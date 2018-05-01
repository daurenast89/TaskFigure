package figure

import (
	"errors"
	"fmt"
)

// Описание фигуры Прямоугольник.
type Rectangle struct {
	Name     string
	AreaRect *AreaType
	Coord    [][2]int
}

//создание фигуры; на вход подаются начальные значения координат x,y и стороны прямоугольника
//при успешности создастся фигура, и метод вернет "nil"
// зачем name
func (r *Rectangle) Create(name string, coordX, coordY int, param ...int) error {
	tempArea := *r.AreaRect                    //создается слепок существующего поля
	coord := make([][2]int, param[0]*param[1]) // массив для хранения
	r.Coord = coord
	for i := coordX; i < i+param[0]; i++ {
		for j := coordY; j < j+param[1]; j++ {
			if tempArea[i][j] == 1 {
				return errors.New("Невозможно построить прямоугольник, на этих координатах есть другая фигура")
			} else {
				tempArea[i][j] = 1
			}
		}
	}
	fmt.Println(r.Coord)
	r.AreaRect = &tempArea
	return nil
}

//Удаление фигуры; на вход подается имя фигуры
func (r *Rectangle) Delete(name string) error {
	tempArea := *r.AreaRect              //создается слепок существующего поля
	allFigInArea:= make(map[string]Painter)
	if _, ok := allFigInArea[name]; ok { //проверка на наличие данных по ключю
		for _, v := range r.Coord { //Производим замену по координатам.
			tempArea[v[0]][v[1]] = 0
		}
		*r.AreaRect = tempArea
		delete(allFigInArea, name)
		return nil
	} else {
		return errors.New("Прямоугольник с таким именем не найден")
	}
}

//Движение фигуры; на вход подается имя фигуры; направление 1-вправо; кол-во ячеек
//															2-влево,
//															3-вверх
//															4-вниз

