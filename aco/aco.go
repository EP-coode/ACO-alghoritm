package aco

import (
	"log"
	"math"
	"math/rand"
	"sync"

	"github.com/EP-coode/ACO-alghoritm/helpers"
)

type AcoParams struct {
	Alpha              float64
	Beta               float64
	Q                  float64
	D                  float64 // greater values makes small changes more significant
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

func (a *Aco) RunAco(iterations int) {
	l := log.Default()

	for i := 0; i < iterations; i++ {
		var antWg sync.WaitGroup
		antWg.Add(a.params.AntsPopulationSize)

		ants := make([]Ant, a.params.AntsPopulationSize)

		for j := 0; j < a.params.AntsPopulationSize; j++ {
			go func(j int) {
				startIndex := rand.Intn(len(*a.enviroment.cities))
				ants[j] = *a.antTraverse(startIndex)
				antWg.Done()
			}(j)
		}

		antWg.Wait()

		// update pheromone
		for _, ant := range ants {
			pheromoneDelta := math.Pow(a.params.Q/ant.Distance, a.params.D)

			for k := 1; k < len(ant.Path); k++ {
				city1 := ant.Path[k-1]
				city2 := ant.Path[k]
				currentPheromone, _ := a.enviroment.cityPheromones.GetEdge(city1, city2)
				a.enviroment.cityPheromones.SetEdge(city1, city2, *currentPheromone+pheromoneDelta)
			}

			if a.bestAnt == nil || ant.Distance < a.bestAnt.Distance {
				a.bestAnt = &Ant{
					Path:     ant.Path,
					Distance: ant.Distance,
				}
				l.Printf("New best Ant { distance: %v }", a.bestAnt.Distance)
			}
		}

		// evaporate pheromone
		for i := 0; i < len(*a.enviroment.cities); i++ {
			for j := i + 1; j < len(*a.enviroment.cities); j++ {
				currentPheromone, _ := a.enviroment.cityPheromones.GetEdge(i, j)
				a.enviroment.cityPheromones.SetEdge(i, j, *currentPheromone*a.params.DegradationFactor)
			}
		}

	}
}

func (a *Aco) GetBestAnt() *Ant {
	return &Ant{
		Path:     a.bestAnt.Path,
		Distance: a.bestAnt.Distance,
	}
}

func (a *Aco) antTraverse(startCityIndex int) *Ant {
	cities := *a.enviroment.cities

	if startCityIndex >= len(cities) || startCityIndex < 0 {
		return nil
	}

	ant := Ant{}

	ant.Path = make([]int, len(cities))
	ant.Path[0] = startCityIndex
	citiesToVisitIndexes := helpers.MakeRange(0, len(cities)-1)
	_, citiesToVisitIndexes = helpers.RemoveFromArray(citiesToVisitIndexes, startCityIndex)

	for i := 1; i < len(cities); i++ {
		// calc neightbour cities connection weights
		currentCity := ant.Path[i-1]
		citiesWeights := a.getNeightbourCitiesWeights(citiesToVisitIndexes, currentCity)

		// pick random city based on connection weight
		pickedCityIndex, _ := helpers.WeightRandomPick(citiesToVisitIndexes, citiesWeights)
		nextCity := citiesToVisitIndexes[*pickedCityIndex]
		ant.Path[i] = nextCity

		// update distance
		ant.Distance += *a.enviroment.GetCitiesDistance(currentCity, nextCity)

		// remove visited city
		_, citiesToVisitIndexes = helpers.RemoveFromArray(citiesToVisitIndexes, *pickedCityIndex)
	}

	ant.Distance += *a.enviroment.GetCitiesDistance(ant.Path[0], ant.Path[len(ant.Path)-1])

	return &ant
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
