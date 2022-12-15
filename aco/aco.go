package aco

import (
	"math"

	"github.com/EP-coode/ACO-alghoritm/helpers"
)

type AcoParams struct {
	Alpha              float64
	Beta               float64
	Q                  float64
	DegradationFactor  float64
	AntsPopulationSize int
}

type Aco struct {
	params     AcoParams
	enviroment *AntEnviroment
	bestAnt    *Ant
}

func NewAco(params AcoParams, cities []City) *Aco {
	return &Aco{
		params:     params,
		bestAnt:    nil,
		enviroment: NewAntEnviroment(cities),
	}
}

func (a *Aco) RunAco() {

}

func (a *Aco) GetBestAnt() *Ant {
	return &Ant{
		Path: a.bestAnt.Path,
	}
}

func (a *Aco) antTraverse(startCityIndex int) *Ant {
	cities := *a.enviroment.cities

	if startCityIndex >= len(cities) || startCityIndex < 0 {
		return nil
	}

	visitedCitiesIndexes := make([]int, len(cities))
	visitedCitiesIndexes[0] = startCityIndex
	citiesToVisitIndexes := helpers.MakeRange(0, len(cities)-1)
	_, citiesToVisitIndexes = helpers.RemoveFromArray(citiesToVisitIndexes, startCityIndex)

	for i := 1; i < len(cities); i++ {
		// calc neightbour cities connection weights
		currentCityIndex := visitedCitiesIndexes[i-1]
		citiesWeights := a.getNeightbourCitiesWeights(citiesToVisitIndexes, currentCityIndex)

		// pick random city based on connection weight
		pickedCityIndex, _ := helpers.WeightRandomPick(citiesToVisitIndexes, citiesWeights)
		visitedCitiesIndexes[i] = citiesToVisitIndexes[*pickedCityIndex]

		// remove visited city
		_, citiesToVisitIndexes = helpers.RemoveFromArray(citiesToVisitIndexes, *pickedCityIndex)
	}

	return &Ant{
		Path: visitedCitiesIndexes,
	}
}

func (a *Aco) getNeightbourCitiesWeights(citiesToVisitIndexes []int, currentCityIndex int) []float64 {
	citiesWeights := make([]float64, len(citiesToVisitIndexes))

	for j, cityToVisitIndex := range citiesToVisitIndexes {
		pheromoneLevel := a.enviroment.GetCitiesPheromone(currentCityIndex, cityToVisitIndex)
		distance := a.enviroment.GetCitiesDistance(currentCityIndex, cityToVisitIndex)
		citiesWeights[j] = math.Pow(*pheromoneLevel, a.params.Alpha) / math.Pow(*distance, a.params.Beta)
	}

	return citiesWeights
}
