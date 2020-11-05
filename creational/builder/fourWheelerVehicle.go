package builder

type fourVehicle struct {
	body  string
	tyres string
}

func newFourVehicle() *fourVehicle {
	return &fourVehicle{}
}

func (v *fourVehicle) buildBody() {
	v.body = "four-wheeler"
}

func (v *fourVehicle) buildTyres() {
	v.tyres = "four tyres"
}

func (v *fourVehicle) getVehicle() Vehicle {
	return Vehicle{
		body:  v.body,
		tyres: v.tyres,
	}
}
