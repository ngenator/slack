package slack

// Description of something returned by the slack api
type Description struct {
	Value   string    `json:"string,omitempty"`
	Creator string    `json:"creator,omitempty"`
	LastSet Timestamp `json:"last_set,omitempty"`
}

type Timestamp uint64
