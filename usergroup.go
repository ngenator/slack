package slack

type UserGroup struct {
	ID          string `json:"id,omitempty"`
	TeamID      string `json:"team_id,omitempty"`
	IsUserGroup bool   `json:"is_usergroup,omitempty"`
	IsSubTeam   bool   `json:"is_subteam,omitempty"`
	IsExternal  bool   `json:"is_external,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"name,omitempty"`
	Handle      string `json:"handle,omitempty"`
	Preferences struct {
		ChannelIDs []string `json:"channels,omitempty"`
		GroupIDs   []string `json:"groups,omitempty"`
	} `json:"prefs,omitempty"`
	AutoType  string   `json:"auto_type,omitempty"`
	UserCount int      `json:"user_count,omitempty"`
	UserIDs   []string `json:"users,omitempty"`
}
