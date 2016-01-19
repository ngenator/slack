package slack

import (
	"fmt"
	"strconv"
	"time"
)

type Timestamp struct {
	time.Time
}

func (t *Timestamp) MarshalJSON() ([]byte, error) {
	ts := t.Time.Unix()

	return []byte(fmt.Sprint(ts)), nil
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 0, 64)
	if err != nil {
		return err
	}

	t.Time = time.Unix(int64(ts), 0)

	return nil
}

// UserID is a reference to a User by User.ID
type UserID string

// ChannelID is a reference to a Channel by Channel.ID
type ChannelID string

// UniqueTimestamp is a unique (per-channel) timestamp
type UniqueTimestamp string
