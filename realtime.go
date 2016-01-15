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

type RTBotMessage struct {
	Attachments []struct {
		ID       int    `json:"id"`
		Fallback string `json:"fallback"`
		Color    string `json:"color"`
		Fields   []struct {
			Title string `json:"title"`
			Value string `json:"value"`
		} `json:"fields"`
	} `json:"attachments"`
}

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

func (r *Realtime) Connect() (users map[string]User, channels map[string]Channel) {
	body, err := r.Get("rtm.start", url.Values{})
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

	r.ws = ws

	return users, channels
}

func (r *Realtime) Listen() {
	m := json.RawMessage{}
	e := RTMEvent{}

	r.check()

	tick := time.NewTicker(30 * time.Second)
	defer tick.Stop()

	for {
		select {
		case <-r.done:
			log.Println("Stopped!")
			close(r.done)
			return
		case <-tick.C:
			log.Println("Ping!")
			err := r.ping()
			if isError(err) {
				r.done <- true
			}
		default:
			err := websocket.JSON.Receive(r.ws, &m)
			if !isError(err) {
				if err := json.Unmarshal(m, &e); err != nil {
					log.Println(err.Error(), string(m))
				} else {
					r.Events <- e
				}
			} else {
				log.Println(err.Error())
			}
		}
	}
}

func (r *Realtime) check() {
	if !r.ws.IsClientConn() {
		log.Panic("ws cannot be nil! did you call Start()?")
	}
}

func (r *Realtime) ping() error {
	r.check()
	err := websocket.JSON.Send(r.ws, &RTMPing{1, "ping"})
	if isError(err) {
		r.done <- true
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
