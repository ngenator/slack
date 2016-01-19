package slack

import (
	"encoding/json"
	"fmt"
)

var EventTypes = map[string]interface{}{
	"hello": &HelloEvent{},
	"pong":  &PongEvent{},

	"message": &MessageEvent{},

	"user_typing": &UserTypingEvent{},

	"presence_change":        &PresenceChangeEvent{},
	"manual_presence_change": &ManualPresenceChangeEvent{},

	"channel_marked":    &ChannelMarkedEvent{},
	"channel_created":   &ChannelCreatedEvent{},
	"channel_joined":    &ChannelJoinedEvent{},
	"channel_left":      &ChannelLeftEvent{},
	"channel_deleted":   &ChannelDeletedEvent{},
	"channel_rename":    &ChannelRenameEvent{},
	"channel_archive":   &ChannelArchiveEvent{},
	"channel_unarchive": &ChannelUnarchiveEvent{},

	"group_marked": &GroupMarkedEvent{},

	"im_marked":          &ImMarkedEvent{},
	"im_created":         &ImCreatedEvent{},
	"im_open":            &ImOpenEvent{},
	"im_close":           &ImCloseEvent{},
	"im_history_changed": &ImHistoryChangedEvent{},
}

func GetEventType(e *Event) interface{} {
	if EventType, ok := EventTypes[e.Type]; ok {
		switch e.Type {
		case "message":
			if e.SubType != "" {
				if MessageEventType, ok := MessageEventTypes[e.SubType]; ok {
					return MessageEventType
				} else {
					ErrorLog.Println("MessageEventType not found:", e.SubType)
				}
			}
			return EventType
		default:
			return EventType
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

// ##################################################

type PresenceChangeEvent struct {
	Presence string `json:"presence,omitempty"`
	UserID   UserID `json:"user,omitempty"`
}

type ManualPresenceChangeEvent struct{}

type UserTypingEvent struct {
	ChannelID ChannelID `json:"channel,omitempty"`
	UserID    UserID    `json:"user,omitempty"`
}

// ##################################################

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

type ChannelCreatedEvent struct {
}

type ChannelJoinedEvent struct {
}

type ChannelLeftEvent struct {
}

type ChannelDeletedEvent struct {
}

type ChannelRenameEvent struct {
}

type ChannelArchiveEvent struct {
}

type ChannelUnarchiveEvent struct {
}

type ChannelHistoryChangedEvent struct {
}

// ##################################################

type DoNotDisturbUpdatedEvent struct{}

type DoNotDisturbUpdatedUserEvent struct{}

// ##################################################

type ReactionAddedEvent struct{}

type ReactionRemovedEvent struct{}

// ##################################################

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

type ImHistoryChangedEvent struct {
	Latest         SlackTimestamp `json:"latest,omitempty"`
	Timestamp      SlackTimestamp `json:"ts,omitempty"`
	EventTimestamp SlackTimestamp `json:"event_ts,omitempty"`
}

// ##################################################

type GroupMarkedEvent struct {
}
