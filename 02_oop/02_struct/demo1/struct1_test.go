package demo1

import (
	"fmt"
	"testing"
)

// 永远在覆盖
func Test1(t *testing.T) {
	category := AnimalCategory{species: "cat"}
	fmt.Printf("The animal category: %s\n", category)

	animal := Animal{
		scientificName: "American Shorthair",
		// 这里注意到 这里需要设置!
		AnimalCategory: category,
	}
	fmt.Printf("The animal: %s\n", animal)

	cat := Cat{
		name:   "little pig",
		Animal: animal,
	}
	fmt.Printf("The cat: %s\n", cat)
}
