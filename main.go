package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: wrong number of arguments")
		return
	}

	content, error := ioutil.ReadFile(os.Args[1])

	if error != nil {
		fmt.Println("Error: file not supported")
	}

	fmt.Println(string(content))
}
