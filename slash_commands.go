package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// SlashCommandResponse is a reply to a slash command
type SlashCommandResponse struct {
	Type        string               `json:"response_type,omitempty"`
	Text        string               `json:"text,omitempty"`
	ChannelID   string               `json:"channel,omitempty"`
	Username    string               `json:"username,omitempty"`
	IconURL     string               `json:"icon_url,omitempty"`
	IconEmoji   string               `json:"icon_emoji,omitempty"`
	UnfurlLinks bool                 `json:"unfurl_links,omitempty"`
	Attachments []*MessageAttachment `json:"attachments,omitempty"`
}

// SlashCommand is a slash command converted to json from application/x-www-form-urlencoded
type SlashCommand struct {
	Token       string `json:"token"`
	TeamID      string `json:"team_id"`
	TeamDomain  string `json:"team_domain"`
	ChannelID   string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	UserID      string `json:"user_id"`
	UserName    string `json:"user_name"`
	Command     string `json:"command"`
	Text        string `json:"text"`
	ResponseURL string `json:"response_url"`
}

// RespondToSlashCommand responds to a slack slash command by posting json to its response url
func RespondToSlashCommand(responseURL string, response *SlashCommandResponse) error {
	body, err := json.Marshal(response)
	if err != nil {
		return err
	}

	resp, err := http.Post(responseURL, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(t))
	}

	return nil
}
