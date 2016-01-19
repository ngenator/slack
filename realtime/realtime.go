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

// Client is a client that uses the slack realtime api
type Client struct {
	done chan bool
	ws   *websocket.Conn
}

// Dial connects to the given websocket address
func (c *Client) Dial(url, origin string) error {
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		slack.ErrorLog.Printf("error dialing websocket address: %v\n\t%s\n", err, url)
		return err
	}

	c.ws = ws

	return nil
}

// Stop stops the processing of new events
func (c *Client) Stop() {
	c.done <- true
}

// ReceiveEvents gets events from the websocket and pushes them through a chan
func (c *Client) ReceiveEvents() <-chan interface{} {
	c.isReady()

	events := make(chan interface{})
	raw := make(chan *json.RawMessage)

	// raw event producer, collecting from the websocket and sending to chan
	go func() {
		defer close(raw)

		for {
			select {
			case <-c.done:
				slack.InfoLog.Println("Stopped receiving events!")
				c.done <- true
				return
			default:
				e := &json.RawMessage{}
				if err := websocket.JSON.Receive(c.ws, &e); err != nil {
					slack.ErrorLog.Printf("error receiving raw event: %v\n", err)
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
				slack.InfoLog.Println("Stopped processing events!")
				c.done <- true
				return
			case <-tick.C:
				ts, err := c.ping()
				if err != nil {
					c.done <- true
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
func (c *Client) Send(message interface{}) error {
	c.isReady()

	err := websocket.JSON.Send(c.ws, &message)
	if err != nil {
		slack.ErrorLog.Printf("error sending realtime message: %v\n", err)
		return err
	}

	return nil
}

func (c *Client) isReady() {
	if c.ws == nil || !c.ws.IsClientConn() {
		slack.ErrorLog.Panic("c.ws cannot be nil! did you call Connect()?")
	}
}

func (c *Client) ping() (int64, error) {
	ts := time.Now().Unix()
	if err := c.Send(OutgoingMessage{ts, "ping"}); err != nil {
		slack.ErrorLog.Printf("error sending ping: %v\n", err)
		return ts, err
	}
	return ts, nil
}

// New creates a Realtime client
func New(token string) *Client {
	return &Client{
		make(chan bool),
		new(websocket.Conn),
	}
}
