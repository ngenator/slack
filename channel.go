package slack

// Channel type
type Channel struct {
	ID ChannelID `json:"id,omitempty"`

	Name    string     `json:"name,omitempty"`
	Creator UserID     `json:"creator,omitempty"`
	Created *Timestamp `json:"created,omitempty"`

	Topic   *ChannelDescription `json:"topic,omitempty"`
	Purpose *ChannelDescription `json:"purpose,omitempty"`

	IsChannel  bool `json:"is_channel,omitempty"`
	IsArchived bool `json:"is_archived,omitempty"`
	IsGeneral  bool `json:"is_general,omitempty"`
	IsMember   bool `json:"is_member,omitempty"`

	LastRead    UniqueTimestamp `json:"last_read,omitempty"`
	Latest      *Message        `json:"latest,omitempty"`
	NumMembers  int             `json:"num_members,omitempty"`
	Members     []UserID        `json:"members,omitempty"`
	UnreadCount int             `json:"unread_count,omitempty"`
}

// ChannelDescription type
type ChannelDescription struct {
	Value   string     `json:"string,omitempty"`
	Creator UserID     `json:"creator,omitempty"`
	LastSet *Timestamp `json:"last_set,omitempty"`
}

// IMChannel type
type IMChannel struct {
	UserID             UserID     `json:"user,omitempty"`
	IsIM               bool       `json:"is_im,omitempty"`
	IsOpen             bool       `json:"is_open,omitempty"`
	IsUserDeleted      bool       `json:"is_user_deleted,omitempty"`
	Created            *Timestamp `json:"created,omitempty"`
	Latest             *Message   `json:"latest,omitempty"`
	UnreadCount        int        `json:"unread_count,omitempty"`
	UnreadCountDisplay int        `json:"unread_count_display,omitempty"`
}
