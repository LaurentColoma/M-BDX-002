package gameLoop

import (
	"fmt"
	"math"
	"strconv"

	gameData "github.com/LaurentColoma/M-BDX-002/gameData"
	"github.com/LaurentColoma/M-BDX-002/pathFinding"
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
	var index int
	lowest := math.Abs(float64(wh.Parcels[0].Pos.X)-float64(pt.Pos.X)) +
		math.Abs(float64(wh.Parcels[0].Pos.Y)-float64(pt.Pos.Y))

	for i := range wh.Parcels {
		if lowest > math.Abs(float64(wh.Parcels[i].Pos.X)-float64(pt.Pos.X))+
			math.Abs(float64(wh.Parcels[i].Pos.Y)-float64(pt.Pos.Y)) && wh.Parcels[i].Aimed == false {
			index = i
		}
	}
	pt.Parcel.Pos.X = wh.Parcels[index].Pos.X
	pt.Parcel.Pos.Y = wh.Parcels[index].Pos.Y
	pt.Index = index
	wh.Parcels[index].Aimed = true
	pt.Parcel.Weight = 1
}

func GameLoop(warehouse gameData.Warehouse) int {
	state := []string{"WAIT", "TAKE", "LEAVE", "WAITING", "GONE", "GO"}
	weight := []string{"YELLOW", "GREEN", "EMPTY", "EMPTY", "BLUE"}
	currentLoad := 0
	waitBeforeComing := 0
	truckLeft := false

	warehouse.Truck.Status = 3
	for i := range warehouse.PalletTrucks {
		warehouse.PalletTrucks[i].Status = 0
	}
	for i := 0; i < warehouse.NbTurn; i++ {
		if len(warehouse.Parcels) == 0 {
			return 1
		}
		minim, index := miniParcel(warehouse)
		if truckLeft == false && warehouse.Truck.Capacity-currentLoad < minim {
			warehouse.Truck.Status = 4
			waitBeforeComing = warehouse.Truck.Upturn
			truckLeft = true
		}
		for i := range warehouse.PalletTrucks {
			warehouse.PalletTrucks[i].Status = 0
		}
		fmt.Printf("tour %v\n", i+1)

		if waitBeforeComing == 0 && truckLeft == true {
			truckLeft = false
			currentLoad = 0
			warehouse.Truck.Status = 3
		} else if truckLeft == true && waitBeforeComing > 0 {
			waitBeforeComing--
		}

		for i := range warehouse.PalletTrucks {
			// we drop the parcel into truck
			if truckLeft == false && warehouse.PalletTrucks[i].Parcel.Weight > 1 && gameData.DropParcel(warehouse.PalletTrucks[i], warehouse) == true {
				currentLoad += warehouse.PalletTrucks[i].Parcel.Weight
				warehouse.PalletTrucks[i].Status = 2
			} else if warehouse.PalletTrucks[i].Parcel.Weight == 1 && gameData.PeekParcel(&warehouse.PalletTrucks[i], &warehouse, warehouse.PalletTrucks[i].Index) == true {
				warehouse.PalletTrucks[i].Status = 1
			} else if warehouse.PalletTrucks[i].Parcel.Weight == 0 {
				giveParcel(&warehouse.PalletTrucks[i], &warehouse)
			}
			if warehouse.PalletTrucks[i].Status != 1 && warehouse.PalletTrucks[i].Status != 2 {
				m := pathFinding.MapFrom(&warehouse, warehouse.PalletTrucks[i].Pos.X, warehouse.PalletTrucks[i].Pos.Y)
				r := pathFinding.GetRoute(m, warehouse.Width, warehouse.Height, warehouse.PalletTrucks[i].Parcel.Pos.X, warehouse.PalletTrucks[i].Parcel.Pos.Y)
				x := r[0][0] - warehouse.PalletTrucks[i].Pos.X
				y := r[0][1] - warehouse.PalletTrucks[i].Pos.Y
				res := strconv.Itoa(x) + strconv.Itoa(y)
				switch res {
				case "10":
					warehouse.PalletTrucks[i].Pos = gameData.Move(warehouse.PalletTrucks[i].Pos, gameData.RIGHT)
				case "01":
					warehouse.PalletTrucks[i].Pos = gameData.Move(warehouse.PalletTrucks[i].Pos, gameData.DOWN)
				case "-10":
					warehouse.PalletTrucks[i].Pos = gameData.Move(warehouse.PalletTrucks[i].Pos, gameData.LEFT)
				case "0-1":
					warehouse.PalletTrucks[i].Pos = gameData.Move(warehouse.PalletTrucks[i].Pos, gameData.UP)
				}
				warehouse.PalletTrucks[i].Status = 5
			}
			if warehouse.PalletTrucks[i].Status == 1 ||
				warehouse.PalletTrucks[i].Status == 2 {
				fmt.Printf("%v %v [%v,%v] %v %v\n", warehouse.PalletTrucks[i].Name,
					state[warehouse.PalletTrucks[i].Status],
					warehouse.PalletTrucks[i].Pos.X,
					warehouse.PalletTrucks[i].Pos.X,
					warehouse.PalletTrucks[i].Parcel.Name,
					weight[warehouse.PalletTrucks[i].Parcel.Weight/100+1])
				if warehouse.PalletTrucks[i].Status == 2 {
					warehouse.PalletTrucks[i].Parcel = gameData.Parcel{}
				}
			} else if warehouse.PalletTrucks[i].Status == 5 {
				fmt.Printf("%v %v [%v,%v]\n", warehouse.PalletTrucks[i].Name,
					state[warehouse.PalletTrucks[i].Status],
					warehouse.PalletTrucks[i].Pos.X,
					warehouse.PalletTrucks[i].Pos.Y)
			} else {
				fmt.Printf("%v %v\n", warehouse.PalletTrucks[i].Name,
					state[warehouse.PalletTrucks[i].Status])
			}
		}
		fmt.Printf("camion %v %v %v\n\n", state[warehouse.Truck.Status], warehouse.Truck.Capacity, currentLoad)
	}
	return 0
}
