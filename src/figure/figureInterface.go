package figure

const AreaSize  = 30 // размер поля (поле квадратного типа)
type AreaType [AreaSize][AreaSize]int //Тип поля, где размещаются все фигуры
var Area *AreaType //Поле где размещаются все фигуры
type allFigInAreaType map[string]Painter
var allFigInArea allFigInAreaType

type Painter interface {
	Create (name string, coordX, coordY int, param ... int) error
	Delete (name string) error
}
