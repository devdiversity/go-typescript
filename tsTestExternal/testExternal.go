package tsTestExternal

// test
type UserRegisterResponse struct {
	Token string `json:"token" `
	User  string `json:"user,omitempty"`
	c     chan string
	empty EmptyStruct
}
