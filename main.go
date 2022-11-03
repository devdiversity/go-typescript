package main

import (
	"fmt"

	"github.com/devdiversity/go-typescript/typescript"
)

func main() {
	fmt.Println("hello")
	typescript.TsInfo.Populate()
}
