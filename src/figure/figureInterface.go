package figure

const AreaSize  = 30 // размер поля (поле квадратного типа)

type Painter interface {
	Create (Area *[AreaSize][AreaSize]int, allFigInArea map[string]Painter,name string, coordX, coordY int, param ...int) error
	Delete(name string, allFigInArea map[string]Painter) error
}
