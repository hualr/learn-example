package _1_bean

import (
	"fmt"
	"testing"
)

type AnimalCategory struct {
	kingdom string // 界。
	phylum  string // 门。
	class   string // 纲。
	order   string // 目。
	family  string // 科。
	genus   string // 属。
	species string // 种。
}

func (ac AnimalCategory) String() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s",
		ac.kingdom, ac.phylum, ac.class, ac.order,
		ac.family, ac.genus, ac.species)
}

type Animal struct {
	scientificName string // 学名。
	AnimalCategory        // 动物基本分类。
}

func (a Animal) String() string {
	return fmt.Sprintf("%s (category: %s)",
		a.scientificName, a.AnimalCategory)
}

type Cat struct {
	name string
	Animal
}

// 该方法会"屏蔽"掉嵌入字段中的同名方法。
func (cat Cat) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q)",
		cat.scientificName, cat.Animal.AnimalCategory, cat.name)
}

// 永远在覆盖
func Test5(t *testing.T) {
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

func Test6(t *testing.T) {

}
