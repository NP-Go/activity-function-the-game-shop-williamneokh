package main

func main() {
	//Create objects
	Minecraft := game{title: "MineCraft", price: 5}
	WorldofWarcraft := game{title: "World of warcraft", price: 19}
	EliteDangerous := game{title: "Elite Dangerous", price: 54}

	Minecraft.printGameDetail()
	WorldofWarcraft.printGameDetail()
	EliteDangerous.printGameDetail()

	//inside above game struct into slice
	var gameList = list{}
	gameList = append(gameList, Minecraft, WorldofWarcraft, EliteDangerous)

	gameList.printAllGame()
}
