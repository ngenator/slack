package slack

import (
	"fmt"
)

// TODO: add sqlite3 db for this
// Storage stores information about the current state
type Storage struct {
	Channels map[ChannelID]*Channel
	Users    map[UserID]*User
}

// FindUser gets a user from a UserID if it's been stored
func (s *Storage) FindUser(id UserID) (*User, error) {
	if user, ok := s.Users[id]; ok {
		return user, nil
	}
	return nil, fmt.Errorf("User not found: %s", id)
}

// FindChannel gets a Channel from a ChannelID if it's been stored
func (s *Storage) FindChannel(id ChannelID) (*Channel, error) {
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
