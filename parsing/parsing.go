package parsing

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"strconv"

	gameData "github.com/LaurentColoma/M-BDX-002/gameData"
)

// Check if the number of arguments is correct
func NbArgsHandler() bool {
	if len(os.Args) < 2 {
		fmt.Println("😱\nError: wrong number of arguments")
		os.Exit(0)
	}
	return true
}

// Check if the file can be opened
func OpenFileHandler(err error) bool {
	if err != nil {
		fmt.Println("😱\nError: file not supported")
		os.Exit(0)
	}
	return true
}

// Check if the first line match with awaited format
func FirstLineHandler(line string, warehouse *gameData.Warehouse) {
	match, _ := regexp.MatchString(`(\d+)\s*(\d+)\s*(\d+)`, line)
	if match == false {
		fmt.Println("😱\nError: format of Warehouse line is wrong")
		fmt.Println(line)
		os.Exit(0)
	}

	//Following lines are used to split the different string in the line
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

// Check if the last line match with awaited format
func LastLineHandler(line string, warehouse *gameData.Warehouse) gameData.Storage {
	match, _ := regexp.MatchString(`(\d+)\s*(\d+)\s*(\d+)\s*(\d+)`, line)
	if match == false {
		fmt.Println("😱\nError: format of Truck line is wrong")
		fmt.Println(line)
		os.Exit(0)
	}

	//Following lines are used to split the different string in the line
	var data = strings.Fields(line)
	var x, _ = strconv.Atoi(data[0])
	var y, _ = strconv.Atoi(data[1])
	if warehouse.Width <= x || warehouse.Height <= y {
		os.Exit(0)
	}
	if x < 0 || y < 0 {
		os.Exit(0)
	}
	var truck gameData.Storage
	truck.Pos.X = x
	truck.Pos.Y = y
	truck.Capacity, _ = strconv.Atoi(data[2])
	// truck.availibility = data[3]

	return truck
}

// Check if parcel line match with awaited format
func ParcelHandler(line string, warehouse *gameData.Warehouse) gameData.Parcel {
	match, _ := regexp.MatchString(`(\w+)\s*(\d+)\s*(\d+)\s*(\w+)`, line)
	if match == false {
		fmt.Println("😱\nError: format of Parcel line is wrong")
		fmt.Println(line)
		os.Exit(0)
	}

	// Following lines are used to split the different string in the line
	data := strings.Fields(line)
	var x, _ = strconv.Atoi(data[1])
	var y, _ = strconv.Atoi(data[2])
	if x < 0 || y < 0 {
		os.Exit(0)
	}
	if warehouse.Width <= x || warehouse.Height <= y {
		os.Exit(0)
	}
	/* Ne pas decommenter
	if data[3] != GREEN || data[3] != YELLOW || data[3] != BLUE {
		os.Exit(0)
	}*/
	var parcel gameData.Parcel
	parcel.Name = data[0]
	parcel.Pos.X = x
	parcel.Pos.Y = y
	parcel.Color = data[3]
	return parcel
}

// Check if pallettruck line match with awaited format
func PalletTruckHandler(line string, warehouse *gameData.Warehouse) gameData.PalletTruck {
	match, _ := regexp.MatchString(`(\w+)\s*(\d+)\s*(\d+)`, line)
	if match == false {
		fmt.Println("😱\nError: format of Pallet Truck is wrong")
		fmt.Println(line)
		os.Exit(0)
	}

	// Following lines are used to split the different string in the line
	data := strings.Fields(line)
	var x, _ = strconv.Atoi(data[1])
	var y, _ = strconv.Atoi(data[2])
	if x < 0 || y < 0 {
		os.Exit(0)
	}
	if warehouse.Width <= x || warehouse.Height <= y {
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
	if strings.Contains(line, "transpalette") {
		//palletTruck *Palletruck
		/*palletTruck.append(PalletTruckHandler(line))*/
		PalletTruckHandler(line, warehouse)
		return
	} else {
		// parcel *Parcel
		/* parcel.append(ParcelHandler(line))*/
		ParcelHandler(line, warehouse)
		return
	}
}
