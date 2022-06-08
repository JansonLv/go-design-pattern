package main

import "fmt"

type house struct {
	windowType string
	doorType   string
	floor      int
}

func NewHouse() *house {
	return &house{}
}

func NewDefaultHouse() *house {
	return &house{
		windowType: "normal",
		doorType:   "nomal",
		floor:      1,
	}
}

func (house *house) setWindowType(windowType string) *house {
	house.windowType = windowType
	return house
}

func (house *house) setDoorType(doorType string) *house {
	house.doorType = doorType
	return house
}

func (house *house) setNumFloor(floorNum int) *house {
	house.floor = floorNum
	return house
}

func (house *house) getHouse() {
	fmt.Printf("\nIgloo House Door Type: %s\n", house.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", house.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", house.floor)
}
