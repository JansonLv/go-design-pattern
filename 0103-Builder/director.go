package main

type director struct {
	builder builder
}

func newDirector(b builder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b builder) {
	d.builder = b
}

func (d *director) buildHouse() house {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}
