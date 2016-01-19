package slack

type Slack struct {
	Channels map[ChannelID]*Channel
	Users    map[UserID]*User
}

func (s *Slack) FindUser(id UserID) *User {
	if user, ok := s.Users[id]; ok {
		return user
	} else {
		return &User{ID: id, Name: "Unknown"}
	}
}

func (s *Slack) FindChannel(id ChannelID) *Channel {
	if channel, ok := s.Channels[id]; ok {
		return channel
	} else if id[0] == 'D' {
		return &Channel{ID: id, Name: "Direct Message"}
	} else {
		return &Channel{ID: id, Name: "Unknown"}
	}
}

// Creates a new Slack instance for holding Channels and Users
func NewSlack() *Slack {
	return &Slack{
		make(map[ChannelID]*Channel),
		make(map[UserID]*User),
	}
}
