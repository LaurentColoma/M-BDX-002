package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	gameData "github.com/LaurentColoma/M-BDX-002/gameData"
	parser "github.com/LaurentColoma/M-BDX-002/parsing"
	pathFinding "github.com/LaurentColoma/M-BDX-002/pathFinding"
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
	if warehouse.Parcels == nil {
		fmt.Println("😱\nError: No Parcel found")
		os.Exit(0)
	}
	if warehouse.PalletTrucks == nil {
		fmt.Println("😱\nError: No Pallet Truck found")
		os.Exit(0)
	}
	file.Close()
	wh, _ := json.MarshalIndent(warehouse, "", " ")
	fmt.Println(string(wh))
	//start game loop here
	var m = pathFinding.MapFrom(&warehouse, 1, 1)
	var p = pathFinding.GetRoute(m, warehouse.Width, warehouse.Height, 3, 3)
	fmt.Println(m, p)
	// if loop.GameLoop(warehouse) == 1 {
	// 	fmt.println("smiley cool")
	// } else {
	// 	fmt.println("smiley")
	// }
	return
}
