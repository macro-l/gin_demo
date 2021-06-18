package main

import "fmt"

// 抽象工厂接口 - 运动装备的牌子工厂
type iSportsFactory interface {
	makeShoe() iShoe
	makeShirt() iShirt
}

func getSportsFactory(brand string) (iSportsFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}
	if brand == "nike" {
		return &nike{}, nil
	}

	return nil, fmt.Errorf("wrong brand type passed")
}
