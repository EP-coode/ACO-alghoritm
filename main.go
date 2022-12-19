package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/EP-coode/ACO-alghoritm/aco"
	"github.com/EP-coode/ACO-alghoritm/ttp"
)

func main() {
	rand.Seed(time.Now().Unix())
	start := time.Now()
	log.Print("Starting trivial_0")
	runAco("trivial_0.ttp")
	log.Print("Starting easy_0")
	runAco("easy_0.ttp")
	log.Print("Starting hard_0")
	runAco("hard_0.ttp")
	elapsed := time.Since(start)
	log.Printf("Execution took %s", elapsed)
}

func runAco(fileName string) {
	rootDir, _ := os.Getwd()
	path := filepath.Join(rootDir, "data", fileName)
	cities, err := ttp.LoadPorblem(path)

	if err != nil {
		log.Printf("Failed to load file %v. cause: %v", fileName, err)
		return
	}

	acoParams := aco.AcoParams{
		Alpha:              1,
		Beta:               2,
		Q:                  20_000,
		D:                  13,
		DegradationFactor:  0.3,
		AntsPopulationSize: 40,
	}

	solver := aco.NewAco(acoParams, cities)
	solver.RunAco(5_000)

	solver.PlotAnt(solver.GetBestAnt(), fmt.Sprintf("doc/results/result_of_%v", fileName))
}
