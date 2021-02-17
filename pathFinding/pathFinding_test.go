package pathFinding

import (
	"reflect"
	"testing"

	gameData "github.com/LaurentColoma/M-BDX-002/gameData"
)

func getPossibleWarehouse() gameData.Warehouse {
	var wh = gameData.Warehouse {
		Width: 5,
		Height: 4,
		NbTurn: 1000,
		PalletTrucks: []gameData.PalletTruck{
		 {
			Pos: gameData.Position{
			 X: 1,
			 Y: 2,
			},
			Parcel: gameData.Parcel{
			 Pos: gameData.Position{
				X: 0,
				Y: 0,
			 },
			 Name: "",
			 Weight: 0,
			},
			Name: "transpalette_1",
			Status: 0,
		 },
		},
		Parcels: []gameData.Parcel{
		 {
			Pos: gameData.Position{
			 X: 2,
			 Y: 1,
			},
			Name: "colis_a_livrer",
			Weight: 200,
		 },
		 {
			Pos: gameData.Position{
			 X: 2,
			 Y: 1,
			},
			Name: "paquet",
			Weight: 500,
		 },
		 {
			Pos: gameData.Position{
			 X: 2,
			 Y: 0,
			},
			Name: "paquet2",
			Weight: 500,
		 },
		 {
			Pos: gameData.Position{
			 X: 1,
			 Y: 0,
			},
			Name: "deadpool",
			Weight: 100,
		 },
		},
		Truck: gameData.Storage{
		 Pos: gameData.Position{
			X: 3,
			Y: 4,
		 },
		 Name: "",
		 Capacity: 4000,
		 Status: 0,
		 Upturn: 5,
		},
	}

	return wh
}

func getImpossibleWarehouse() gameData.Warehouse {
	var wh = gameData.Warehouse {
		Width: 5,
		Height: 4,
		NbTurn: 1000,
		PalletTrucks: []gameData.PalletTruck{
		 {
			Pos: gameData.Position{
			 X: 1,
			 Y: 2,
			},
			Parcel: gameData.Parcel{
			 Pos: gameData.Position{
				X: 0,
				Y: 0,
			 },
			 Name: "",
			 Weight: 0,
			},
			Name: "transpalette_1",
			Status: 0,
		 },
		},
		Parcels: []gameData.Parcel{
		 {
			Pos: gameData.Position{
			 X: 2,
			 Y: 1,
			},
			Name: "colis_a_livrer",
			Weight: 200,
		 },
		 {
			Pos: gameData.Position{
			 X: 2,
			 Y: 1,
			},
			Name: "paquet",
			Weight: 500,
		 },
		 {
			Pos: gameData.Position{
			 X: 3,
			 Y: 0,
			},
			Name: "paquet2",
			Weight: 500,
		 }, 
		 {
			Pos: gameData.Position{
			 X: 1,
			 Y: 0,
			},
			Name: "deadpool",
			Weight: 100,
		 },
		},
		Truck: gameData.Storage{
		 Pos: gameData.Position{
			X: 3,
			Y: 4,
		 },
		 Name: "",
		 Capacity: 4000,
		 Status: 0,
		 Upturn: 5,
		},
	}

	return wh
}

func TestMapFrom(t *testing.T) {
	var wh = getPossibleWarehouse()

	var m = MapFrom(&wh, 1, 2)
	var expectedM = []int{ 3, -1, -1, 4, 5, 2, 1, -1, 3, 4, 1, -2, 1, 2, 3, 2, 1, 2, 3, 4 }

	if !reflect.DeepEqual(m, expectedM) {
		t.Errorf("Expected %v got %v", expectedM, m)
	}

}

func TestGetRoute(t *testing.T) {
	var wh = getPossibleWarehouse()

	var m = MapFrom(&wh, 1, 2)
	var p = GetRoute(m, wh.Width, wh.Height, 2, 0)
	var expectedP = [][2]int{{2, 2}, {3, 2}, {3, 1}, {3, 0}, {2, 0}}

	if !reflect.DeepEqual(p, expectedP) {
		t.Errorf("Expected %v got %v", expectedP, p)
	}

	var iwh = getImpossibleWarehouse()

	m = MapFrom(&iwh, 1, 2)
	p = GetRoute(m, wh.Width, wh.Height, 2, 0)
	expectedP = [][2]int{{2, 2}, {3, 2}, {3, 1}, {3, 0}, {2, 0}}

	if reflect.DeepEqual(p, expectedP) {
		t.Errorf("Expected %v != %v", expectedP, p)
	}

}