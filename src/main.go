package main

import (
	"fmt"
	"figure"
)


func main() {
	var Area [figure.AreaSize][figure.AreaSize]int //Поле где размещаются все фигуры
	for i := 0; i < figure.AreaSize; i++ { //Запольняем поле Нолями
	for j := 0; j < figure.AreaSize; j++ {
		Area[i][j] = 0
	}
}
	allFigInArea:= make(map[string]figure.Painter) //карта для хранения фигур
	var err error // переменная для ошибок
	nameRect:= "Rect1"
	if _, ok := allFigInArea[nameRect]; ok {

	} else {
		NewReactangle := new(figure.Rectangle)// Тут описывается движение фигуры
		err = NewReactangle.Create(&Area, "Rect", 10, 10, 5, 4)
		if err != nil {
			fmt.Println(err)
		}
		for _, v := range Area { //Вывод Нашего поля на экран
			fmt.Println(v)
		}
	}
}
