package figure

const AreaSize  = 30 // размер поля (поле квадратного типа)
//type AreaType [AreaSize][AreaSize]int //Тип поля, где размещаются все фигуры

type Painter interface {
	Create (Area *[AreaSize][AreaSize]int, name string, coordX, coordY int, param ... int) error
	Delete (name string) error
}
