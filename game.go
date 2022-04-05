package main

import "fmt"

//Declare Struct for Game
type game struct {
	title string
	price float64
}

//Declare method for Game
func (g game) printGameDetail() {

	fmt.Printf("Game Title: %v - Price: $%v\n", g.title, g.price)
}
