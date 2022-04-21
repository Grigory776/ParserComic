package main

import (
	"ParserComic/xkcd"
	"fmt"
	"time"
)

func main(){
	
	for {
		fmt.Println("Введите номер коммикса (1-2600): ")
		var num string
		fmt.Scan(&num)
		tmp, er:= xkcd.SearchComic(num)
		if er!=nil{
			panic(er)
		}
		(*tmp).Print()
		fmt.Println("Требуется ли еще информация по коммиксам? (да/нет)")
		var r string
		fmt.Scan(&r)
		if r == "нет"{
			break
		} else if r == "да"{
			for i:=0; i<10; i++{
				fmt.Println()
			}
		} else {
			fmt.Println("Вы ввели неверную комманду! ")
			time.Sleep(3 * time.Second)
			break
		}
}
}

