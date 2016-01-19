package slack

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const APIBaseURL = "https://slack.com/api"

type APIResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}

type APIClient struct {
	Token  string
	client *http.Client
}

func (c *APIClient) Call(method string, params url.Values) ([]byte, error) {
	params.Add("token", c.Token)

	u := fmt.Sprintf("%s/%s?%s", APIBaseURL, method, params.Encode())

	resp, err := c.client.Get(u)
	if err != nil {
		ErrorLog.Printf("error sending api request: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ErrorLog.Printf("error reading response body: %v\n", err)
		return nil, err
	}

	return body, nil
}

func (c *APIClient) AddReaction(name, channel, timestamp string) error {
	values := url.Values{}
	values.Add("name", name)
	values.Add("channel", channel)
	values.Add("timestamp", timestamp)

	resp, err := c.Call("reactions.add", values)
	if err != nil {
		ErrorLog.Printf("error sending reaction: %v\n", err)
		return err
	}

	raw := json.RawMessage{}
	response := APIResponse{}

	if err := json.Unmarshal(resp, &raw); err != nil {
		ErrorLog.Printf("error unmarshaling raw reaction response: %v\n", err)
		return err
	} else {
		if err := json.Unmarshal(raw, &response); err != nil {
			ErrorLog.Printf("error unmarshaling reaction: %v\n\t%s\n", err, string(raw))
			return err
		}
	}

	if response.Ok {
		return nil
	}

	return errors.New(fmt.Sprintf("Response was not ok! %s", response.Error))
}

func NewAPIClient(token string) *APIClient {
	return &APIClient{
		token,
		&http.Client{},
	}
}
