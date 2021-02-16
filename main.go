package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func NbArgsHandler() bool {
	if len(os.Args) < 2 {
		fmt.Println("Error: wrong number of arguments")
		os.Exit(0)
	}
	return true
}

func OpenFileHandler(err error) bool {
	if err != nil {
		fmt.Println("Error: file not supported")
		os.Exit(0)
	}
	return true
}

func FirstLineHandler(line string) bool {
	match, _ := regexp.MatchString("[0-9][\t\n\v\f\r ][0-9][\t\n\v\f\r ][0-9]", line)
	if match == false {
		fmt.Println("Error: format of first line is wrong")
		os.Exit(0)
	}
	return match
}

func LastLineHandler(line string) bool {
	match, _ := regexp.MatchString("[0-9][\t\n\v\f\r ][0-9][\t\n\v\f\r ][0-9][\t\n\v\f\r ][0-9]", line)
	if match == false {
		fmt.Println("Error: format of last line is wrong")
		os.Exit(0)
	}
	return match
}

func ParcelHandler(line string) bool {
	match, _ := regexp.MatchString("[A-Za-z][\t\n\v\f\r ][0-9][\t\n\v\f\r ][0-9][\t\n\v\f\r ][A-Za-z]", line)
	if match == false {
		fmt.Println("Error: format of parcel line is wrong")
	}
	return match
}

func PalletTruckHandler(line string) bool {
	match, _ := regexp.MatchString("[A-Za-z][\t\n\v\f\r ][0-9][\t\n\v\f\r ][0-9]", line)
	if match == false {
		fmt.Println("Error: format of pallet truck is wrong")
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

	return
}
