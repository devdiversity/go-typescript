package tstest

import (
	"time"

	testexternal "github.com/devdiversity/go-typescript/tsTestExternal"
)

type httpSessions struct {
	ID         int64
	Key        string
	Data       string
	CreatedOn  time.Time
	ModifiedOn time.Time
	ExpiresOn  time.Time
}

type TestTest struct {
	Info  string `json:"info"`
	Count string
}

// Typescript: interface
type HttpSessions struct {
	ID         int64
	Key        string
	Data       string
	CreatedOn  time.Time
	ModifiedOn time.Time
	ExpiresOn  time.Time
}

// Typescript: type
type TestType []int

// Typescript: type
type TestTypeMap map[string]map[int]string

// Typescript: type
type TestTypeTime time.Time

// Typescript: interface
type TestStruct1 struct {
	CreatedOn time.Time    `json:"CreatedOn"`
	TestT     TestType     `json:"TestT"`
	Session   HttpSessions `json:"Session"`
	ID        int64        `json:"id"`
	Key       []string     `json:"key"`
	Data      *string      `json:"data"`
	DataPTR   *[]string
	UserPswd  testexternal.UserRegisterResponse `json:"newpassword"`

	ModifiedOn time.Time `json:"ModifiedOn"`
	ExpiresOn  time.Time `json:"-"`

	MapsArray     []map[string]time.Time
	Maps          map[string]time.Time         `json:"Maps"`
	MapsNested    map[string]map[int]string    `json:"MapsNested"`
	MapsNestedPTR map[string]map[int]*[]string `json:"MapsNestedPtr"`
	TestTest      TestTest                     `json:"testTest"`
	TestType      TestType                     `json:"testType"`
	TestTypeMap   TestTypeMap                  `json:"TestTypeMap"`
}

// Typescript: type TestTypeStruct
type TestTypeStruct TestStruct1
