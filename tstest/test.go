package tstest

import (
	"time"
	"typescript/tsTestExternal"
)

// Typescript: TSDeclaration= Nullable<T> = T | null;
// Typescript: TSDeclaration= Record<K extends string | number | symbol, T> = { [P in K]: T; }

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

// Typescript: TStype=  MyType = number

// Typescript: type
type TestType []int

// Typescript: type
type TestTypeMap map[string]map[int]string

// Typescript: type=Date
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
	EnumTest      Direction                            `json:"direction"`
	EnumSeason    Season                               `json:"season"`
}

// Typescript: type
type TestTypeStruct TestStruct1

type Season string

// Typescript: const
const Timeout = 1000

// Typescript: const
const (
	Uno   string = "uno"
	Cento int    = 100
)

// Typescript: enum=Season
const (
	Summer Season = "summer"
	Autumn        = "autumn"
	Winter        = "winter"
	Spring        = "spring"
)

// Typescript: enum=Test
const (
	A int = iota
	B
	C
	D
)

type Direction int

// Typescript: enum=Direction
const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}
