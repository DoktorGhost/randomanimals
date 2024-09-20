package randomanimals

import (
	"math/rand"
)

// Генерирует случайное животное
func RandomAnimal() string {
	numbs := len(animals)
	randomIndex := rand.Intn(numbs)
	return animals[randomIndex]
}
