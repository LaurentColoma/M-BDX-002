package gameLoop

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	gameData "github.com/LaurentColoma/M-BDX-002/gameData"
	pathFinding "github.com/LaurentColoma/M-BDX-002/pathFinding"
)

func miniParcel(warehouse gameData.Warehouse) (mini int, index int) {
	minim := 501
	index = 0
	for i := range warehouse.Parcels {
		if warehouse.Parcels[i].Weight < minim {
			minim = warehouse.Parcels[i].Weight
			index = i
			if minim == 100 {
				return minim, index
			}
		}
	}
	return minim, index
}

func giveParcel(pt *gameData.PalletTruck, wh *gameData.Warehouse) {
	m := pathFinding.MapFrom(wh, pt.Pos.X, pt.Pos.Y)
	var paths [][2]int
	var p [][2]int

	for i := range wh.Parcels {
		p = pathFinding.GetRoute(m, wh.Width, wh.Height, wh.Parcels[i].Pos.X, wh.Parcels[i].Pos.Y)
		paths[i] = p[len(paths)-1]
	}
	lowest := math.Abs(float64(paths[0][0])-float64(pt.Pos.X)) +
		math.Abs(float64(paths[0][1])-float64(pt.Pos.Y))
	index := 0
	for j := range paths {
		if lowest > math.Abs(float64(paths[j][0])-float64(pt.Pos.X))+
			math.Abs(float64(paths[j][1])-float64(pt.Pos.Y)) && wh.Parcels[j].Aimed == false {
			index = j
		}
	}
	pt.Parcel.Pos.X = wh.Parcels[index].Pos.X
	pt.Parcel.Pos.Y = wh.Parcels[index].Pos.Y
	wh.Parcels[index].Aimed = true
	pt.Parcel.Weight = 1
}

func GameLoop(warehouse gameData.Warehouse) int {
	state := []string{"WAIT", "TAKE", "LEAVE", "WAITING", "GONE"}
	weight := []string{"YELLOW", "GREEN", "EMPTY", "EMPTY", "BLUE"}
	currentLoad := 0
	waitBeforeComing := 0
	truckLeft := false

	warehouse.Truck.Status = sort.SearchStrings(state, "WAITING")
	for i := range warehouse.PalletTrucks {
		warehouse.PalletTrucks[i].Status = sort.SearchStrings(state, "WAIT")
	}
	for i := 0; i < warehouse.NbTurn; i++ {
		if len(warehouse.Parcels) == 0 {
			return 1
		}
		minim, index := miniParcel(warehouse)
		if truckLeft == false && warehouse.Truck.Capacity-currentLoad < minim {
			warehouse.Truck.Status = sort.SearchStrings(state, "GONE")
			waitBeforeComing = warehouse.Truck.Upturn
			truckLeft = true
		}
		for i := range warehouse.PalletTrucks {
			warehouse.PalletTrucks[i].Status = sort.SearchStrings(state, "WAIT")
		}
		fmt.Printf("tour %v\n", i+1)

		if waitBeforeComing == 0 && truckLeft == true {
			truckLeft = false
			currentLoad = 0
			warehouse.Truck.Status = sort.SearchStrings(state, "WAITING")
		} else if truckLeft == true && waitBeforeComing > 0 {
			waitBeforeComing--
		}

		for i := range warehouse.PalletTrucks {
			// we drop the parcel into truck
			if truckLeft == false && warehouse.PalletTrucks[i].Parcel.Weight > 1 && gameData.DropParcel(warehouse.PalletTrucks[i], warehouse) == true {
				currentLoad += warehouse.PalletTrucks[i].Parcel.Weight
				warehouse.PalletTrucks[i].Status = sort.SearchStrings(state, "LEAVE")
			} else if warehouse.PalletTrucks[i].Parcel.Weight == 1 && gameData.PeekParcel(&warehouse.PalletTrucks[i], &warehouse, index) == true {
				warehouse.PalletTrucks[i].Status = sort.SearchStrings(state, "TAKE")
			} else if warehouse.PalletTrucks[i].Parcel.Weight == 0 {
				giveParcel(&warehouse.PalletTrucks[i], &warehouse)
			}
			if warehouse.PalletTrucks[i].Status != sort.SearchStrings(state, "TAKE") && warehouse.PalletTrucks[i].Status != sort.SearchStrings(state, "LEAVE") {
				x := warehouse.PalletTrucks[i].Path[0][0] - warehouse.PalletTrucks[i].Pos.X
				y := warehouse.PalletTrucks[i].Path[0][1] - warehouse.PalletTrucks[i].Pos.Y
				res := strconv.Itoa(x) + strconv.Itoa(y)
				switch res {
				case "10":
					warehouse.PalletTrucks[i].Pos = gameData.Move(warehouse.PalletTrucks[i].Pos, gameData.RIGHT)
				case "01":
					warehouse.PalletTrucks[i].Pos = gameData.Move(warehouse.PalletTrucks[i].Pos, gameData.DOWN)
				case "-10":
					warehouse.PalletTrucks[i].Pos = gameData.Move(warehouse.PalletTrucks[i].Pos, gameData.LEFT)
				case "0-1":
					warehouse.PalletTrucks[i].Pos = gameData.Move(warehouse.PalletTrucks[i].Pos, gameData.DOWN)
				}
				warehouse.PalletTrucks[i].Status = sort.SearchStrings(state, "GO")
			}
			if warehouse.PalletTrucks[i].Status == sort.SearchStrings(state, "TAKE") ||
				warehouse.PalletTrucks[i].Status == sort.SearchStrings(state, "LEAVE") {
				fmt.Printf("%v %v [%v,%v] %v %v\n", warehouse.PalletTrucks[i].Name,
					state[warehouse.PalletTrucks[i].Status],
					warehouse.PalletTrucks[i].Pos.X,
					warehouse.PalletTrucks[i].Pos.X,
					warehouse.PalletTrucks[i].Parcel.Name,
					weight[warehouse.PalletTrucks[i].Parcel.Weight/100+1])
				if warehouse.PalletTrucks[i].Status == sort.SearchStrings(state, "LEAVE") {
					warehouse.PalletTrucks[i].Parcel = gameData.Parcel{}
				}
			} else {
				fmt.Printf("%v %v [%v,%v]\n", warehouse.PalletTrucks[i].Name,
					state[warehouse.PalletTrucks[i].Status],
					warehouse.PalletTrucks[i].Pos.X,
					warehouse.PalletTrucks[i].Pos.X)
			}
		}
		fmt.Printf("camion %v %v %v\n\n", state[warehouse.Truck.Status], warehouse.Truck.Capacity, currentLoad)
	}
	return 0
}
