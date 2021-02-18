package parsing

import (
	"reflect"
	"testing"

	gameData "github.com/LaurentColoma/M-BDX-002/gameData"
)

var expectedWarehouse = gameData.Warehouse{
	Width:  5,
	Height: 5,
	NbTurn: 10,
	PalletTrucks: []gameData.PalletTruck{
		{
			Pos: gameData.Position{
				X: 0,
				Y: 0,
			},
			Parcel: gameData.Parcel{
				Pos: gameData.Position{
					X: 0,
					Y: 0,
				},
				Name:   "",
				Weight: 0,
			},
			Name:   "transpalette_1",
			Status: 0,
		},
	},
	Parcels: []gameData.Parcel{
		{
			Pos: gameData.Position{
				X: 2,
				Y: 1,
			},
			Name:   "colis_a_livrer",
			Weight: 200,
		},
		{
			Pos: gameData.Position{
				X: 2,
				Y: 2,
			},
			Name:   "paquet",
			Weight: 500,
		},
		{
			Pos: gameData.Position{
				X: 0,
				Y: 3,
			},
			Name:   "deadpool",
			Weight: 100,
		},
		{
			Pos: gameData.Position{
				X: 4,
				Y: 1,
			},
			Name:   "colere_DU_dragon",
			Weight: 200,
		},
	},
	Truck: gameData.Storage{
		Pos: gameData.Position{
			X: 3,
			Y: 4,
		},
		Name:     "",
		Capacity: 4,
		Status:   0,
		Upturn:   5,
	},
}

func TestFirstLineHandler(t *testing.T) {
	var wh gameData.Warehouse

	FirstLineHandler("5 5 10", &wh)

	if wh.Width != 5 || wh.Height != 5 || wh.NbTurn != 10 {
		t.Error("First line is not parsed correctly")
	}
}

func TestLastLineHandler(t *testing.T) {
	var wh gameData.Warehouse

	wh.Width = 5
	wh.Height = 5

	var truck = LastLineHandler("3 4 4 5", &wh)

	if truck.Pos.X != 3 || truck.Pos.Y != 4 || truck.Capacity != 4 || truck.Upturn != 5 {
		t.Error("Last line is not parsed correctly")
	}
}

func TestParcelHandler(t *testing.T) {
	var wh gameData.Warehouse

	wh.Width = 5
	wh.Height = 5

	var parcel = ParcelHandler("paquet 2 3 BLUE", &wh)

	if parcel.Pos.X != 2 || parcel.Pos.Y != 3 || parcel.Weight != 500 || parcel.Name != "paquet" {
		t.Error("Parcel line is not parsed correctly")
	}
}

func TestTruckHandler(t *testing.T) {
	var wh gameData.Warehouse

	wh.Width = 5
	wh.Height = 5

	var palletTruck = PalletTruckHandler("transpalette 1 2", &wh)

	if palletTruck.Pos.X != 1 || palletTruck.Pos.Y != 2 || palletTruck.Name != "transpalette" {
		t.Error("PalletTruck line is not parsed correctly")
	}
}

func TestParsingHandler(t *testing.T) {
	var lines = []string{
		"5 5 10",
		"colis_a_livrer 2 1 green",
		"paquet 2 2 BLUE",
		"deadpool 0 3 yellow",
		"colere_DU_dragon 4 1 green",
		"transpalette_1 0 0",
		"3 4 4 5",
	}
	var wh gameData.Warehouse

	for i, line := range lines {
		ParsingHandler(line, len(lines), i+1, &wh)
	}

	if !reflect.DeepEqual(wh, expectedWarehouse) {
		t.Errorf("Expected %v got %v", expectedWarehouse, wh)
	}
}
