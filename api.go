package slack

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const BaseURL = "https://slack.com/api"
const OriginURL = "https://slack.com"

type APIClient struct {
	Token  string
	client *http.Client
}

func (c *APIClient) raiseOnError(raw json.RawMessage) error {
	response := Response{}
	if err := json.Unmarshal(raw, &response); err != nil {
		ErrorLog.Println("failed to unmarshal raw api response")
		return err
	}

	if !response.Ok {
		msg := fmt.Sprintf("response not ok: %s", response.Error)
		ErrorLog.Println(msg)
		return fmt.Errorf(msg)
	}

	return nil
}

func (c *APIClient) Call(method string, params *url.Values, response interface{}) error {
	params.Add("token", c.Token)

	u := fmt.Sprintf("%s/%s?%s", BaseURL, method, params.Encode())

	resp, err := c.client.Get(u)
	if err != nil {
		ErrorLog.Printf("error sending api request for %s: %v\n", method, err)
		return err
	}
	defer resp.Body.Close()

	raw := json.RawMessage{}
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		ErrorLog.Printf("error unmarshaling raw api response for %s: %v\n", method, err)
	}

	if err := c.raiseOnError(raw); err != nil {
		return err
	}

	if err := json.Unmarshal(raw, response); err != nil {
		ErrorLog.Printf("error unmarshaling raw api response for %s to %T: %v\n\t%s\n", method, response, err, string(raw))
		return err
	}

	return nil
}

func (c *APIClient) RTMStart() (*RTMStartResponse, error) {
	values := url.Values{}
	// values.Add("mpim_aware", "1")

	response := RTMStartResponse{}
	if err := c.Call("rtm.start", &values, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *APIClient) AddReaction(name, channel, timestamp string) error {
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

func (c *APIClient) GetUser(id string) (*User, error) {
	values := url.Values{}
	values.Add("user", id)

	response := UserResponse{}
	if err := c.Call("users.info", &values, &response); err != nil {
		return nil, err
	}

	return response.User, nil
}

func NewAPIClient(token string) *APIClient {
	return &APIClient{
		token,
		&http.Client{},
	}
}
