import (
	"fmt"
)

type Position struct {
	x int
	y int
}

funct isAdjacent(pos1 Position, pos2 Position) bool {
	return (
		(pos1.x + 1 == pos2.x && pos1.y == pos2.y) ||
		(pos1.x - 1 == pos2.x && pos1.y == pos2.y) ||
		(pos1.x == pos2.x && pos1.y + 1 == pos2.y) ||
		(pos1.x == pos2.x && pos1.y - 1 == pos2.y)
	)
}

type PalletTruck struct {
	pos Position
	parcel Parcel
}

// Todo: move these struct in another file
type Parcel struct {
	pos Position
	weight int
	color int
}

type Warehouse struct {
	width int
	height int
	nbTurn int
	palletTrucks []PalletTruck
	parcels []Parcel
}

type MoveDirection int

const(
	UP = iota
	DOWN = iota
	RIGHT = iota
	LEFT = iota
)

// Check if a position is valid for a pallet truck
func checkPosition(pos Position, wh Warehouse) bool {
	if (pos.x >= wh.width || pos.x < 0 || pos.y >= wh.height || pos.y < 0) {
		fmt.Println("Error: Cannot move outside of warehouse")
		return false;
	}
	for _, s := range wh.palletTrucks {
		if (s.pos.x == pos.x && s.pos.y == pos.y) {
			fmt.Println("Error: An other pallet truck is on this position")
			return false;
		}
	}
	for _, s := range wh.parcels {
		if (s.pos.x == pos.x && s.pos.y == pos.y) {
			fmt.Println("Error: A pacel is on this position")
			return false;
		}
	}
	return true;
}

// Return a new position that have applied the direction move
func move(pos Position, direction MoveDirection) pos {
	newPos := pos
	switch direction {
	case UP:
		newPos.y = pos.y - 1
	case DOWN:
		newPos.y = pos.y + 1
	case RIGHT:
		newPos.x = pos.x + 1
	case LEFT:
		newPos.x = pos.x - 1
	return newPos;
}

func applyPosition(pt PalletTruck, pos Position) {
	pt.pos = Position{pos}
}

func remove(s []Parcel, i int) []Parcel {
	s[len(s) - 1], s[i] = s[i], s[len(s) - 1]
	return s[:len(s) - 1]
}

func peekParcel(pt PalletTruck, wh Warehouse, idx int) {
	if (isAdjacent(pt.pos, wh.parcels[idx].pos)) {
		pt.parcel = wh.parcels[idx]
		wh.parcels = remove(wh.parcels, idx)
		return true
	}
	return false
}