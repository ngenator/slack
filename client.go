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
	client http.Client
}

func (self *Client) Get(method string, params url.Values) ([]byte, error) {
	params.Add("token", self.Token)

	u := fmt.Sprintf("%s/%s?%s", APIBaseURL, method, params.Encode())

	resp, err := self.client.Get(u)
	if err != nil {
		log.Println("Error: ", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error: ", err.Error())
		return nil, err
	}

	return body, nil
}

func (self *Client) AddReaction(name, channel, timestamp string) error {
	values := url.Values{}
	values.Add("name", name)
	values.Add("channel", channel)
	values.Add("timestamp", timestamp)

	resp, err := self.Get("reactions.add", values)
	if err != nil {
		log.Println("Error: ", err.Error())
		return err
	}

	response := new(Response)

	if err := json.Unmarshal(resp, &response); err != nil {
		log.Println("Error: ", err.Error())
		return err
	}

	if response.Ok {
		return nil
	}

	return errors.New("Response was not ok")
}
