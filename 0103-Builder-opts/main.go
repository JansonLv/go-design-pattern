package main

func main() {
	normalHouse := NewHouse(setWindowType("normal"), setDoorType("normal"), setNumFloor(2))
	normalHouse.getHouse()

	// 根据选项设置默认模式
	iglooHouse := NewHouse(setWindowType("snow"))
	iglooHouse.getHouse()
}
