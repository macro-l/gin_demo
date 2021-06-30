package main

// 主管程序
type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildHouse() house {
	setDoorType()
	setWindowType()
	setNumFloor()
	return getHouse()
}
