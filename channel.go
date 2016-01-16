package slack

type Channel struct {
	ID          string       `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	Creator     string       `json:"creator,omitempty"`
	Topic       *Description `json:"topic,omitempty"`
	Purpose     *Description `json:"purpose,omitempty"`
	IsArchived  bool         `json:"is_archived,omitempty"`
	IsGeneral   bool         `json:"is_general,omitempty"`
	IsMember    bool         `json:"is_member,omitempty"`
	LastRead    string       `json:"last_read,omitempty"`
	Latest      *Message     `json:"latest,omitempty"`
	Members     []string     `json:"members,omitempty"`
	UnreadCount int          `json:"unread_count,omitempty"`
}
