package main

import (
	"fmt"
	"github.com/vikashkumarjha/golang_design_patterns/creational/builder"
)

func main() {
	b := builder.GetVehicle("twowheeler")

	dir := builder.NewDirector(b)

	v := dir.BuildVehicle()

	fmt.Printf("%v", v)

}
