package randomanimals

import (
	"math/rand"
	"time"
)

// Генерирует случайное животное
func RandomAnimal() string {
	numbs := len(animals)
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	randomIndex := r.Intn(numbs)
	return animals[randomIndex]
}
