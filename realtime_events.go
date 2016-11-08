package slack

import "encoding/json"

// TODO: create event structs for each event
var EventTypes = map[string]interface{}{
	"hello": &HelloEvent{},

	"pong": &PongEvent{},

	"message": &MessageEvent{},

	"user_typing": &UserTypingEvent{},

	"channel_marked":          &ChannelMarkedEvent{},
	"channel_created":         &ChannelCreatedEvent{},
	"channel_joined":          &ChannelJoinedEvent{},
	"channel_left":            &ChannelLeftEvent{},
	"channel_deleted":         &ChannelDeletedEvent{},
	"channel_rename":          &ChannelRenameEvent{},
	"channel_archive":         &ChannelArchiveEvent{},
	"channel_unarchive":       &ChannelUnarchiveEvent{},
	"channel_history_changed": &Event{},

	"dnd_updated":      &Event{},
	"dnd_updated_user": &Event{},

	"im_marked":          &ImMarkedEvent{},
	"im_created":         &ImCreatedEvent{},
	"im_open":            &ImOpenEvent{},
	"im_close":           &ImCloseEvent{},
	"im_history_changed": &ImHistoryChangedEvent{},

	"group_joined":          &Event{},
	"group_left":            &Event{},
	"group_open":            &Event{},
	"group_close":           &Event{},
	"group_archive":         &Event{},
	"group_unarchive":       &Event{},
	"group_rename":          &Event{},
	"group_marked":          &GroupMarkedEvent{},
	"group_history_changed": &Event{},

	"file_created":         &Event{},
	"file_shared":          &Event{},
	"file_unshared":        &Event{},
	"file_public":          &Event{},
	"file_private":         &Event{},
	"file_change":          &Event{},
	"file_deleted":         &Event{},
	"file_comment_added":   &Event{},
	"file_comment_edited":  &Event{},
	"file_comment_deleted": &Event{},

	"pin_added":   &Event{},
	"pin_removed": &Event{},

	"presence_change":        &PresenceChangeEvent{},
	"manual_presence_change": &ManualPresenceChangeEvent{},

	"pref_change": &Event{},

	"user_change": &Event{},

	"team_join": &Event{},

	"star_added":   &Event{},
	"star_removed": &Event{},

	"reaction_added":   &Event{},
	"reaction_removed": &Event{},

	"emoji_changed": &Event{},

	"commands_changed": &Event{},

	"team_plan_change":     &Event{},
	"team_pref_change":     &Event{},
	"team_rename":          &Event{},
	"team_domain_changed":  &Event{},
	"email_domain_changed": &Event{},
	"team_profile_change":  &Event{},
	"team_profile_delete":  &Event{},
	"team_profile_reorder": &Event{},

	"bot_added":   &Event{},
	"bot_changed": &Event{},

	"accounts_changed": &Event{},

	"subteam_created":      &Event{},
	"subteam_updated":      &Event{},
	"subteam_self_added":   &Event{},
	"subteam_self_removed": &Event{},
}

func GetEventType(e *Event) interface{} {
	if EventType, ok := EventTypes[e.Type]; ok {
		if e.Type == "message" && e.SubType != "" {
			if MessageEventType, ok := MessageEventTypes[e.SubType]; ok {
				return MessageEventType
			}
			ErrorLog.Println("MessageEventType not found:", e.SubType)
		}
		return EventType
	}
	ErrorLog.Println("EventType not found:", e.Type)
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

type HelloEvent struct{}

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
	ChannelID           ChannelID       `json:"channel,omitempty"`
	Timestamp           UniqueTimestamp `json:"ts,omitempty"`
	UnreadCount         int             `json:"unread_count,omitempty"`
	UnreadCountDisplay  int             `json:"unread_count_display,omitempty"`
	NumMentions         int             `json:"num_mentions,omitempty"`
	NumMentionsDisplay  int             `json:"num_mentions_display,omitempty"`
	MentionCount        int             `json:"mention_count,omitempty"`
	MentionCountDisplay int             `json:"mention_count_display,omitempty"`
}

type ChannelCreatedEvent struct{}
type ChannelJoinedEvent struct{}
type ChannelLeftEvent struct{}
type ChannelDeletedEvent struct{}
type ChannelRenameEvent struct{}
type ChannelArchiveEvent struct{}
type ChannelUnarchiveEvent struct{}
type ChannelHistoryChangedEvent struct{}

// ##################################################

type DoNotDisturbUpdatedEvent struct{}
type DoNotDisturbUpdatedUserEvent struct{}

// ##################################################

type ReactionAddedEvent struct{}
type ReactionRemovedEvent struct{}

// ##################################################

type ImMarkedEvent struct {
	ChannelID           ChannelID       `json:"channel,omitempty"`
	Timestamp           UniqueTimestamp `json:"ts,omitempty"`
	DMCount             int             `json:"dm_count,omitempty"`
	UnreadCountDisplay  int             `json:"unread_count_display,omitempty"`
	NumMentionsDisplay  int             `json:"num_mentions_display,omitempty"`
	MentionCountDisplay int             `json:"mention_count_display,omitempty"`
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
	Latest         UniqueTimestamp `json:"latest,omitempty"`
	Timestamp      UniqueTimestamp `json:"ts,omitempty"`
	EventTimestamp UniqueTimestamp `json:"event_ts,omitempty"`
}

// ##################################################

type GroupMarkedEvent struct{}

// ##################################################

type FileCreatedEvent struct{}
type FileSharedEvent struct{}
