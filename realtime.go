package slack

import (
	"encoding/json"
	"time"

	"golang.org/x/net/websocket"
)

// OutgoingMessage type
type OutgoingMessage struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
}

// RealtimeClient is a client that uses the slack realtime api
type RealtimeClient struct {
	done chan bool
	ws   *websocket.Conn
}

// Dial connects to the given websocket address
func (c *RealtimeClient) Dial(url, origin string) error {
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		ErrorLog.Printf("error dialing websocket address: %v\n\t%s\n", err, url)
		return err
	}

	c.ws = ws

	return nil
}

// Stop stops the processing of new events
func (c *RealtimeClient) Stop() {
	InfoLog.Println("Stopping realtime event receiver...")
	c.done <- true
	<-c.done
	InfoLog.Println("Done!")
}

// ReceiveEvents gets events from the websocket and pushes them through a chan
func (c *RealtimeClient) ReceiveEvents() <-chan interface{} {
	c.isReady()

	events := make(chan interface{})
	raw := make(chan *json.RawMessage)

	// raw event producer, collecting from the websocket and sending to chan
	go func() {
		defer close(raw)

		for {
			select {
			case <-c.done:
				InfoLog.Println("Stopped receiving events!")
				c.done <- true
				return
			default:
				e := &json.RawMessage{}
				if err := websocket.JSON.Receive(c.ws, &e); err != nil {
					ErrorLog.Printf("error receiving raw event: %v\n", err)
					c.done <- true
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
			case <-c.done:
				InfoLog.Println("Stopped processing events!")
				c.done <- true
				return
			case <-tick.C:
				ts, err := c.ping()
				if err != nil {
					c.done <- true
				}
				InfoLog.Printf("Ping! %d\n", ts)
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

// Send sends a message via the websocket
func (c *RealtimeClient) Send(message interface{}) error {
	c.isReady()

	err := websocket.JSON.Send(c.ws, &message)
	if err != nil {
		ErrorLog.Printf("error sending realtime message: %v\n", err)
		return err
	}

	return nil
}

func (c *RealtimeClient) isReady() {
	if c.ws == nil || !c.ws.IsClientConn() {
		ErrorLog.Panic("c.ws cannot be nil! did you call Connect()?")
	}
}

func (c *RealtimeClient) ping() (int64, error) {
	ts := time.Now().Unix()
	if err := c.Send(OutgoingMessage{ts, "ping"}); err != nil {
		ErrorLog.Printf("error sending ping: %v\n", err)
		return ts, err
	}
	return ts, nil
}

// NewRealtimeClient creates a Realtime client
func NewRealtimeClient() *RealtimeClient {
	return &RealtimeClient{
		make(chan bool),
		new(websocket.Conn),
	}
}
