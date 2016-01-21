package slack

type SlashCommandResponse struct {
	Type        string               `json:"response_type,omitempty"`
	Text        string               `json:"text,omitempty"`
	Attachments []*MessageAttachment `json:"attachments,omitempty"`
}
