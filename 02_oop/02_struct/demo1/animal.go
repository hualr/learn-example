package demo1

import "fmt"

type Animal struct {
	scientificName string // 学名。
	AnimalCategory        // 动物基本分类。
}

func (a Animal) String() string {
	return fmt.Sprintf("%s (category: %s)",
		a.scientificName, a.AnimalCategory)
}
