package figure

import (
	"errors"
	"fmt"
)

const AreaSize  = 30 // размер поля (поле квадратного типа)
// Описание фигуры Прямоугольник.
type Rectangle struct {
	Name        string
	Area        *[AreaSize][AreaSize]int
	Coord [][2]int
	Height      int //высота  прямоугольника.
	Width       int //ширина  прямоугольника.
}

func (r *Rectangle) Create (coordX, coordY int) error {
	tempArea := *r.Area
	coord := make([][2]int, r.Height*r.Width)
	r.Coord = coord
	for i := coordX; i < i+r.Height; i++ {
		for j := coordY; j < j+r.Width; j++ {
			if tempArea[i][j] == 1 {
				return errors.New("Невозможно построить фигуру, на этих координатах есть другая фигура")
			}else{
				tempArea[i][j] = 1
			}
		}
	}
	fmt.Println(r.Coord)
	*r.Area = tempArea
	return nil
}
