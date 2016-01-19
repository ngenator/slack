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

	return nil
}

func New(token string) *Slack {
	return &Slack{
		*NewStorage(),
		NewAPIClient(token),
		NewRealtimeClient(),
	}
}
