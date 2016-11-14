package slack

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
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
	urlParams := url.Values{}
	urlParams.Add("token", c.Token)

	// Create the url from the BaseURL, api endpoint, and the api key
	u := fmt.Sprintf("%s/%s?%s", BaseURL, method, urlParams.Encode())

	resp, err := c.client.Post(u, "application/x-www-form-urlencoded", strings.NewReader(params.Encode()))
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

// TODO: handle optional args
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

// TODO: handle more of the optional args
func (c *APIClient) PostChatMessage(channel, text, username, iconEmoji string) error {
	values := url.Values{}
	values.Add("channel", channel)
	values.Add("text", text)
	values.Add("username", username)
	values.Add("icon_emoji", iconEmoji)
	values.Add("link_names", "1")

	response := Response{}
	if err := c.Call("chat.postMessage", &values, &response); err != nil {
		return err
	}

	return nil
}

func (c *APIClient) GetUsers() ([]*User, error) {
	values := url.Values{}

	response := UsersResponse{}
	if err := c.Call("users.list", &values, &response); err != nil {
		return nil, err
	}

	return response.Users, nil
}

func (c *APIClient) GetUserGroups() ([]*UserGroup, error) {
	values := url.Values{}
	values.Add("include_users", "1")

	response := UserGroupsResponse{}
	if err := c.Call("usergroups.list", &values, &response); err != nil {
		return nil, err
	}

	return response.UserGroups, nil
}

func (c *APIClient) GetUserIDsInUserGroup(usergroup string) ([]string, error) {
	values := url.Values{}
	values.Add("usergroup", usergroup)

	response := UserIDsResponse{}
	if err := c.Call("usergroups.users.list", &values, &response); err != nil {
		return nil, err
	}

	return response.UserIDs, nil
}

func (c *APIClient) UpdateUserIDsInUserGroup(userGroup string, userIDs []string) error {
	values := url.Values{}
	values.Add("usergroup", userGroup)
	values.Add("users", strings.Join(userIDs, ","))

	response := UserGroupUpdateResponse{}
	if err := c.Call("usergroups.users.update", &values, &response); err != nil {
		return err
	}

	return nil
}

func (c *APIClient) GetUserDirectMessageChannel(userID string) (*Channel, error) {
	values := url.Values{}
	values.Add("user", userID)
	values.Add("return_im", "1")

	response := ChannelResponse{}
	if err := c.Call("im.open", &values, &response); err != nil {
		return nil, err
	}

	return response.Channel, nil
}

func NewAPIClient(token string) *APIClient {
	return &APIClient{
		token,
		&http.Client{},
	}
}
