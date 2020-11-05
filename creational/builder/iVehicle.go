package builder

type iVehicle interface {
	buildBody()
	buildTyres()
	getVehicle() Vehicle
}

func GetVehicle(vechicleType string) iVehicle {
	switch vechicleType {
	case "twowheeler":
		return &twoVehicle{}

	case "fourWheeler":
		return &fourVehicle{}
	}
	return nil

}
