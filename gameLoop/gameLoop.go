package gameLoop

import (
	"fmt"
	"sort"

	gameData "github.com/LaurentColoma/M-BDX-002/gameData"
)

func weightLeft(parcels gameData.Warehouse.[]Parcel) (rest int) {
	rest := 0
	for i := range parcels {
		rest += parcels[i].Weight
	}
}

func miniParcel(parcels gameData.Warehouse.[]Parcel) (mini int, index int) {
	minim := 501
	index := 0
	for i := range parcels {
		if parcels[i].Weight < minim {
			minim = parcels[i].Weight
			index = i
			if minim == 100 {
				return minim, index//, parcels[i].Pos.X, parcels[i].Pos.Y
			}
		}
	}
	return minim, index
}

func GameLoop(warehouse gameData.Warehouse) int {
	state := [...]string{"WAIT", "TAKE", "LEAVE", "WAITING", "GONE"}
	weight := [...]string{"YELLOW", "GREEN", "BLUE"}
	// totalLoad := 0
	currentLoad := 0
	waitBeforeComing := 0
	truckLeft := false

	warehouse.Truck.Status = sort.SearchStrings(state, "WAITING")
	for i := range warehouse.PalletTrucks {
		warehouse.PalletTrucks[i].Status = sort.SearchStrings(state, "WAIT")
	}
	// for i := range warehouse.Parcels {
	// 	totalLoad += warehouse.Parcels[i].Weight
	// }
	for i := 0; i < warehouse.NbTurn; i++ {
		if len(warehouse.Parcels) == 0
			return 1
		minim, index = miniParcel(warehouse.Parcels)
		if truckLeft == false && warehouse.Truck.Capacity - currentLoad < minim {
			warehouse.Truck.Status = sort.SearchStrings(state, "GONE")
			waitBeforeComing = warehouse.Truck.Upturn
			truckLeft = true
			for i := range warehouse.PalletTrucks {
				warehouse.PalletTrucks[i].Status = sort.SearchStrings(state, "WAIT")
			}
		}
		fmt.Printf("tour %v\n", i+1)

		if waitBeforeComing == 0 && truckLeft == true {
			truckLeft == false
			currentLoad = 0
		} else if truckLeft == true && waitBeforeComing > 0 {
			waitBeforeComing--
		}
		
		for i := range warehouse.PalletTrucks {
			
			if truckLeft == false && warehouse.PalletTrucks[i].Status != sort.SearchStrings(state, "TAKE") && gameData.PeekParcel(&warehouse.PalletTrucks[i], &warehouse, index) {
				warehouse.PalletTrucks[i].Status = sort.SearchStrings(state, "TAKE")
			}
			// get Path here
			fmt.Printf("%v %v [%v,%v]\n", warehouse.PalletTrucks[i].Name,
				state[warehouse.PalletTrucks[i].Status],
				warehouse.PalletTrucks[i].Pos.X,
				warehouse.PalletTrucks[i].Pos.X)
		} // le camion fait bipbip le batard
		fmt.printf("camion %v %v %v\n\n", state[warewarehouse.Truck.Status], warehwarehouse.Truck.Capacity, currcurrentLoad)
	}
	return 0
}
