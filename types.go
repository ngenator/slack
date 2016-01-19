package slack

// User type
type User struct {
	ID UserID `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Deleted bool `json:"deleted,omitempty"`

	Color string `json:"color,omitempty"`

	Profile *UserProfile `json:"profile,omitempty"`

	IsAdmin           bool `json:"is_admin,omitempty"`
	IsOwner           bool `json:"is_owner,omitempty"`
	IsPrimaryOwner    bool `json:"is_primary_owner,omitempty"`
	IsRestricted      bool `json:"is_restricted,omitempty"`
	IsUltraRestricted bool `json:"is_ultra_restricted,omitempty"`
	IsBot             bool `json:"is_bot,omitempty"`

	HasTwoFactorAuth bool   `json:"has_2fa,omitempty"`
	TwoFactorType    string `json:"two_factor_type,omitempty"`
	HasFiles         bool   `json:"has_files,omitempty"`
}

type UserProfile struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	RealName  string `json:"real_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Skype     string `json:"skype,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

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

	LastRead    SlackTimestamp `json:"last_read,omitempty"`
	Latest      *Message       `json:"latest,omitempty"`
	Members     []UserID       `json:"members,omitempty"`
	UnreadCount int            `json:"unread_count,omitempty"`
}

// ChannelDescription type
type ChannelDescription struct {
	Value   string     `json:"string,omitempty"`
	Creator UserID     `json:"creator,omitempty"`
	LastSet *Timestamp `json:"last_set,omitempty"`
}

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

// Message type
type Message struct {
	Text        string              `json:"text,omitempty"`
	UserID      UserID              `json:"user,omitempty"`
	ChannelID   ChannelID           `json:"channel,omitempty"`
	Timestamp   SlackTimestamp      `json:"ts,omitempty"`
	Attachments []MessageAttachment `json:"attachments,omitempty"`
	ReplyTo     int                 `json:"reply_to,omitempty"`
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

	ImageUrl string `json:"image_url,omitempty"`
	ThumbUrl string `json:"thumb_url,omitempty"`

	Fields []MessageAttachmentField `json:"fields,omitempty"`
}

// MessageAttachmentField type
type MessageAttachmentField struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short bool   `json:"short,omitempty"`
}
