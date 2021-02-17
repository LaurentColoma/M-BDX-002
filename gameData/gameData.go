package gameData

import (
	"fmt"
)

type position struct {
	X int
	Y int
}

func isAdjacent(pos1 position, pos2 position) bool {
	return ((pos1.X+1 == pos2.X && pos1.Y == pos2.Y) ||
		(pos1.X-1 == pos2.X && pos1.Y == pos2.Y) ||
		(pos1.X == pos2.X && pos1.Y+1 == pos2.Y) ||
		(pos1.X == pos2.X && pos1.Y-1 == pos2.Y))
}

type PalletTruck struct {
	Pos    position
	Parcel Parcel
	Name   string
}

type Storage struct {
	Pos      position
	Name     string // by default "camion"
	Capacity int
}

// Todo: move these struct in another file
type Parcel struct {
	Pos    position
	Weight int
}

type Warehouse struct {
	Width        int
	Height       int
	NbTurn       int
	PalletTrucks []PalletTruck
	Parcels      []Parcel
	Truck        Storage
}

type MoveDirection int

const (
	UP    = iota
	DOWN  = iota
	RIGHT = iota
	LEFT  = iota
)

// Check if a position is valid for a pallet truck
func checkPosition(pos position, wh Warehouse) bool {
	if pos.X >= wh.Width || pos.X < 0 || pos.Y >= wh.Height || pos.Y < 0 {
		fmt.Println("Error: Cannot move outside of warehouse")
		return false
	}
	for _, s := range wh.PalletTrucks {
		if s.Pos.X == pos.X && s.Pos.Y == pos.Y {
			fmt.Println("Error: An other pallet truck is on this position")
			return false
		}
	}
	for _, s := range wh.Parcels {
		if s.Pos.X == pos.X && s.Pos.Y == pos.Y {
			fmt.Println("Error: A pacel is on this position")
			return false
		}
	}
	return true
}

// Return a new position that have applied the direction move
func move(pos position, direction MoveDirection) (newPos position) {
	newPos = pos
	switch direction {
	case UP:
		newPos.Y = pos.Y - 1
	case DOWN:
		newPos.Y = pos.Y + 1
	case RIGHT:
		newPos.X = pos.X + 1
	case LEFT:
		newPos.X = pos.X - 1
	}
	return
}

func applyPosition(pt PalletTruck, pos position) {
	pt.Pos = pos
}

func remove(s []Parcel, i int) []Parcel {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func peekParcel(pt PalletTruck, wh Warehouse, idx int) bool {
	if isAdjacent(pt.Pos, wh.Parcels[idx].Pos) {
		pt.Parcel = wh.Parcels[idx]
		wh.Parcels = remove(wh.Parcels, idx)
		return true
	}
	return false
}
