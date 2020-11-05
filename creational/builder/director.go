package builder

type director struct {
	builder iVehicle
}

func NewDirector(b iVehicle) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iVehicle) {
	d.builder = b
}

func (d *director) BuildVehicle() Vehicle {
	d.builder.buildBody()
	d.builder.buildTyres()
	return d.builder.getVehicle()
}
