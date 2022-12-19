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
	cities := ttp.LoadPorblem(path)
	acoParams := aco.AcoParams{
		Alpha:              1,
		Beta:               1,
		Q:                  10_000,
		DegradationFactor:  0.1,
		AntsPopulationSize: 10,
	}

	solver := aco.NewAco(acoParams, cities)
	solver.RunAco(5000)

	solver.PlotAnt(solver.GetBestAnt(), fmt.Sprintf("doc/results/result_of_%v", fileName))
}
