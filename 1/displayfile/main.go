package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arr := os.Args[1:]
	lens := 0
	for range arr {
		lens++
	}

	if lens == 0 || os.Args[1] != "quest8.txt" {
		fmt.Println("File name missing")
	} else if lens > 1 {
		fmt.Println("Too many arguments")
	} else {
		file, err := ioutil.ReadFile(os.Args[1])

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(string(file))
		}
	}

}
