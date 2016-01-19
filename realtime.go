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
	done chan bool
	ws   *websocket.Conn
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

func (r *RealtimeClient) Stop() {
	r.done <- true
}

func (r *RealtimeClient) ReceiveEvents() <-chan interface{} {
	r.isReady()

	events := make(chan interface{})
	raw := make(chan *json.RawMessage)

	// raw event producer, collecting from the websocket and sending to chan
	go func() {
		defer close(raw)

		for {
			select {
			case <-r.done:
				Log.Println("Stopped receiving events!")
				r.done <- true
				return
			default:
				e := &json.RawMessage{}
				if err := websocket.JSON.Receive(r.ws, &e); err != nil {
					ErrorLog.Printf("error receiving raw event: %v\n", err)
					r.done <- true
				} else {
					raw <- e
				}
			}
		}
	}()

	// consumer of raw events, producer of known slack events
	go func() {
		defer close(events)

		tick := time.NewTicker(30 * time.Second)
		defer tick.Stop()

		for {
			select {
			case <-r.done:
				Log.Println("Stopped processing events!")
				r.done <- true
				return
			case <-tick.C:
				ts, err := r.ping()
				if err != nil {
					r.done <- true
				}
				Log.Printf("Ping! %d\n", ts)
			case e := <-raw:
				EventLog.Println(string(*e))

				realtimeEvent := &Event{}
				if err := UnmarshalRaw(e, &realtimeEvent); err == nil {
					event := GetEventType(realtimeEvent)
					if err := UnmarshalRaw(e, &event); err == nil {
						events <- event
					}
				}
			}
		}
	}()

	return events
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
		make(chan bool),
		new(websocket.Conn),
	}
}
