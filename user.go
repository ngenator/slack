package slack

// User type
type User struct {
	ID UserID `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Deleted bool `json:"deleted,omitempty"`

	Color string `json:"color,omitempty"`

	Profile *UserProfile `json:"profile,omitempty"`

	IsAdmin           bool `json:"is_admin,omitempty"`
	IsOwner           bool `json:"is_owner,omitempty"`
	IsPrimaryOwner    bool `json:"is_primary_owner,omitempty"`
	IsRestricted      bool `json:"is_restricted,omitempty"`
	IsUltraRestricted bool `json:"is_ultra_restricted,omitempty"`
	IsBot             bool `json:"is_bot,omitempty"`

	HasTwoFactorAuth bool   `json:"has_2fa,omitempty"`
	TwoFactorType    string `json:"two_factor_type,omitempty"`
	HasFiles         bool   `json:"has_files,omitempty"`
}

type UserProfile struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	RealName  string `json:"real_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Skype     string `json:"skype,omitempty"`
	Phone     string `json:"phone,omitempty"`
}
