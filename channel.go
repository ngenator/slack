package slack

type Channel struct {
	ID          string      `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Creator     string      `json:"creator,omitempty"`
	Topic       *Topic      `json:"topic,omitempty"`
	Purpose     *Purpose    `json:"purpose,omitempty"`
	IsArchived  bool        `json:"is_archived,omitempty"`
	IsGeneral   bool        `json:"is_general,omitempty"`
	IsMember    bool        `json:"is_member,omitempty"`
	LastRead    string      `json:"last_read,omitempty"`
	Latest      interface{} `json:"latest,omitempty"`
	Members     []string    `json:"members,omitempty"`
	UnreadCount int         `json:"unread_count,omitempty"`
}

type Topic struct {
	Value   string `json:"string,omitempty"`
	Creator string `json:"creator,omitempty"`
	LastSet int    `json:"last_set,omitempty"`
}

type Purpose struct {
	Topic
}

type ChannelResponse struct {
	Ok       bool      `json:"ok,omitempty"`
	Channels []Channel `json:"channels,omitempty"`
}
