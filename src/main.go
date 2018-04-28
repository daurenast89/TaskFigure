package main

import (
	"fmt"
	"TaskFigure/figure"
)



type Area [figure.AreaSize][figure.AreaSize]int //Поле, где размещаются все фигуры



func main(){
	allFigInArea := make(map[string]figure.Painter) //хранение всех существующих фигур

	RectCreate:=new(rectangle)
	//RectCreate.Create("test", 1,2,3,4)
	fmt.Println(RectCreate.Create("test", 1,2,3,4))
}
