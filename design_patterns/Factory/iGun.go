package main

// 产品接口 - 枪支模型定义
type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}
