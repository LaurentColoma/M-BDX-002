package gameLoop

import (
	"fmt"

	gameData "github.com/LaurentColoma/M-BDX-002/gameData"
)

func gameLoop(warehouse gameData.Warehouse) {
	for i := 0; i < warehouse.NbTurn; i++ {
		fmt.Printf("tour %v\n", i+1)

		for i := range warehouse.PalletTrucks {
			fmt.Printf("%v %v [%v,%v]\n", warehouse.PalletTrucks[i].Name,
			state[], warwarehouse.PalletTrucks[i].Pos.X, warwarehouse.PalletTrucks[i].Pos.X)
		}
	}

}
