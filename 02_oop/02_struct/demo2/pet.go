package demo2

type Pet interface {
	SetName(name string)
	Name() string
	Category() string
	ScientificName() string
	test() Pet
}
