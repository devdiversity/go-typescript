package main

import (
	"fmt"
	"typescript/typescript"
)

func main() {
	var tsInfoData = typescript.TSInfo{}
	var tsSoucesData = typescript.TSSouces{}
	tsInfoData.Populate()
	tsSoucesData.New()
	tsSoucesData.Populate(tsInfoData)

	for _, v := range tsSoucesData.Errors {
		fmt.Println(v)
	}
	for p, _ := range tsSoucesData.Pakages {
		fmt.Printf("\nexport namespace %s {\n", p)
		for _, v1 := range tsSoucesData.Pakages[p].Structs {
			fmt.Println(v1)
		}
		for _, v1 := range tsSoucesData.Pakages[p].Types {
			fmt.Println(v1)
		}
		fmt.Printf("}\n// end namespace %s \n\n", p)
	}
}
