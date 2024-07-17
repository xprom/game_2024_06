package main

import (
	"encoding/json"
	"fmt"
	term "github.com/nsf/termbox-go"
	"os"
	"time"
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
	/// сделал много изменений тут
	/// правки никто еще не видел и не проверял, я хочу просто их показать
	///
	///
	/// ////zz
	for _, obsticel := range gameMap.Obsticles {
		if obsticel.X == x && obsticel.Y == y {
			return false
		}
	}

	return true
}

func drawGameField(playerPositionX, playerPositionY int, gameMap Map, bullets []*Bullet) {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			isBulletsFound := false
			for _, b := range bullets {
				if b.x == i && b.y == j {
					term.SetCell(j, i, '.', 0, term.ColorLightBlue)
					isBulletsFound = true
				}
			}

			if isBulletsFound {
				continue
			}

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

type Bullet struct {
	x         int
	y         int
	direction string // пока не используем
}

// обновляет позицию пули
func (b *Bullet) fly() bool {
	b.y = b.y + 1
	if b.y > 100 {
		return false
	}

	return true
}

func main() {
	bulletesOnFiled := make([]*Bullet, 0)

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
	// стартовые координаты игрока
	x := 0
	y := 0
	drawGameField(x, y, gameMap, bulletesOnFiled)
	if err := term.Sync(); err != nil {
		panic(err)
	}

	// запускаем фоновый процесс обновления позиций пуль на карте
	go func() {
		for {
			time.Sleep(time.Millisecond * 100)
			if len(bulletesOnFiled) == 0 {
				continue
			}

			for _, b := range bulletesOnFiled {
				b.fly()
			}
			drawGameField(x, y, gameMap, bulletesOnFiled)
			if err := term.Sync(); err != nil {
				panic(err)
			}
		}
	}()

	for {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyArrowRight:
				y++
				drawGameField(x, y, gameMap, bulletesOnFiled)
				if err := term.Sync(); err != nil {
					panic(err)
				}
			case term.KeyArrowLeft:
				y--
				drawGameField(x, y, gameMap, bulletesOnFiled)
				if err := term.Sync(); err != nil {
					panic(err)
				}
			case term.KeyArrowDown:
				x++
				drawGameField(x, y, gameMap, bulletesOnFiled)
				if err := term.Sync(); err != nil {
					panic(err)
				}
			case term.KeyArrowUp:
				x--
				drawGameField(x, y, gameMap, bulletesOnFiled)
				if err := term.Sync(); err != nil {
					panic(err)
				}
			case term.KeySpace:
				bulletesOnFiled = append(bulletesOnFiled, &Bullet{
					x:         x,
					y:         y,
					direction: "",
				})
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
