package builder

type twoVehicle struct {
	body  string
	tyres string
}

func newTwoVehicle() *twoVehicle {
	return &twoVehicle{}
}

func (v *twoVehicle) buildBody() {
	v.body = "two-wheeler"
}

func (v *twoVehicle) buildTyres() {
	v.tyres = "two tyres"
}

func (v *twoVehicle) getVehicle() Vehicle {
	return Vehicle{
		body:  v.body,
		tyres: v.tyres,
	}
}
