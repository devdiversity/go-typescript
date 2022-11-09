package tstest

import (
	"time"
	"typescript/tsTestExternal"
)

type HttpSessions struct {
	ID         int64     `json:"id"`
	Key        string    `json:"key"`
	Data       string    `json:"data"`
	CreatedOn  time.Time `json:"created" ts:"type=Date"`
	ModifiedOn time.Time `json:"modified" ts:"type=Date"`
	ExpiresOn  time.Time `json:"expire" ts:"type=Date"`
}

type TestTest struct {
	tsTestExternal.UserRegisterResponse `ts:"expand"`
	Info                                string `json:"info"`
	Count                               string
	TypeName                            string    `json:"typename" ts:"type=MyType"`
	CreatedOn                           time.Time `json:"created" ts:"type=Date"`
}

// Typescript: type
type TestType []int

// Typescript: type
type TestTypeMap map[string]map[int]string

// Typescript: type
type TestTypeTime time.Time

// Typescript: interface
type TestStruct1 struct {
	CreatedOn time.Time    `json:"created" ts:"type=Date"`
	TestT     TestType     `json:"TestT"`
	Session   HttpSessions `json:"Session"`
	ID        int64        `json:"id"`
	Key       []string     `json:"key"`
	Data      *string      `json:"data"`
	DataPTR   *[]string
	UserPswd  tsTestExternal.UserRegisterResponse `json:"newpassword"`

	ModifiedOn time.Time `json:"modified" ts:"type=Date"`
	ExpiresOn  time.Time `json:"-"`

	MapsArray     []map[string]time.Time
	Maps          map[string]time.Time                 `json:"maps" ts:"type=Date"`
	MapsNested    map[string]map[int]string            `json:"MapsNested"`
	MapsNestedPTR map[string]map[int]*[]string         `json:"MapsNestedPtr"`
	TestTest      TestTest                             `json:"testTest"`
	TestType      TestType                             `json:"testType"`
	TestTypeMap   TestTypeMap                          `json:"TestTypeMap"`
	TestDep       tsTestExternal.UserRegisterResponse2 `json:"testdep"`
}

// Typescript: type TestTypeStruct
type TestTypeStruct TestStruct1
