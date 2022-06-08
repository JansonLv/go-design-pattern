package main

func main() {
	house1 := NewHouse()
	normalHouse := house1.setDoorType("normal door").setNumFloor(100).setWindowType("normal window")
	normalHouse.getHouse()

	house2 := NewHouse()
	iglooHouse := house2.setDoorType("Snow door").setNumFloor(1).setWindowType("Snow window")
	iglooHouse.getHouse()

	//  默认房子
	house3 := NewDefaultHouse().setNumFloor(2)
	house3.getHouse()
}
