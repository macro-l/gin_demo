package Extends

import "fmt"

// 动物 - 类
type Animal struct {
	name string
}

// （动物）叫声
func (a *Animal) Call() string {
	return "叫声"
}

// （动物）名称
func (a *Animal) GetName() string {
	return a.name
}

// 狗狗 - 类 (继承 动物类)
type Dog struct {
	Animal
}

// 引用结构比值结构内存消耗小
//type Dog struct {
//	animal *Animal
//	name   string
//}

// （狗狗）叫声
func (d *Dog) Call() string {
	return "汪汪汪"
}

// 导盲犬 - 类
type GuideDog struct {
	Animal
	Dog
}

// (导盲犬) 叫声
func (gd *GuideDog) Call() string {
	return "小声汪汪"
}

func main() {
	// 实例化+填充
	animal := Animal{"狗狗"}
	dog := Dog{animal}
	fmt.Println(dog.GetName(), "叫声：", dog.Call())
	// 父属性输出
	fmt.Println(dog.Animal.GetName(), "叫声：", dog.Animal.Call())
	fmt.Println(animal.GetName(), "叫声：", animal.Call())

	// 实例化
	dog2 := Dog{}
	fmt.Println(dog2.Call())

	// 导盲犬 （多继承测试）
	// 组合的类型中包含同名方法，如果子类没有重写，调用的时候会报错
	// 不重写重复的同名方法时， 报错信息：ambiguous selector guideDog.Call
	guideDog := GuideDog{}
	fmt.Println(guideDog.Call())

}
