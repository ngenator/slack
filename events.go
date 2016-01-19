package slack

import (
	"encoding/json"
	"fmt"
)

type Event struct {
	Type    string      `json:"type"`
	SubType string      `json:"subtype,omitempty"`
	Error   *EventError `json:"error,omitempty"`
}

type EventError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"msg,omitempty"`
}

func (e *EventError) String() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

type HelloEvent struct {
}

type PongEvent struct {
	ReplyTo int64 `json:"reply_to,omitempty"`
}

type PresenceChangeEvent struct {
	Presence string `json:"presence,omitempty"`
	UserID   UserID `json:"user,omitempty"`
}

type UserTypingEvent struct {
	ChannelID ChannelID `json:"channel,omitempty"`
	UserID    UserID    `json:"user,omitempty"`
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
	DeletedTimestamp SlackTimestamp `json:"deleted_ts,omitempty"`
	Timestamp        SlackTimestamp `json:"ts,omitempty"`
	Hidden           bool           `json:"hidden,omitempty"`
}

type ChannelMarkedEvent struct {
	ChannelID           ChannelID      `json:"channel,omitempty"`
	Timestamp           SlackTimestamp `json:"ts,omitempty"`
	UnreadCount         int            `json:"unread_count,omitempty"`
	UnreadCountDisplay  int            `json:"unread_count_display,omitempty"`
	NumMentions         int            `json:"num_mentions,omitempty"`
	NumMentionsDisplay  int            `json:"num_mentions_display,omitempty"`
	MentionCount        int            `json:"mention_count,omitempty"`
	MentionCountDisplay int            `json:"mention_count_display,omitempty"`
}

type ChannelJoinEvent struct {
	Text      string         `json:"text,omitempty"`
	Inviter   UserID         `json:"inviter,omitempty"`
	Timestamp SlackTimestamp `json:"ts,omitempty"`
	UserID    UserID         `json:"user,omitempty"`
}

type ChannelLeaveEvent struct {
}

type ChannelTopicEvent struct {
}

type ChannelPurposeEvent struct {
}

type ChannelNameEvent struct {
}

type ChannelArchiveEvent struct {
}

type ChannelUnArchiveEvent struct {
}

type ReactionAddedEvent struct {
}

type ReactionRemovedEvent struct {
}

type ImMarkedEvent struct {
	ChannelID           ChannelID      `json:"channel,omitempty"`
	Timestamp           SlackTimestamp `json:"ts,omitempty"`
	DmCount             int            `json:"dm_count,omitempty"`
	UnreadCountDisplay  int            `json:"unread_count_display,omitempty"`
	NumMentionsDisplay  int            `json:"num_mentions_display,omitempty"`
	MentionCountDisplay int            `json:"mention_count_display,omitempty"`
}

type ImCreatedEvent struct {
	UserID  UserID    `json:"user,omitempty"`
	Channel IMChannel `json:"channel,omitempty"`
}

type ImOpenEvent struct {
	UserID    UserID    `json:"user,omitempty"`
	ChannelID ChannelID `json:"channel,omitempty"`
}

type ImCloseEvent struct {
	UserID    UserID    `json:"user,omitempty"`
	ChannelID ChannelID `json:"channel,omitempty"`
}

type GroupMarkedEvent struct {
}

var EventTypes = map[string]interface{}{
	"hello":           &HelloEvent{},
	"pong":            &PongEvent{},
	"message":         &MessageEvent{},
	"user_typing":     &UserTypingEvent{},
	"presence_change": &PresenceChangeEvent{},
	"channel_marked":  &ChannelMarkedEvent{},
	"group_marked":    &GroupMarkedEvent{},
	"im_marked":       &ImMarkedEvent{},
	"im_created":      &ImCreatedEvent{},
	"im_open":         &ImOpenEvent{},
	"im_close":        &ImCloseEvent{},
}

var EventSubTypes = map[string]interface{}{
	"bot_message": &BotMessageEvent{},
	// TODO: add types for these events
	"me_message":        &MeMessageEvent{},
	"message_changed":   &MessageChangedEvent{},
	"message_deleted":   &MessageDeletedEvent{},
	"channel_join":      &ChannelJoinEvent{},
	"channel_leave":     &ChannelLeaveEvent{},
	"channel_topic":     &ChannelTopicEvent{},
	"channel_purpose":   &ChannelPurposeEvent{},
	"channel_name":      &ChannelNameEvent{},
	"channel_archive":   &ChannelArchiveEvent{},
	"channel_unarchive": &ChannelUnArchiveEvent{},
	"group_join":        &MessageEvent{},
	"group_leave":       &MessageEvent{},
	"group_topic":       &MessageEvent{},
	"group_purpose":     &MessageEvent{},
	"group_name":        &MessageEvent{},
	"group_archive":     &MessageEvent{},
	"group_unarchive":   &MessageEvent{},
	"file_share":        &MessageEvent{},
	"file_comment":      &MessageEvent{},
	"file_mention":      &MessageEvent{},
	"pinned_item":       &MessageEvent{},
	"unpinned_item":     &MessageEvent{},
}

func GetEventType(e *Event) interface{} {
	if EventType, ok := EventTypes[e.Type]; ok {
		if e.SubType == "" {
			return EventType
		} else if EventSubType, ok := EventSubTypes[e.SubType]; ok {
			return EventSubType
		} else {
			ErrorLog.Println("EventSubType not found:", e.SubType)
		}
	} else {
		ErrorLog.Println("EventType not found:", e.Type)
	}
	return nil
}

func UnmarshalRaw(raw *json.RawMessage, event interface{}) error {
	if err := json.Unmarshal(*raw, &event); err != nil {
		ErrorLog.Printf("error unmarshaling %T to %T:\n\t%+[1]v\n", raw, event)
		return err
	}
	return nil
}
