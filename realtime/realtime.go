package realtime

import (
	"encoding/json"
	"time"

	"github.com/ngenator/slack"
	"golang.org/x/net/websocket"
)

// OutgoingMessage type
type OutgoingMessage struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
}

// Realtime is a client that uses the slack realtime api
type Realtime struct {
	done chan bool
	ws   *websocket.Conn
}

// Dial connects to the given websocket address
func (r *Realtime) Dial(url, origin string) error {
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		slack.ErrorLog.Printf("error dialing websocket address: %v\n\t%s\n", err, url)
		return err
	}

	r.ws = ws

	return nil
}

// Stop stops the processing of new events
func (r *Realtime) Stop() {
	r.done <- true
}

// ReceiveEvents gets events from the websocket and pushes them through a chan
func (r *Realtime) ReceiveEvents() <-chan interface{} {
	r.isReady()

	events := make(chan interface{})
	raw := make(chan *json.RawMessage)

	// raw event producer, collecting from the websocket and sending to chan
	go func() {
		defer close(raw)

		for {
			select {
			case <-r.done:
				slack.InfoLog.Println("Stopped receiving events!")
				r.done <- true
				return
			default:
				e := &json.RawMessage{}
				if err := websocket.JSON.Receive(r.ws, &e); err != nil {
					slack.ErrorLog.Printf("error receiving raw event: %v\n", err)
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
				slack.InfoLog.Println("Stopped processing events!")
				r.done <- true
				return
			case <-tick.C:
				ts, err := r.ping()
				if err != nil {
					r.done <- true
				}
				slack.InfoLog.Printf("Ping! %d\n", ts)
			case e := <-raw:
				slack.EventLog.Println(string(*e))

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

// Send sends a message via the websocket
func (r *Realtime) Send(message interface{}) error {
	r.isReady()

	err := websocket.JSON.Send(r.ws, &message)
	if err != nil {
		slack.ErrorLog.Printf("error sending realtime message: %v\n", err)
		return err
	}

	return nil
}

func (r *Realtime) isReady() {
	if r.ws == nil || !r.ws.IsClientConn() {
		slack.ErrorLog.Panic("r.ws cannot be nil! did you call Connect()?")
	}
}

func (r *Realtime) ping() (int64, error) {
	ts := time.Now().Unix()
	if err := r.Send(OutgoingMessage{ts, "ping"}); err != nil {
		slack.ErrorLog.Printf("error sending ping: %v\n", err)
		return ts, err
	}
	return ts, nil
}

// New creates a Realtime client
func New(token string) *Realtime {
	return &Realtime{
		make(chan bool),
		new(websocket.Conn),
	}
}
