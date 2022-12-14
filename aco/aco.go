package aco

import (
	"math"
)

type AcoParams struct {
}

type AntEnviroment struct {
	Cities         []*City
	CityConections [][]*CityConnection
}

type CityConnection struct {
	Pheromone float64
	Distance  float64
}

type City struct {
	X float64
	Y float64
}

func (src *City) GetDistance(dst *City) float64 {
	dx := src.X - dst.X
	dy := src.Y - dst.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (e *AntEnviroment) GetConnectionInfo(cityIndex1 int, cityIndex2 int) *CityConnection {
	if cityIndex2 == cityIndex1 {
		return nil
	}

	if cityIndex2 >= len(e.Cities) || cityIndex2 < 0 {
		return nil
	}

	if cityIndex1 >= len(e.Cities) || cityIndex1 < 0 {
		return nil
	}

	// make sure index 2 is bigger than 1
	if cityIndex1 > cityIndex2 {
		tmp := cityIndex1
		cityIndex1 = cityIndex2
		cityIndex2 = tmp
	}

	// add ofset becaouse diagonal is skipped
	return e.CityConections[cityIndex2 - 1][cityIndex1]
}

func NewAntEnviroment(cities []*City) *AntEnviroment {
	cityConnections := make([][]*CityConnection, len(cities)-1)
	for i := range cityConnections {
		cityConnections[i] = make([]*CityConnection, len(cityConnections)-i)
	}

	enviroment := &AntEnviroment{
		Cities:         cities,
		CityConections: cityConnections,
	}

	for i := range cityConnections {
		for j := range cityConnections[i] {
			cityConnections[i][j] = &CityConnection{
				Pheromone: 0,
				Distance:  cities[i].GetDistance(cities[j]),
			}
		}
	}

	return enviroment
}
