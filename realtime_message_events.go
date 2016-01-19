package slack

var MessageEventTypes = map[string]interface{}{
	"bot_message": &BotMessageEvent{},
	"me_message":  &MeMessageEvent{},

	"message_changed": &MessageChangedEvent{},
	"message_deleted": &MessageDeletedEvent{},

	"channel_join":      &ChannelJoinMessageEvent{},
	"channel_leave":     &ChannelLeaveMessageEvent{},
	"channel_topic":     &ChannelTopicMessageEvent{},
	"channel_purpose":   &ChannelPurposeMessageEvent{},
	"channel_name":      &ChannelNameMessageEvent{},
	"channel_archive":   &ChannelArchiveMessageEvent{},
	"channel_unarchive": &ChannelUnarchiveMessageEvent{},

	"group_join":      &GroupJoinMessageEvent{},
	"group_leave":     &GroupLeaveMessageEvent{},
	"group_topic":     &GroupTopicMessageEvent{},
	"group_purpose":   &GroupPurposeMessageEvent{},
	"group_name":      &GroupNameMessageEvent{},
	"group_archive":   &GroupArchiveMessageEvent{},
	"group_unarchive": &GroupUnarchiveMessageEvent{},

	"file_share":   &FileShareMessageEvent{},
	"file_comment": &FileCommentMessageEvent{},
	"file_mention": &FileMentionMessageEvent{},

	"pinned_item":   &PinnedItemMessageEvent{},
	"unpinned_item": &UnpinnedItemMessageEvent{},
}

type MessageEvent struct {
	Message
}

type BotMessageEvent struct {
	BotID string `json:"bot_id,omitempty"`
	Message
}

type MeMessageEvent struct {
	Message
}

type MessageChangedEvent struct {
	Hidden bool `json:"hidden,omitempty"`
	Message
}

type MessageDeletedEvent struct {
	DeletedTimestamp UniqueTimestamp `json:"deleted_ts,omitempty"`
	Timestamp        UniqueTimestamp `json:"ts,omitempty"`
	Hidden           bool            `json:"hidden,omitempty"`
}

type ChannelJoinMessageEvent struct {
	Text      string          `json:"text,omitempty"`
	Inviter   UserID          `json:"inviter,omitempty"`
	Timestamp UniqueTimestamp `json:"ts,omitempty"`
	UserID    UserID          `json:"user,omitempty"`
}

type ChannelLeaveMessageEvent struct{}
type ChannelTopicMessageEvent struct{}
type ChannelPurposeMessageEvent struct{}
type ChannelNameMessageEvent struct{}
type ChannelArchiveMessageEvent struct{}
type ChannelUnarchiveMessageEvent struct{}
type GroupJoinMessageEvent struct{}
type GroupLeaveMessageEvent struct{}
type GroupTopicMessageEvent struct{}
type GroupPurposeMessageEvent struct{}
type GroupNameMessageEvent struct{}
type GroupArchiveMessageEvent struct{}
type GroupUnarchiveMessageEvent struct{}
type FileShareMessageEvent struct{}
type FileCommentMessageEvent struct{}
type FileMentionMessageEvent struct{}
type PinnedItemMessageEvent struct{}
type UnpinnedItemMessageEvent struct{}
