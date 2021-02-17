package gameLoop

import (
	"fmt"

	gameData "github.com/LaurentColoma/M-BDX-002/gameData"
)

func gameLoop(warehouse gameData.Warehouse) {
	for i := 0; i < warehouse.nbTurn; i++ {
		fmt.Printf("tour %v\n", i+1)
		for i, s := range warehouse.PalletTrucks {
			fmt.Println("%v %v [%v,%v]\n")
		}
	}

}
