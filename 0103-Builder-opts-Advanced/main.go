package main

func main() {
	normalHouse := NewHouse(createMaterialByCementName("海螺"), setWindowType("normal"), setDoorType("normal"), setNumFloor(2))
	normalHouse.getHouse()

	cementSlab := createCementSlab("普通")
	iglooHouse := NewHouse(createMaterialByMaterial(cementSlab))
	iglooHouse.getHouse()
}
