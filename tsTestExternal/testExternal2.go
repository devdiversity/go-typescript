package tsTestExternal

import "typescript/moduleExt"

// test
type UserRegisterResponse2 struct {
	Token        string `json:"token" `
	User         string `json:"user,omitempty"`
	c            chan string
	empty        EmptyStruct             `json:"empty" `
	TestExternal moduleExt.ModuleExtTest `json:"testdep" `
}

type EmptyStruct struct {
	test  string
	test2 string
}
