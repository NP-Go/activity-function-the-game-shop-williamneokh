package main

import "fmt"

//declare List
type list []game

//declare method for List Print

func (l list) printAllGame() {
	for _, item := range l {
		fmt.Printf("Game title: %v - Price: $%v\n", item.title, item.price)
	}
}
