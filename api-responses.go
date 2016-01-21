package slack

type Response struct {
	Ok    bool      `json:"ok"`
	Error *APIError `json:"error,omitempty"`
}

type RTMStartResponse struct {
	URL      string     `json:"url,omitempty"`
	Channels []*Channel `json:"channels,omitempty"`
	Users    []*User    `json:"users,omitempty"`
	// TODO: add the rest of the initial data
	Response
}

type UserResponse struct {
	User *User `json:"user,omitempty"`
	Response
}
type AuthTestResponse struct {
	UserName string `json:"user,omitempty"`
	UserID   UserID `json:"user_id,omitempty"`
	Response
}
