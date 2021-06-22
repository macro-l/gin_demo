package animal

type Animal struct {
	name string
}

func NewAnimal(name string) *Animal {
	return &Animal{name: name}
}

func (a Animal) Call() string {
	return "动物的叫声..." + a.happy()
}

func (a Animal) FavorFood() string {
	return "爱吃的食物..."
}

func (a Animal) GetName() string {
	return a.name
}

func (a Animal) happy() string {
	return "哈哈汪~~"
}

func (a Animal) Unhappy() string {
	return "滋滋汪~~"
}
