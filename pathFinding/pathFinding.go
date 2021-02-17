package pathFinding

import (
	gameData "github.com/LaurentColoma/M-BDX-002/gameData"
)

func mapFormStep(mapf []int, pointsToStartFrom [][2]int, w int, h int, count int) [][2]int {
	var newPointsToStartFrom [][2]int

	for _, c := range pointsToStartFrom {
		var x = c[0]
		var y = c[1]
		var pointsAround = [4][2]int{{x + 1, y}, {x - 1, y}, {x, y + 1}, {x, y - 1}}

		for _, nc := range pointsAround {
			var x = nc[0]
			var y = nc[1]
			if x >= 0 && y >= 0 && x < w && y < h {
				if mapf[y*w+x] == 0 {
					newPointsToStartFrom = append(newPointsToStartFrom, [2]int{x, y})
					mapf[y*w+x] = count
				}
			}
		}
	}
	return newPointsToStartFrom
}

func mapIsFullFilled(mapf []int, w int, h int) bool {
	for _, c := range mapf {
		if c == 0 {
			return false
		}
	}
	return true
}

// return a map of path length from a position where -1 is a wall and -2 the start point
// There is an example
// +---+----+----+---+---+
// | 3 | -1 | 5  | 4 | 5 |
// +---+----+----+---+---+
// | 2 | 1  | -1 | 3 | 4 |
// +---+----+----+---+---+
// | 1 | -2  | 1  | 2 | 3 |
// +---+----+----+---+---+
// | 2 | 1  | 2  | 3 | 4 |
// +---+----+----+---+---+
func MapFrom(wh *gameData.Warehouse, x int, y int) []int {
	var mapf = make([]int, wh.Width*wh.Height)

	for _, s := range wh.PalletTrucks {
		mapf[s.Pos.Y*wh.Width+s.Pos.X] = -1
	}
	for _, s := range wh.Parcels {
		mapf[s.Pos.Y*wh.Width+s.Pos.X] = -1
	}
	if wh.Truck.Status == 3 {
		mapf[wh.Truck.Pos.Y*wh.Width+wh.Truck.Pos.X] = -1
	}

	mapf[y*wh.Width+x] = -2

	var pointsToStartFrom = [][2]int{{x, y}}
	var count = 1
	for !mapIsFullFilled(mapf, wh.Width, wh.Height) && len(pointsToStartFrom) > 0 {
		pointsToStartFrom = mapFormStep(mapf, pointsToStartFrom, wh.Width, wh.Height, count)
		count += 1
	}

	return mapf
}

func GetRoute(mapf []int, w int, h int, x int, y int) (res [][2]int) {
	var c = mapf[y*w+x]

	if c <= 0 && c != -1 {
		return nil
	} else {
		c = w * h // we just want a very big number
	}

	res = append(res, [2]int{x, y})

	for c != -2 {
		var pointsAround = [4][2]int{{x + 1, y}, {x - 1, y}, {x, y + 1}, {x, y - 1}}
		var locked = true

		for _, pa := range pointsAround {
			var paX = pa[0]
			var paY = pa[1]
			if paX >= 0 && paY >= 0 && paX < w && paY < h {
				if mapf[paY*w+paX] == -1 {
					continue
				}
				if mapf[paY*w+paX] < c {
					c = mapf[paY*w+paX]
					y = paY
					x = paX
					locked = false
					if c > 0 {
						res = append([][2]int{{x, y}}, res...)
					}
					break
				}
			}
		}

		if locked {
			return nil
		}
	}
	return
}
