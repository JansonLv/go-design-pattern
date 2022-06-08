package main

import "fmt"

type house struct {
	windowType string
	doorType   string
	floor      int
}

func NewHouse(opts ...OptionFunc) *house {
	h := &house{
		windowType: "default",
		doorType:   "default",
		floor:      0,
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

type OptionFunc func(*house)

func setWindowType(windowType string) OptionFunc {
	return func(h *house) {
		h.windowType = windowType
	}
}

func setDoorType(doorType string) OptionFunc {
	return func(h *house) {
		h.doorType = doorType
	}
}

func setNumFloor(floorNum int) OptionFunc {
	return func(h *house) {
		h.floor = floorNum
	}
}

func (house *house) getHouse() {
	fmt.Printf("\nIgloo House Door Type: %s\n", house.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", house.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", house.floor)
}
