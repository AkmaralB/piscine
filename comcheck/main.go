package main

import (
	"fmt"
	"os"
)

func main() {
	argum := os.Args[1:]
	count := 0
	for _, v := range argum {
		if v == "01" || v == "galaxy" || v == "galaxy 01" {
			count++
		}
	}
	if count > 0 {
		fmt.Println("Alert!!!")
	}
}
