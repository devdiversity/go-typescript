package main

import (
	"errors"
	"fmt"
	"os"
	"typescript/typescript"

	"github.com/northwood-labs/golang-utils/exiterrorf"
)

func main() {
	var tsInfoData = typescript.TSInfo{}
	var tsSoucesData = typescript.TSSouces{}
	tsInfoData.Populate()
	tsSoucesData.Populate(tsInfoData)

	if len(tsSoucesData.Errors) != 0 {
		err := ""
		for _, v := range tsSoucesData.Errors {
			err += fmt.Sprintln(v)
		}
		exiterrorf.ExitErrorf(errors.New(fmt.Sprintf("Some errors...\n %s", err)))
	}

	tsSource := ""

	tsSource += fmt.Sprintln("\n// Global Declarations ")
	for p := range tsSoucesData.Pakages {
		for _, v1 := range tsSoucesData.Pakages[p].GTypes {
			tsSource += fmt.Sprintln(v1)
		}
	}

	for p := range tsSoucesData.Pakages {
		tsSource += fmt.Sprintf("\n//\n// namespace %s\n//\n", p)
		tsSource += fmt.Sprintf("\nexport namespace %s {\n", p)
		for _, v1 := range tsSoucesData.Pakages[p].Structs {
			tsSource += fmt.Sprintln(v1)
		}
		for _, v1 := range tsSoucesData.Pakages[p].Types {
			tsSource += fmt.Sprintln(v1)
		}
		for _, v1 := range tsSoucesData.Pakages[p].Enums {
			tsSource += fmt.Sprintln(v1)
		}
		for _, v1 := range tsSoucesData.Pakages[p].Consts {
			tsSource += fmt.Sprintln(v1)
		}
		tsSource += fmt.Sprintf("}\n\n")
	}

	err := os.WriteFile("test.ts", []byte(tsSource), 0644)
	if err != nil {
		exiterrorf.ExitErrorf(errors.New(fmt.Sprintf("error riting file\n %s", err.Error())))
	}

}
