package game

import (
	"math/rand"
)

type Ship struct {
	size         int
	isHorizontal bool
	x            int
	y            int
}

func PlaceAllShips(board *GameBoard) {
	ships := []Ship{
		{size: 4},            // четырёх-палубный
		{size: 3}, {size: 3}, // два трёх-палубных
		{size: 2}, {size: 2}, {size: 2}, // три двух-палубных
		{size: 1}, {size: 1}, {size: 1}, {size: 1}, // четыре одно-палубных
	}

	for _, ship := range ships {
		placed := false

		for !placed {
			ship.x = rand.Intn(BOARD_SIZE) + 1
			ship.y = rand.Intn(BOARD_SIZE) + 1
			ship.isHorizontal = true //rand.Intn(2) == 0 // случайно выбираем ориентацию корабля

			if canPlaceShip(board, ship) {
				placeShip(board, ship)
				placed = true
			}
		}
	}
}

func placeShip(board *GameBoard, ship Ship) {
	if ship.isHorizontal {
		for i := 0; i < ship.size; i++ {
			field, _ := GetFieldByCords(board, Cords{
				X: ship.x + i,
				Y: ship.y,
			})

			field.isShip = true
		}
	} else {
		for i := 0; i < ship.size; i++ {
			field, _ := GetFieldByCords(board, Cords{
				X: ship.x,
				Y: ship.y + i,
			})

			field.isShip = true
		}
	}
}

// Функция для проверки, можно ли разместить корабль на указанной позиции
func canPlaceShip(board *GameBoard, ship Ship) bool {
	if ship.isHorizontal {
		if ship.x+ship.size-1 > BOARD_SIZE {
			return false
		}
	} else {
		if ship.y+ship.size-1 > BOARD_SIZE {
			return false
		}
	}

	for shift := -1; shift < 2; shift++ {
		for i := -1; i <= ship.size; i++ {
			var cords Cords

			if ship.isHorizontal {
				cords = Cords{
					X: ship.x + i,
					Y: ship.y + shift,
				}
			} else {
				cords = Cords{
					X: ship.x + shift,
					Y: ship.y + i,
				}
			}

			isProjectionsShip := shift == 0 && cords.X > 0 && cords.X < 11 && cords.Y > 0 && cords.Y < 11

			field, err := GetFieldByCords(board, cords)

			if err != nil {
				if isProjectionsShip {
					return false
				}
				continue
			}

			if field.isShip {
				return false
			}
		}
	}

	return true
}
