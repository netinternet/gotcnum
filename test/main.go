package main

import (
	"fmt"

	"github.com/netinternet/gotcnum"
)

func main() {
	if valid := gotcnum.Tcnum("10148343778"); valid {
		fmt.Println("TC Number Is Valid")
	} else {
		fmt.Println("TC Number Is Not Valid")
	}
}
