package parsing

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	gameData "github.com/LaurentColoma/M-BDX-002/gameData"
)

func NbArgsHandler() bool {
	if len(os.Args) < 2 {
		fmt.Println("😱\nError: wrong number of arguments")
		os.Exit(0)
	}
	return true
}

func OpenFileHandler(err error) bool {
	if err != nil {
		fmt.Println("😱\nError: file not supported")
		os.Exit(0)
	}
	return true
}

func FirstLineHandler(line string, warehouse *gameData.Warehouse) {
	match, _ := regexp.MatchString(`(\d+)\s*(\d+)\s*(\d+)`, line)
	if match == false {
		fmt.Println("😱\nError: format of Warehouse line is wrong")
		os.Exit(0)
	}

	var data = strings.Fields(line)
	var w, _ = strconv.Atoi(data[0])
	var h, _ = strconv.Atoi(data[1])
	var n, _ = strconv.Atoi(data[2])
	if w <= 0 || h <= 0 {
		fmt.Println("😱\nError: warehouse cannot be null")
		os.Exit(0)
	}
	if n < 10 || n > 100000 {
		fmt.Println("😱\nError: number of turn is out of range")
		os.Exit(0)
	}
	warehouse.Width = w
	warehouse.Height = h
	warehouse.NbTurn = n
}

func LastLineHandler(line string, warehouse *gameData.Warehouse) gameData.Storage {
	match, _ := regexp.MatchString(`(\d+)\s*(\d+)\s*(\d+)\s*(\d+)`, line)
	if match == false {
		fmt.Println("😱\nError: format of Truck line is wrong")
		os.Exit(0)
	}

	var data = strings.Fields(line)
	var x, _ = strconv.Atoi(data[0])
	var y, _ = strconv.Atoi(data[1])
	if warehouse.Width <= x || warehouse.Height <= y {
		fmt.Println("😱\nError: Out of range")
		os.Exit(0)
	}
	if x < 0 || y < 0 {
		fmt.Println("😱\nError: Out of range")
		os.Exit(0)
	}
	var truck gameData.Storage
	truck.Pos.X = x
	truck.Pos.Y = y
	truck.Capacity, _ = strconv.Atoi(data[2])
	truck.Upturn, _ = strconv.Atoi(data[3])
	return truck
}

func ParcelHandler(line string, warehouse *gameData.Warehouse) gameData.Parcel {
	match, _ := regexp.MatchString(`(\w+)\s*(\d+)\s*(\d+)\s*(\w+)`, line)
	if match == false {
		fmt.Println("😱\nError: format of Parcel line is wrong")
		os.Exit(0)
	}

	data := strings.Fields(line)
	weight := 0
	var x, _ = strconv.Atoi(data[1])
	var y, _ = strconv.Atoi(data[2])
	if x < 0 || y < 0 {
		fmt.Println("😱\nError: Out of range")
		os.Exit(0)
	}
	if warehouse.Width <= x || warehouse.Height <= y {
		fmt.Println("😱\nError: Out of range")
		os.Exit(0)
	}
	if strings.ToUpper(data[3]) == "YELLOW" {
		weight = 100
	} else if strings.ToUpper(data[3]) == "GREEN" {
		weight = 200
	} else if strings.ToUpper(data[3]) == "BLUE" {
		weight = 500
	} else {
		fmt.Println("Error: wrong color")
		os.Exit(0)
	}
	var parcel gameData.Parcel
	parcel.Name = data[0]
	parcel.Pos.X = x
	parcel.Pos.Y = y
	parcel.Weight = weight
	parcel.Aimed = false
	parcel.Dead = false
	return parcel
}

func PalletTruckHandler(line string, warehouse *gameData.Warehouse) gameData.PalletTruck {
	match, _ := regexp.MatchString(`(\w+)\s*(\d+)\s*(\d+)`, line)
	if match == false {
		fmt.Println("😱\nError: format of Pallet Truck is wrong")
		os.Exit(0)
	}

	data := strings.Fields(line)
	var x, _ = strconv.Atoi(data[1])
	var y, _ = strconv.Atoi(data[2])
	if x < 0 || y < 0 {
		fmt.Println("😱\nError: Out of range")
		os.Exit(0)
	}
	if warehouse.Width <= x || warehouse.Height <= y {
		fmt.Println("😱\nError: Out of range")
		os.Exit(0)
	}
	var palletTruck gameData.PalletTruck
	palletTruck.Name = data[0]
	palletTruck.Pos.X = x
	palletTruck.Pos.Y = y
	return palletTruck
}

func ParsingHandler(line string, nb_lines int, count int, warehouse *gameData.Warehouse) {
	if count == 1 {
		FirstLineHandler(line, warehouse)
		return
	}
	if count == nb_lines {
		warehouse.Truck = LastLineHandler(line, warehouse)
		return
	}
	testLine := strings.Fields(line)
	if len(testLine) == 3 {
		warehouse.PalletTrucks = append(warehouse.PalletTrucks, PalletTruckHandler(line, warehouse))
		return
	}
	if len(testLine) == 4 {
		warehouse.Parcels = append(warehouse.Parcels, ParcelHandler(line, warehouse))
		return
	}
}
