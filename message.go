package slack

// Message type
type Message struct {
	Text        string               `json:"text,omitempty"`
	UserID      UserID               `json:"user,omitempty"`
	ChannelID   ChannelID            `json:"channel,omitempty"`
	Timestamp   UniqueTimestamp      `json:"ts,omitempty"`
	Attachments []*MessageAttachment `json:"attachments,omitempty"`
	ReplyTo     int64                `json:"reply_to,omitempty"`
}

// MessageAttachment type
type MessageAttachment struct {
	ID int `json:"id,omitempty"`

	Text     string `json:"text,omitempty"`
	PreText  string `json:"pretext,omitempty"`
	Fallback string `json:"fallback,omitempty"`

	Color string `json:"color,omitempty"`

	AuthorName string `json:"author_name,omitempty"`
	AuthorLink string `json:"author_link,omitempty"`
	AuthorIcon string `json:"author_icon,omitempty"`

	Title     string `json:"title,omitempty"`
	TitleLink string `json:"title_link,omitempty"`

	ImageURL string `json:"image_url,omitempty"`
	ThumbURL string `json:"thumb_url,omitempty"`

	Fields []*MessageAttachmentField `json:"fields,omitempty"`
}

// MessageAttachmentField type
type MessageAttachmentField struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short bool   `json:"short,omitempty"`
}
