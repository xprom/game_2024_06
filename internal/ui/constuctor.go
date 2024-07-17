package ui

import (
	"fmt"
	"math/rand"
)

func DrawHelloMessage() {
	fmt.Println("Привет " + helperRandomName())
}

func helperRandomName() string {
	switch rand.Intn(3) {
	case 0:
		return "Иван"
	case 1:
		return "Семен"
	case 2:
		return "Игорь"
	}

	panic("недостижимое значение")
}
