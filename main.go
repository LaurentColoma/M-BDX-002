package main

import (
	"bufio"
	"fmt"
	"os"

	"encoding/json"

	gameData "github.com/LaurentColoma/M-BDX-002/gameData"
	parser "github.com/LaurentColoma/M-BDX-002/parsing"
)

func main() {
	if parser.NbArgsHandler() == false {
		os.Exit(0)
	}

	content, err := os.Open(os.Args[1])
	if parser.OpenFileHandler(err) == false {
		os.Exit(0)
	}
	nb_lines := 0
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		nb_lines += 1
	}
	content.Close()

	file, err := os.Open(os.Args[1])
	if parser.OpenFileHandler(err) == false {
		os.Exit(0)
	}
	count := 0
	scanner2 := bufio.NewScanner(file)
	var warehouse gameData.Warehouse
	for scanner2.Scan() {
		count += 1
		parser.ParsingHandler(scanner2.Text(), nb_lines, count, &warehouse)
	}
	file.Close()
	wh, _ := json.MarshalIndent(warehouse, "", " ")
	fmt.Println(string(wh))
	//start game loop here
	return
}
