package gameLoop

import (
	"fmt"

	gameData "github.com/LaurentColoma/M-BDX-002/gameData"
)

<<<<<<< HEAD
func gameLoop(warehouse gameData.Warehouse) {
=======
func gameLoop(warehouse gameData.Warehouse)  {
	state := [...]string{"GO", "WAIT", "TAKE", "LEAVE", "GONE", "WAITING"}

	gameData.Warehouse.Truck.Status = 5
	for i := range warehouse.PalletTrucks {
		gameData.Warehouse.PalletTrucks[i].Status = 1
	}
>>>>>>> 7bc1be5ad53bd8d5998f3e33b9ac89dee8c2e13d
	for i := 0; i < warehouse.NbTurn; i++ {
		fmt.Printf("tour %v\n", i+1)

		for i := range warehouse.PalletTrucks {
			fmt.Printf("%v %v [%v,%v]\n", warehouse.PalletTrucks[i].Name,
			state[], warwarehouse.PalletTrucks[i].Pos.X, warwarehouse.PalletTrucks[i].Pos.X)
		}
	}

}