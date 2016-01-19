package slack

import (
	"encoding/json"
	"net/url"
	"time"

	"golang.org/x/net/websocket"
)

type StartResponse struct {
	Ok       bool       `json:"ok,omitempty"`
	URL      string     `json:"url,omitempty"`
	Channels []*Channel `json:"channels,omitempty"`
	Users    []*User    `json:"users,omitempty"`
	// TODO: add the rest of the initial data
}

type OutgoingMessage struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
}

type RealtimeClient struct {
	APIClient
	Slack
	RawEvents chan *json.RawMessage
	Events    chan interface{}
	done      chan bool
	ws        *websocket.Conn
}

func (r *RealtimeClient) Connect() error {
	body, err := r.APIClient.Call("rtm.start", url.Values{})
	if err != nil {
		ErrorLog.Printf("error sending rtm.start request: %v\n", err)
		return err
	}

	response := new(StartResponse)

	if err := json.Unmarshal(body, &response); err != nil {
		ErrorLog.Printf("error unmarshaling rtm.start response: %v\n", err)
		return err
	}

	for _, u := range response.Users {
		r.Slack.Users[u.ID] = u
	}

	for _, c := range response.Channels {
		r.Slack.Channels[c.ID] = c
	}

	ws, err := websocket.Dial(response.URL, "", "https://slack.com")
	if err != nil {
		ErrorLog.Printf("error dialing websocket address: %v\n\t%s\n", err, response.URL)
		return err
	}

	r.ws = ws

	return nil
}

func (r *RealtimeClient) Start() chan bool {
	done := make(chan bool)
	r.ReceiveRawEvents(done)
	r.ProcessEvents(done)
	return done
}

func (r *RealtimeClient) ReceiveRawEvents(done chan bool) {
	go func() {
		defer close(r.RawEvents)

		tick := time.NewTicker(30 * time.Second)
		defer tick.Stop()

		for {
			r.isReady()

			select {
			case <-done:
				Log.Println("Stopped processing raw events!")
				done <- true
				break
			case <-tick.C:
				ts, err := r.ping()
				Log.Printf("Ping! %d\n", ts)
				if err != nil {
					done <- true
				}
			default:
				raw := &json.RawMessage{}
				if err := websocket.JSON.Receive(r.ws, &raw); err != nil {
					ErrorLog.Printf("error unmarshaling raw event: %v\n", err)
				} else {
					r.RawEvents <- raw
				}
			}
		}
	}()
}

func (r *RealtimeClient) ProcessEvents(done chan bool) {
	go func() {
		defer close(r.Events)

		for {
			select {
			case <-done:
				Log.Println("Stopped processing events!")
				done <- true
				break
			case raw := <-r.RawEvents:
				EventLog.Println(string(*raw))
				realtimeEvent := &Event{}
				if err := UnmarshalRaw(raw, &realtimeEvent); err == nil {
					event := GetEventType(realtimeEvent)
					if err := UnmarshalRaw(raw, &event); err == nil {
						r.Events <- event
					}
				}
			}
		}
	}()
}

func (r *RealtimeClient) Send(message interface{}) error {
	r.isReady()

	err := websocket.JSON.Send(r.ws, &message)
	if err != nil {
		ErrorLog.Printf("error sending realtime message: %v\n", err)
		return err
	}

	return nil
}

func (r *RealtimeClient) isReady() {
	if !r.ws.IsClientConn() {
		ErrorLog.Panic("r.ws cannot be nil! did you call Connect()?")
	}
}

func (r *RealtimeClient) ping() (int64, error) {
	ts := time.Now().Unix()
	if err := r.Send(OutgoingMessage{ts, "ping"}); err != nil {
		ErrorLog.Printf("error sending ping: %v\n", err)
		return ts, err
	}
	return ts, nil
}

func NewRealtimeClient(token string) *RealtimeClient {
	return &RealtimeClient{
		*NewAPIClient(token),
		*NewSlack(),
		make(chan *json.RawMessage),
		make(chan interface{}),
		make(chan bool),
		new(websocket.Conn),
	}
}
