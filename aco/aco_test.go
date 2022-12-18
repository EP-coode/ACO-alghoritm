package aco

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestEnviroemntCreation(t *testing.T) {
	cities := []City{
		{X: 1, Y: 3},
		{X: 2, Y: 2},
		{X: 3, Y: 1},
	}

	enviroment := NewAntEnviroment(cities)

	distance := enviroment.GetCitiesDistance(0, 1)

	if *distance-math.Sqrt2 > 0.0001 {
		t.Errorf("Wrong city distance computed betwen city 0 and 1. Expected %v but got %v",
			math.Sqrt2, distance)
	}

	distance = enviroment.GetCitiesDistance(1, 0)

	if *distance-math.Sqrt2 > 0.0001 {
		t.Errorf("Wrong city distance computed betwen city 1 and 0. Expected %v but got %v",
			math.Sqrt2, distance)
	}

	distance = enviroment.GetCitiesDistance(2, 0)

	if *distance-math.Sqrt2*2 > 0.0001 {
		t.Errorf("Wrong city distance computed betwen city 2 and 0. Expected %v but got %v",
			math.Sqrt2, distance)
	}

	distance = enviroment.GetCitiesDistance(2, 2)

	if distance != nil {
		t.Errorf("Sloud not return connection to itself")
	}

	distance = enviroment.GetCitiesDistance(0, 3)

	if distance != nil {
		t.Errorf("Sloud not return if city is out of range")
	}
}

func TestAntTraversing(t *testing.T) {
	cities := []City{
		{X: 1, Y: 3},
		{X: 2, Y: 2},
		{X: 3, Y: 1},
	}

	acoParams := AcoParams{
		Alpha:              1,
		Beta:               1,
		Q:                  1,
		DegradationFactor:  1,
		AntsPopulationSize: 10,
	}

	rand.Seed(time.Now().Unix())

	aco := NewAco(acoParams, cities)

	ant := aco.antTraverse(0)

	if len(ant.Path) != 3 {
		t.Errorf("Wrong path lenght. Expected 3 but got %v", len(ant.Path))
	}
}
