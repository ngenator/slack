package slack

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"golang.org/x/net/websocket"
)

type RTMStart struct {
	Ok       bool      `json:"ok,omitempty"`
	URL      string    `json:"url,omitempty"`
	Channels []Channel `json:"channels,omitempty"`
	Users    []User    `json:"users,omitempty"`
	// TODO: add the rest of the initial data
}

type RTMEvent struct {
	// TODO: move the message specific stuff elsewhere
	Type      string     `json:"type,omitempty"`
	SubType   string     `json:"subtype,omitempty"`
	Hidden    bool       `json:"hidden,omitempty"`
	Timestamp string     `json:"ts,omitempty"`
	Username  string     `json:"username,omitempty"`
	User      string     `json:"user,omitempty"`
	Channel   string     `json:"channel,omitempty"`
	Text      string     `json:"text,omitempty"`
	Edited    *RTMEdited `json:"edited,omitempty"`
	Error     *RTMError  `json:"error,omitempty"`
}

type RTMError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"msg,omitempty"`
}

func (e *RTMError) String() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

type RTMEdited struct {
	Timestamp string `json:"ts,omitempty"`
	User      string `json:"user,omitempty"`
}

type RTMMessage struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

// TODO: split each event type into it's own chan and push messages to the appropriate chan
type RealtimeClient struct {
	Client
	Events   chan RTMEvent
	Messages chan RTMMessage
	done     chan bool
	ws       *websocket.Conn
}

func (r *RealtimeClient) Connect() (users map[string]User, channels map[string]Channel) {
	body, err := r.Get("rtm.start", url.Values{})
	if err != nil {
		ErrorLog.Fatalf("error sending rtm.start request: %v\n", err)
	}

	start := new(RTMStart)

	if err := json.Unmarshal(body, &start); err != nil {
		ErrorLog.Fatalf("error unmarshaling rtm.start response: %v\n", err)
	}

	// TODO: clean this up and store these on the instance
	users = make(map[string]User)
	channels = make(map[string]Channel)

	for _, u := range start.Users {
		users[u.ID] = u
	}

	for _, c := range start.Channels {
		channels[c.ID] = c
	}

	ws, err := websocket.Dial(start.URL, "", "https://slack.com")
	if err != nil {
		ErrorLog.Fatalf("error dialing websocket address: %v\n\t%s\n", err, start.URL)
	}

	r.ws = ws

	return users, channels
}

func (r *RealtimeClient) Listen() {
	r.isReady()

	m := json.RawMessage{}
	e := RTMEvent{}

	tick := time.NewTicker(30 * time.Second)
	defer tick.Stop()

	for {
		select {
		case <-r.done:
			Log.Println("Stopped!")
			close(r.done)
			return
		case <-tick.C:
			Log.Println("Ping!")
			err := r.ping()
			if err != nil {
				r.done <- true
			}
		default:
			if err := websocket.JSON.Receive(r.ws, &m); err != nil {
				ErrorLog.Printf("error unmarshaling raw event: %v\n", err)
			} else {
				EventLog.Println(string(m))
				if err := json.Unmarshal(m, &e); err != nil {
					ErrorLog.Printf("error unmarshaling event: %v\n\t%s\n", err, string(m))
				} else {
					r.Events <- e
				}
			}
		}
	}
}

func (r *RealtimeClient) Send(message interface{}) error {
	r.isReady()

	err := websocket.JSON.Send(r.ws, &message)
	if err != nil {
		ErrorLog.Printf("error sending realtime message: %v\n", err)
		r.done <- true
		return err
	}

	return nil
}

func (r *RealtimeClient) isReady() {
	if !r.ws.IsClientConn() {
		ErrorLog.Panic("r.ws cannot be nil! did you call Connect()?")
	}
}

func (r *RealtimeClient) ping() error {
	if err := r.Send(RTMMessage{1, "ping"}); err != nil {
		ErrorLog.Printf("error sending ping: %v\n", err)
		return err
	}
	return nil
}

func NewRealtimeClient(token string) *RealtimeClient {
	return &RealtimeClient{
		*NewClient(token),
		make(chan RTMEvent),
		make(chan RTMMessage),
		make(chan bool),
		new(websocket.Conn),
	}
}
