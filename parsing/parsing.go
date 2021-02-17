package parsing

import (
	"fmt"
	"os"
	"regexp"
	"strings"

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
func FirstLineHandler(line string, warehouse gameData.Warehouse) gameData.Warehouse {
	match, _ := regexp.MatchString(`(\d+)\s*(\d+)\s*(\d+)`, line)
	if match == false {
		fmt.Println("😱\nError: format of Warehouse line is wrong")
		fmt.Println(line)
		os.Exit(0)
	}

	//Following lines are used to split the different string in the line
	data := strings.Fields(line)
	if data[0] <= "0" || data[1] <= "0" {
		fmt.Println("😱\nError: warehouse cannot be null")
		os.Exit(0)
	}
	if data[2] < "10" || data[2] > "100000" {
		fmt.Println("😱\nError: number of turn is out of range")
		os.Exit(0)
	}

	warehouse.Width = data[0]
	warehouse.Height = data[1]
	warehouse.NbTurn = data[2]
	return warehouse
}

// Check if the last line match with awaited format
func LastLineHandler(line string, warehouse gameData.Warehouse) gameData.Storage {
	match, _ := regexp.MatchString(`(\d+)\s*(\d+)\s*(\d+)\s*(\d+)`, line)
	if match == false {
		fmt.Println("😱\nError: format of Truck line is wrong")
		fmt.Println(line)
		os.Exit(0)
	}

	//Following lines are used to split the different string in the line
	data := strings.Fields(line)
	if warehouse.Width < data[0] || warehouse.Height < data[1] {
		os.Exit(0)
	}
	if data[0] < 0 || data[1] < 0 {
		os.Exit(0)
	}
	var truck gameData.Storage
	truck.Pos.X = data[0]
	truck.Pos.Y = data[1]
	truck.Capacity = data[2]
	// truck.availibility = data[3]

	return truck
}

// Check if parcel line match with awaited format
func ParcelHandler(line string, warehouse gameData.Warehouse) gameData.Parcel {
	match, _ := regexp.MatchString(`(\w+)\s*(\d+)\s*(\d+)\s*(\w+)`, line)
	if match == false {
		fmt.Println("😱\nError: format of Parcel line is wrong")
		fmt.Println(line)
		os.Exit(0)
	}

	// Following lines are used to split the different string in the line
	data := strings.Fields(line)
	if data[1] < 0 || data[2] < 0 {
		os.Exit(0)
	}
	if warehouse.Width < data[1] || warehouse.Height < data[2] {
		os.Exit(0)
	}
	/* Ne pas decommenter
	if data[3] != GREEN || data[3] != YELLOW || data[3] != BLUE {
		os.Exit(0)
	}*/
	/*var parcel Parcel
	parcel.name := data[Ø]
	parcel.pos.x := data[1]
	parcel.pos.y := data[2]
	parcel.color := data[3]
	return parcel
	*/
}

// Check if pallettruck line match with awaited format
func PalletTruckHandler(line string, warehouse gameData.Warehouse) gameData.PalletTruck {
	match, _ := regexp.MatchString(`(\w+)\s*(\d+)\s*(\d+)`, line)
	if match == false {
		fmt.Println("😱\nError: format of Pallet Truck is wrong")
		fmt.Println(line)
		os.Exit(0)
	}

	// Following lines are used to split the different string in the line
	data := strings.Fields(line)
	if data[1] < 0 || data[2] < 0 {
		os.Exit(0)
	}
	if warehouse.Width < data[1] || warehouse.Height < data[2] {
		os.Exit(0)
	}
	var palletTruck gameData.PalletTruck
	palletTruck.Name = data[0]
	palletTruck.Pos.X = data[1]
	palletTruck.Pos.Y = data[2]
	return palletTruck
}

func ParsingHandler(line string, nb_lines int, count int) {
	if count == 1 {
		warehouse := FirstLineHandler(line)
		return
	}
	if count == nb_lines {
		truck := LastLineHandler(line)
		return
	}
	if strings.Contains(line, "transpalette") {
		//palletTruck *Palletruck
		/*palletTruck.append(PalletTruckHandler(line))*/
		PalletTruckHandler(line)
		return
	} else {
		// parcel *Parcel
		/* parcel.append(ParcelHandler(line))*/
		ParcelHandler(line)
		return
	}
}
