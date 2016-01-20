package slack

type Slack struct {
	Storage
	API      *APIClient
	Realtime *RealtimeClient
}

func (s *Slack) ConnectRealtime() error {
	response, err := s.API.RTMStart()
	if err != nil {
		return err
	}

	if err := s.Realtime.Dial(response.URL, OriginURL); err != nil {
		return err
	}

	for _, u := range response.Users {
		s.Users[u.ID] = u
	}

	for _, c := range response.Channels {
		s.Channels[c.ID] = c
	}

	return nil
}

func New(token string) *Slack {
	return &Slack{
		*NewStorage(),
		NewAPIClient(token),
		NewRealtimeClient(),
	}
}
