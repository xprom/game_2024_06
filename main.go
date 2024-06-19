package main

import (
	"encoding/json"
	"fmt"
	"os"

	term "github.com/nsf/termbox-go"
)

type Map struct {
	Title     string     `json:"title"`
	Obsticles []Position `json:"obsticles"`
}

// привет, Анастасия
// привет
type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// isClearPoint проверяем есть ли препятствие на указанной клетке
func isClearPoint(x, y int, gameMap Map) bool {
	for _, obsticel := range gameMap.Obsticles {
		if obsticel.X == x && obsticel.Y == y {
			return false
		}
	}

	return true
}

func drawGameField(playerPositionX, playerPositionY int, gameMap Map) {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if playerPositionX == i && playerPositionY == j {
				term.SetCell(j, i, '*', 0, term.ColorGreen)
			} else {
				if isClearPoint(i, j, gameMap) {
					term.SetCell(j, i, ' ', 0, term.ColorLightBlue)
				} else {
					term.SetCell(j, i, ' ', 0, term.ColorLightRed)
				}
			}
		}
	}
}

func main() {
	file, err := os.Open("map.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := make([]byte, 1000)
	n, err := file.Read(data)

	var gameMap Map
	err = json.Unmarshal(data[0:n], &gameMap)

	fmt.Println("game started")
	err = term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()
	x := 0
	y := 0
	drawGameField(x, y, gameMap)
	if err := term.Sync(); err != nil {
		panic(err)
	}

	for {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyArrowRight:
				y++
				drawGameField(x, y, gameMap)
				if err := term.Sync(); err != nil {
					panic(err)
				}
			case term.KeyArrowLeft:
				y--
				drawGameField(x, y, gameMap)
				if err := term.Sync(); err != nil {
					panic(err)
				}
			case term.KeyArrowDown:
				x++
				drawGameField(x, y, gameMap)
				if err := term.Sync(); err != nil {
					panic(err)
				}
			case term.KeyArrowUp:
				x--
				drawGameField(x, y, gameMap)
				if err := term.Sync(); err != nil {
					panic(err)
				}
			default:
				term.Clear(0, 0)
				fmt.Println("game finished")
				os.Exit(0)
			}
		case term.EventError:
			panic(ev.Err)
		}
	}
}
