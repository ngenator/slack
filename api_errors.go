package slack

import (
	"encoding/json"
	"fmt"
)

type APIError string

func (e *APIError) MarshalJSON() ([]byte, error) {
	return []byte(*e), nil
}

func (e *APIError) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if message, ok := APIErrorMessages[s]; ok {
		*e = APIError(fmt.Sprintf("%s: %s", s, message))
	} else {
		*e = APIError(s)
	}
	return nil
}

func (e *APIError) String() string {
	return string(*e)
}

// TODO: add the rest of the error messages
var APIErrorMessages = map[string]string{
	// common
	"not_authed":       "No authentication token provided.",
	"invalid_auth":     "Invalid authentication token.",
	"account_inactive": "Authentication token is for a deleted user or team.",
	// chat.postMessage
	"channel_not_found": "Value passed for `channel` was invalid.",
	"not_in_channel":    "Cannot post user messages to a channel they are not in.",
	"is_archived":       "Channel has been archived.",
	"msg_too_long":      "Message text is too long.",
	"no_text":           "No message text provided.",
	"rate_limited":      "Application has posted too many messages, read the Rate Limit documentation for more information.",
	// reactions.add
	"bad_timestamp":          "Value passed for `timestamp` was invalid.",
	"file_not_found":         "File specified by `file` does not exist.",
	"file_comment_not_found": "File comment specified by `file_comment` does not exist.",
	"message_not_found":      "Message specified by `channel` and `timestamp` does not exist.",
	"no_item_specified":      "`file`, `file_comment`, or combination of `channel` and `timestamp` was not specified.",
	"invalid_name":           "Value passed for `name` was invalid.",
	"already_reacted":        "The specified item already has the user/reaction combination.",
	"too_many_emoji":         "The limit for distinct reactions (i.e emoji) on the item has been reached.",
	"too_many_reactions":     "The limit for reactions a person may add to the item has been reached.",
}
