package aco

import (
	"math"

	"github.com/EP-coode/ACO-alghoritm/graph"
	"github.com/EP-coode/ACO-alghoritm/helpers"
)

type Ant struct {
	Path         []int
	Distance     float64
}

type AntEnviroment struct {
	cities         *[]City
	cityDistances  *graph.UndirectedNoLoopGraph[float64]
	cityPheromones *graph.UndirectedNoLoopGraph[float64]
}

type City struct {
	X float64
	Y float64
}

// TODO: improve error handling
// TODO: remove redundant code

func (src *City) GetDistance(dst *City) float64 {
	dx := src.X - dst.X
	dy := src.Y - dst.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (e *AntEnviroment) GetCitiesDistance(cityIndex1 int, cityIndex2 int) *float64 {
	value, _ := e.cityDistances.GetEdge(cityIndex1, cityIndex2)
	return value
}

func (e *AntEnviroment) GetCitiesPheromone(cityIndex1 int, cityIndex2 int) *float64 {
	value, _ := e.cityPheromones.GetEdge(cityIndex1, cityIndex2)
	return value
}

func (e *AntEnviroment) SetCitiesPheromone(cityIndex1 int, cityIndex2 int, newPheromoneLevel float64) {
	e.cityPheromones.SetEdge(cityIndex1, cityIndex2, newPheromoneLevel)
}

func NewAntEnviroment(cities []City) *AntEnviroment {
	cityIndexes := helpers.MakeRange(0, len(cities)-1)
	cityDistances := graph.NewGraph[int, float64](cityIndexes, 0)
	pheromones := graph.NewGraph[int, float64](cityIndexes, 1)

	enviroment := &AntEnviroment{
		cities:         &cities,
		cityDistances:  cityDistances,
		cityPheromones: pheromones,
	}

	for i := 0; i < len(cities); i++ {
		for j := i + 1; j < len(cities); j++ {
			cityDistance := cities[i].GetDistance(&cities[j])
			enviroment.cityDistances.SetEdge(i, j, cityDistance)
		}
	}

	return enviroment
}
