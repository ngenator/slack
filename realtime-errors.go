package slack

import "fmt"

type EventError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"msg,omitempty"`
}

func (e *EventError) String() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}