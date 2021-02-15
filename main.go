package main

import (
	"bufio"
	"fmt"
	"os"
)

func NbArgsHandler() bool {
	if len(os.Args) < 2 {
		fmt.Println("Error: wrong number of arguments")
		return false
	}
	return true
}

func OpenFileHandler(err error) bool {
	if err != nil {
		fmt.Println("Error: file not supported")
		return false
	}
	return true
}

func FirstLineHandler(line string) bool {
	return true
}

func ParsingHandler(line string, count int) bool {
	if count == 0 {
		return FirstLineHandler(line)
	}
	return true
}

func main() {
	if NbArgsHandler() == false {
		return
	}

	content, err := os.Open(os.Args[1])
	if OpenFileHandler(err) == false {
		return
	}
	defer content.Close()

	count := 0
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		if ParsingHandler(scanner.Text(), count) == false {
			return
		}
	}

	return
}
