package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	//"strings"
)

// Check if the number of arguments is correct
func NbArgsHandler() bool {
	if len(os.Args) < 2 {
		fmt.Println("Error: wrong number of arguments\n😱")
		os.Exit(0)
	}
	return true
}

// Check if the file can be opened
func OpenFileHandler(err error) bool {
	if err != nil {
		fmt.Println("Error: file not supported\n😱")
		os.Exit(0)
	}
	return true
}

// Check if the first line match with awaited format
func FirstLineHandler(line string /*, warehouse Warehouse*/) /*, Warehouse */ {
	match, _ := regexp.MatchString(`(\d+)\s*(\d+)\s*(\d+)`, line)
	if match == false {
		fmt.Println("Error: format of Warehouse line is wrong\n😱")
		fmt.Println(line)
		os.Exit(0)
	}

	//Following lines are used to split the different string in the line
	data := strings.Fields(line)
	if data[0] <= "0" || data[1] <= "0" {
		fmt.Println("Error: warehouse cannot be null\n😱")
		os.Exit(0)
	}
	if data[2] < "10" || data[2] > "100000" {
		fmt.Println("Error: number of turn is out of range\n😱")
		os.Exit(0)
	}
	/*
		warehouse.width = data[0]
		warehouse.length = data[1]
		warehouse.turn = data[2]*/
	/* return warehouse */
}

// Check if the last line match with awaited format
func LastLineHandler(line string /*, warehouse Warehouse) /*, Truck */ {
	match, _ := regexp.MatchString(`(\d+)\s*(\d+)\s*(\d+)\s*(\d+)`, line)
	if match == false {
		fmt.Println("Error: format of Truck line is wrong\n😱")
		fmt.Println(line)
		os.Exit(0)
	}

	/* Following lines are used to split the different string in the line
	data := strings.Fields(line)
	 if warehouse.width < data[0] || warehouse.length < data[1] {
		 os.Exit(0)
	 }
	 if data[Ø] < 0 || data[1] < 0 {
		 os.Exit(0)
	 }
	 var truck Truck
	 truck.pos.x = data[0]
	 truck.pos.y = data[1]
	 truck.max_charge = data[2]
	 truck.availibility = data[3]
	*/
	//return truck
}

// Check if parcel line match with awaited format
func ParcelHandler(line string /*warehouse Warehouse*/) /* Parcel */ {
	match, _ := regexp.MatchString(`(\w+)\s*(\d+)\s*(\d+)\s*(\w+)`, line)
	if match == false {
		fmt.Println("Error: format of Parcel line is wrong\n😱")
		fmt.Println(line)
		os.Exit(0)
	}

	/* Following lines are used to split the different string in the line
	data := strings.Fields(line)
	if data[1] < 0 || data[2] < 0 {
		os.Exit(0)
	}
	if warehouse.width < data[1] || warehouse.length < data[2] {
		os.Exit(0)
	}*/
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
func PalletTruckHandler(line string /*warehouse Warehouse*/) /*PalletTruck*/ {
	match, _ := regexp.MatchString(`(\w+)\s*(\d+)\s*(\d+)`, line)
	if match == false {
		fmt.Println("Error: format of Pallet Truck is wrong\n😱")
		fmt.Println(line)
		os.Exit(0)
	}

	/* Following lines are used to split the different string in the line
	data := strings.Fields(line)
	if data[1] < 0 || data[2] < 0 {
		os.Exit(0)
	}
	if warehouse.width < data[1] || warehouse.length < data[2] {
		os.Exit(0)
	}
	var palletTruck PalletTruck
	palletTruck.name := data[Ø]
	palletTruck.pos.x := data[1]
	palletTruck.pos.y := data[2]
	return palletTruck
	*/
}

func ParsingHandler(line string, nb_lines int, count int) {
	if count == 1 {
		/* warehouse Warehouse := */ FirstLineHandler(line)
		return
	}
	if count == nb_lines {
		/* truck Truck := */ LastLineHandler(line)
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

func main() {
	if NbArgsHandler() == false {
		os.Exit(0)
	}

	content, err := os.Open(os.Args[1])
	if OpenFileHandler(err) == false {
		os.Exit(0)
	}
	nb_lines := 0
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		nb_lines += 1
	}
	content.Close()

	file, err := os.Open(os.Args[1])
	if OpenFileHandler(err) == false {
		os.Exit(0)
	}
	count := 0
	scanner2 := bufio.NewScanner(file)
	for scanner2.Scan() {
		count += 1
		ParsingHandler(scanner2.Text(), nb_lines, count)
	}
	file.Close()
	fmt.Println("Parsing done succesfully")
	//start game loop here
	return
}
