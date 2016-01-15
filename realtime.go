package slack

import (
	"encoding/json"
	"log"
	"net/url"
	"time"

	"golang.org/x/net/websocket"
)

type RTMStart struct {
	Ok       bool      `json:"ok,omitempty"`
	URL      string    `json:"url,omitempty"`
	Channels []Channel `json:"channels,omitempty"`
	Users    []User    `json:"users,omitempty"`
}

type RTMEvent struct {
	Type      string `json:"type,omitempty"`
	SubType   string `json:"subtype,omitempty"`
	Hidden    bool   `json:"hidden,omitempty"`
	Timestamp string `json:"ts,omitempty"`
	Username  string `json:"username,omitempty"`
	User      string `json:"user,omitempty"`
	Channel   string `json:"channel,omitempty"`
	Text      string `json:"text,omitempty"`
	Edited    struct {
		Timestamp string `json:"ts,omitempty"`
		User      string `json:"user,omitempty"`
	} `json:"edited,omitempty"`
	Error struct {
		Code    int    `json:"code,omitempty"`
		Message string `json:"msg,omitempty"`
	} `json:"error,omitempty"`
}

type RTMMessage struct{}

type RTMPing struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

type Realtime struct {
	Client
	Events   chan RTMEvent
	Messages chan RTMMessage
	done     chan bool
	ws       *websocket.Conn
}

func (self *Realtime) Connect() (users map[string]User, channels map[string]Channel) {
	body, err := self.Get("rtm.start", url.Values{})
	if err != nil {
		log.Fatal(err.Error())
	}

	start := new(RTMStart)

	if err := json.Unmarshal(body, &start); err != nil {
		log.Fatal(err.Error())
	}

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
		log.Fatal(err.Error())
	}

	self.ws = ws

	return users, channels
}

func (self *Realtime) Listen() {
	var e RTMEvent

	self.check()

	tick := time.NewTicker(30 * time.Second)
	defer tick.Stop()

	for {
		select {
		case <-self.done:
			log.Println("Stopped!")
			close(self.done)
			return
		case <-tick.C:
			log.Println("Ping!")
			err := self.ping()
			if isError(err) {
				self.done <- true
			}
		default:
			err := websocket.JSON.Receive(self.ws, &e)
			if !isError(err) {
				self.Events <- e
			}
		}
	}
}

func (self *Realtime) check() {
	if !self.ws.IsClientConn() {
		log.Panic("ws cannot be nil! did you call Start()?")
	}
}

func (self *Realtime) ping() error {
	self.check()
	err := websocket.JSON.Send(self.ws, &RTMPing{1, "ping"})
	if isError(err) {
		self.done <- true
		return err
	}
	return nil
}

func isError(err error) bool {
	if err != nil {
		log.Printf("%10s %s\n", "Error:", err.Error())
		return true
	}
	return false
}

func NewRealtime(token string) *Realtime {
	return &Realtime{
		Client{
			Token: token,
		},
		make(chan RTMEvent),
		make(chan RTMMessage),
		make(chan bool),
		new(websocket.Conn),
	}
}
