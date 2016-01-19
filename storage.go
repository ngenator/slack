package slack

import (
	"fmt"
)

// Storage stores information about the current state
type Storage struct {
	Channels map[ChannelID]*Channel
	Users    map[UserID]*User
}

// GetUser gets a user from a UserID
func (s *Storage) GetUser(id UserID) (*User, error) {
	if user, ok := s.Users[id]; ok {
		return user, nil
	}
	return nil, fmt.Errorf("User not found: %s", id)
}

// GetChannel gets a Channel from a ChannelID
func (s *Storage) GetChannel(id ChannelID) (*Channel, error) {
	if channel, ok := s.Channels[id]; ok {
		return channel, nil
	}
	return nil, fmt.Errorf("Channel not found: %s", id)
}

// NewStorage creates a Storage instance for holding Channels and Users
func NewStorage() *Storage {
	return &Storage{
		make(map[ChannelID]*Channel),
		make(map[UserID]*User),
	}
}
