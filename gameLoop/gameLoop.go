package gameLoop

import (
	"fmt"
	"sort"
	"strconv"

	PathFinding "github.com/LaurentColoma/M-BDX-002/PathFinding"
	gameData "github.com/LaurentColoma/M-BDX-002/gameData"
)

func miniParcel(warehouse gameData.Warehouse) (mini int, index int) {
	minim := 501
	index = 0
	for i := range warehouse.Parcels {
		if warehouse.Parcels[i].Weight < minim {
			minim = warehouse.Parcels[i].Weight
			index = i
			if minim == 100 {
				return minim, index //, parcels[i].Pos.X, parcels[i].Pos.Y
			}
		}
	}
	return minim, index
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
		} else if truckLeft == true && waitBeforeComing > 0 {
			waitBeforeComing--
		}

		for i := range warehouse.PalletTrucks {
			// we drop the parcel into truck
			if truckLeft == false && warehouse.PalletTrucks[i].Parcel.Weight != 0 && gameData.DropParcel(warehouse.PalletTrucks[i], warehouse) == true {
				currentLoad += warehouse.PalletTrucks[i].Parcel.Weight
				warehouse.PalletTrucks[i].Status = sort.SearchStrings(state, "LEAVE")
			}
			// taking a parcel if one is adjacent
			if warehouse.PalletTrucks[i].Parcel.Weight == 0 && gameData.PeekParcel(&warehouse.PalletTrucks[i], &warehouse, index) == true {
				warehouse.PalletTrucks[i].Status = sort.SearchStrings(state, "TAKE")
			}
			// each palletTruck looking for a parcel if nothing in sight yet
			if warehouse.PalletTrucks[i].Parcel.Weight == 0 {
				//warehouse.PalletTrucks[i].Parcel.Weight =
				m := PathFinding.MapFrom(&warehouse, warehouse.PalletTrucks[i].Pos.X, warehouse.PalletTrucks[i].Pos.Y)
				warehouse.PalletTrucks[i].Path = PathFinding.GetRoute(m, warehouse.Width, warehouse.Height, warehouse.PalletTrucks[i].Parcel.Pos.X, warehouse.PalletTrucks[i].Parcel.Pos.Y)
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
		} // le camion fait bipbip le batard
		fmt.Printf("camion %v %v %v\n\n", state[warehouse.Truck.Status], warehouse.Truck.Capacity, currentLoad)
	}
	return 0
}
