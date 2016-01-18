package slack

type Channel struct {
	ID          string       `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	Creator     string       `json:"creator,omitempty"`
	Created     Timestamp    `json:"created,omitempty"`
	Topic       *Description `json:"topic,omitempty"`
	Purpose     *Description `json:"purpose,omitempty"`
	IsChannel   bool         `json:"is_channel,omitempty"`
	IsArchived  bool         `json:"is_archived,omitempty"`
	IsGeneral   bool         `json:"is_general,omitempty"`
	IsMember    bool         `json:"is_member,omitempty"`
	LastRead    string       `json:"last_read,omitempty"`
	Latest      *Message     `json:"latest,omitempty"`
	Members     []UserID     `json:"members,omitempty"`
	UnreadCount int          `json:"unread_count,omitempty"`
}

type IM struct {
	ID            string    `json:"id,omitempty"`
	IsIM          bool      `json:"is_im,omitempty"`
	User          UserID    `json:"user,omitempty"`
	IsUserDeleted bool      `json:"is_user_deleted,omitempty"`
	Created       Timestamp `json:"created,omitempty"`
}
