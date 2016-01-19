package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/ngenator/slack"
)

const BaseURL = "https://slack.com/api"

type Response struct {
	Ok    bool  `json:"ok"`
	Error Error `json:"error,omitempty"`
}

type RTMStartResponse struct {
	Response
	URL      string           `json:"url,omitempty"`
	Channels []*slack.Channel `json:"channels,omitempty"`
	Users    []*slack.User    `json:"users,omitempty"`
	// TODO: add the rest of the initial data
}

type Client struct {
	Token  string
	client *http.Client
}

func (c *Client) raiseOnError(response interface{}) error {
	r := response.(Response)
	if !r.Ok {
		msg := fmt.Sprintf("response not ok: %s", r.Error)
		log.Println(msg)
		return fmt.Errorf(msg)
	}
	return nil
}

func (c *Client) Call(method string, params *url.Values, response interface{}) error {
	params.Add("token", c.Token)

	u := fmt.Sprintf("%s/%s?%s", BaseURL, method, params.Encode())

	resp, err := c.client.Get(u)
	if err != nil {
		log.Printf("error sending api request for %s: %v\n", method, err)
		return err
	}
	defer resp.Body.Close()

	raw := json.RawMessage{}
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		log.Printf("error unmarshaling raw api response for %s: %v\n", method, err)
	}

	if err := json.Unmarshal(raw, response); err != nil {
		log.Printf("error unmarshaling raw api response for %s to %T: %v\n\t%s\n", method, response, err, string(raw))
		return err
	}

	return nil
}

func (c *Client) RTMStart() (*RTMStartResponse, error) {
	values := url.Values{}
	// values.Add("mpim_aware", "1")

	response := RTMStartResponse{}
	if err := c.Call("rtm.start", &values, &response); err != nil {
		return nil, err
	}

	if err := c.raiseOnError(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) AddReaction(name, channel, timestamp string) error {
	values := url.Values{}
	values.Add("name", name)
	values.Add("channel", channel)
	values.Add("timestamp", timestamp)

	response := Response{}
	if err := c.Call("reactions.add", &values, &response); err != nil {
		return err
	}

	return nil
}

func New(token string) *Client {
	return &Client{
		token,
		&http.Client{},
	}
}
