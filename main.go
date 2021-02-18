package main

import (
	"bufio"
	"fmt"
	"os"

	gameData "github.com/LaurentColoma/M-BDX-002/gameData"
	Loop "github.com/LaurentColoma/M-BDX-002/gameLoop"
	parser "github.com/LaurentColoma/M-BDX-002/parsing"
	pathFinding "github.com/LaurentColoma/M-BDX-002/pathFinding"
)

func main() {
	parser.NbArgsHandler()

	content, err := os.Open(os.Args[1])
	parser.OpenFileHandler(err)
	nb_lines := 0
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		nb_lines += 1
	}
	content.Close()

	file, err := os.Open(os.Args[1])
	parser.OpenFileHandler(err)
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
	var m = pathFinding.MapFrom(&warehouse, 1, 2)
	pathFinding.GetRoute(m, warehouse.Width, warehouse.Height, 2, 0)
	if Loop.GameLoop(warehouse) == 1 {
		fmt.Println("🤩\nAll according to our plan")
	} else {
		fmt.Println("😏\nTurn elapsed")
	}
	return
}
