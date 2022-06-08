package main

import (
	"fmt"
)

type house struct {
	material   cementSlab // 假设材料只有水泥或者是水泥板
	windowType string
	doorType   string
	floor      int
}

func NewHouse(createMaterialFunc createMaterialFunc, opts ...OptionFunc) *house {
	h := &house{
		windowType: "default",
		doorType:   "default",
		floor:      0,
	}
	if createMaterialFunc != nil {
		h.material = createMaterialFunc()
	}

	for _, opt := range opts {
		opt(h)
	}
	return h
}

type createMaterialFunc func() cementSlab

// 根据提供的水泥名称直接制成水泥板
func createMaterialByCementName(cementName string) func() cementSlab {
	return func() cementSlab {
		return createCementSlab(cementName)
	}
}

func createMaterialByMaterial(material cementSlab) func() cementSlab {
	// 直接使用水泥板
	return func() cementSlab {
		return material
	}
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
	fmt.Printf("House material: %s \n", house.material)
	fmt.Printf("Igloo House Door Type: %s\n", house.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", house.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n\n", house.floor)
}
