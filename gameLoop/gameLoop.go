package gameLoop

import (
	"fmt"
	"sort"

	gameData "github.com/LaurentColoma/M-BDX-002/gameData"
)

func gameLoop(warehouse gameData.Warehouse) {
	state := [...]string{"WAIT", "TAKE", "LEAVE", "WAITING", "GONE"}
	totalLoad := 0

	warehouse.Truck.Status = sort.SearchStrings(state, "WAITING")
	for i := range warehouse.PalletTrucks {
		warehouse.PalletTrucks[i].Status = sort.SearchStrings(state, "WAIT")
	}
	for i := range warehouse.Parcels {
		totalLoad += warehouse.Parcels[i].Weight
	}
	for i := 0; i < warehouse.NbTurn; i++ {
		fmt.Printf("tour %v\n", i+1)

		for i := range warehouse.PalletTrucks {
			fmt.Printf("%v %v [%v,%v]\n", warehouse.PalletTrucks[i].Name,
				state[warehouse.PalletTrucks[i].Status],
				warwarehouse.PalletTrucks[i].Pos.X,
				warwarehouse.PalletTrucks[i].Pos.X)
		}
	}

}
