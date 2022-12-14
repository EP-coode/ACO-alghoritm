package tests

import (
	"math"
	"testing"

	"github.com/EP-coode/ACO-alghoritm/aco"
)

func TestEnviroemntCreation(t *testing.T) {
	cities := []*aco.City{
		{X: 1, Y: 3},
		{X: 2, Y: 2},
		{X: 3, Y: 1},
	}

	enviroment := aco.NewAntEnviroment(cities)

	connection := enviroment.GetConnectionInfo(0, 1)

	if connection.Distance-math.Sqrt2 > 0.0001 {
		t.Errorf("Wrong city distance computed betwen city 0 and 1. Expected %v but got %v",
			math.Sqrt2, connection.Distance)
	}

	connection = enviroment.GetConnectionInfo(1, 0)

	if connection.Distance-math.Sqrt2 > 0.0001 {
		t.Errorf("Wrong city distance computed betwen city 1 and 0. Expected %v but got %v",
			math.Sqrt2, connection.Distance)
	}

	connection = enviroment.GetConnectionInfo(2, 0)

	if connection.Distance-math.Sqrt2*2 > 0.0001 {
		t.Errorf("Wrong city distance computed betwen city 2 and 0. Expected %v but got %v",
			math.Sqrt2, connection.Distance)
	}

	connection = enviroment.GetConnectionInfo(2, 2)

	if connection != nil {
		t.Errorf("Sloud not return connection to itself")
	}

	connection = enviroment.GetConnectionInfo(0, 3)

	if connection != nil {
		t.Errorf("Sloud not return if city is out of range")
	}
}
