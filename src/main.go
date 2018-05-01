package main

import (
	"fmt"
	"figure"
)

func main(){
	NewReactangle := new(figure.Rectangle)
	Area:=*figure.Area
	for i:=0; i < figure.AreaSize; i++ {
		for j := 0; j < figure.AreaSize; j++ {
			Area[i][j] = 0
		}
	}


	NewReactangle.Create("Rect",10, 10, 5,4)
	for _, v := range figure.Area { //Вывод Нашего поля на экран
		fmt.Println(v)
	}
}
