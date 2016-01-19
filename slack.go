package slack

type Slack struct {
	Storage
	API      *APIClient
	Realtime *RealtimeClient
}

func New(token string) *Slack {
	return &Slack{
		*NewStorage(),
		NewAPIClient(token),
		NewRealtimeClient(),
	}
}
