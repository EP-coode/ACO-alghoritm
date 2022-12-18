package main

import (
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/EP-coode/ACO-alghoritm/aco"
	"github.com/EP-coode/ACO-alghoritm/ttp"
)

func main() {
	rand.Seed(time.Now().Unix())
	rootDir, _ := os.Getwd()
	path := filepath.Join(rootDir, "data", "trivial_0.ttp")
	cities := ttp.LoadPorblem(path)
	acoParams := aco.AcoParams{
		Alpha:              1,
		Beta:               1,
		Q:                  1,
		DegradationFactor:  0.1,
		AntsPopulationSize: 10,
	}

	solver := aco.NewAco(acoParams, cities)
	solver.RunAco(100)

}
