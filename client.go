package slack

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const APIBaseURL = "https://slack.com/api"

type Response struct {
	Ok bool `json:"ok"`
}

type Client struct {
	Token  string
	client *http.Client
}

func (c *Client) Get(method string, params url.Values) ([]byte, error) {
	params.Add("token", c.Token)

	u := fmt.Sprintf("%s/%s?%s", APIBaseURL, method, params.Encode())

	resp, err := c.client.Get(u)
	if err != nil {
		log.Printf("error sending api request: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error reading response body: %v\n", err)
		return nil, err
	}

	return body, nil
}

func (c *Client) AddReaction(name, channel, timestamp string) error {
	values := url.Values{}
	values.Add("name", name)
	values.Add("channel", channel)
	values.Add("timestamp", timestamp)

	resp, err := c.Get("reactions.add", values)
	if err != nil {
		log.Printf("error sending reaction: %v\n", err)
		return err
	}

	response := &Response{}

	if err := json.Unmarshal(resp, &response); err != nil {
		log.Printf("error unmarshaling reaction response: %v\n", err)
		return err
	}

	if response.Ok {
		return nil
	}

	return errors.New("Response was not ok")
}

func NewClient(token string) *Client {
	return &Client{
		Token:  token,
		client: &http.Client{},
	}
}
