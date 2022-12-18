package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/EP-coode/ACO-alghoritm/aco"
	"github.com/EP-coode/ACO-alghoritm/ttp"
)

func main() {
	rand.Seed(time.Now().Unix())
	runAco("trivial_0.ttp")
	runAco("easy_0.ttp")
	runAco("medium_0.ttp")
	runAco("hard_0.ttp")
}

func runAco(fileName string) {
	rootDir, _ := os.Getwd()
	path := filepath.Join(rootDir, "data", fileName)
	cities := ttp.LoadPorblem(path)
	acoParams := aco.AcoParams{
		Alpha:              1,
		Beta:               1,
		Q:                  1,
		DegradationFactor:  0.1,
		AntsPopulationSize: 20,
	}

	solver := aco.NewAco(acoParams, cities)
	solver.RunAco(5000)

	solver.PlotAnt(solver.GetBestAnt(), fmt.Sprintf("doc/results/result_of_%v", fileName))
}
