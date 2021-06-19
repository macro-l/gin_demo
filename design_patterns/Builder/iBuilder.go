package main

// 生成器接口  - 房屋
type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

// 房屋生成器
func getBuilder(builderType string) iBuilder {
	if builderType == "normal" {
		return &normalBuilder{}
	}
	if builderType == "igloo" {
		return &iglooBuilder{}
	}
	return nil
}
