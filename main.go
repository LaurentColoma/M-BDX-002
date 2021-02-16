package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// mapCoords stores every infos about the passed file
type mapCoords struct {
	roundLeft         int      // number of total rounds left to be decremented each turn
	truckCapacity     int      // admissible charge by the truck, shoudln't change
	truckBeforeArrive int      // decrementive rounds when the truck has to leave (set by the 3rd value of truck array)
	size              [2]int   // size of the map
	truck             [3]int   // truck position and number of round required when it leaves
	boxes             [][3]int // first dimension is number of boxes, second dimension values are x and y positions
	pallets           [][2]int // first dimension is number of pallets, second dimension values are x and y positions
}

// NbArgsHandler check the numbers of parameters
func NbArgsHandler() bool {
	if len(os.Args) < 2 {
		fmt.Println("Error: wrong number of arguments\n😱")
		os.Exit(0)
	}
	return true
}

// OpenFileHandler tries to open the file
func OpenFileHandler(err error) bool {
	if err != nil {
		fmt.Println("Error: file not supported\n😱")
		os.Exit(0)
	}
	return true
}

func FirstLineHandler(line string) bool {
	match, _ := regexp.MatchString("[0-9][\t\n\v\f\r ][0-9][\t\n\v\f\r ][0-9]", line)
	if match == false {
		fmt.Println("Error: format of first line is wrong\n😱")
		os.Exit(0)
	}
	return match
}

func LastLineHandler(line string) bool {
	match, _ := regexp.MatchString("[0-9][\t\n\v\f\r ][0-9][\t\n\v\f\r ][0-9][\t\n\v\f\r ][0-9]", line)
	if match == false {
		fmt.Println("Error: format of last line is wrong\n😱")
		os.Exit(0)
	}
	return match
}

func ParcelHandler(line string) bool {
	match, _ := regexp.MatchString("[A-Za-z][\t\n\v\f\r ][0-9][\t\n\v\f\r ][0-9][\t\n\v\f\r ][A-Za-z]", line)
	if match == false {
		fmt.Println("Error: format of parcel line is wrong\n😱")
	}
	return match
}

func PalletTruckHandler(line string) bool {
	match, _ := regexp.MatchString("[A-Za-z][\t\n\v\f\r ][0-9][\t\n\v\f\r ][0-9]", line)
	if match == false {
		fmt.Println("Error: format of pallet truck is wrong\n😱")
	}
	return match
}

func ParsingHandler(line string, nb_lines int, count int) bool {
	match := false
	if count == 0 {
		return FirstLineHandler(line)
	}
	if count == nb_lines {
		return LastLineHandler(line)
	}
	if ParcelHandler(line) == true {
		match = true
	} else if PalletTruckHandler(line) == true {
		match = true
	}
	return match
}

func main() {
	game := mapCoords

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
		if ParsingHandler(scanner2.Text(), nb_lines, count) == false {
			os.Exit(0)
		}
		fmt.Println(scanner2.Text())
		count += 1
	}
	file.Close()

	if game.roundLeft == 0 {
		fmt.Println("🙂")
	} else {
		fmt.Println("😎")
	}
	return
}
