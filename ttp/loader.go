package ttp

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/EP-coode/ACO-alghoritm/aco"
)

var findDimesion = regexp.MustCompile(`(?m)DIMENSION:\s(\d+)`)

func LoadPorblem(filePath string) []aco.City {
	file, _ := os.Open(filePath)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	dim, _ := findDimSize(fileScanner)
	cities, _ := loadCities(fileScanner, dim)

	file.Close()

	return cities
}

func loadCities(scanner *bufio.Scanner, citiesCount int) ([]aco.City, error) {
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "NODE_COORD_SECTION") {
			break
		}
	}

	citiesLoaded := 0
	cities := make([]aco.City, citiesCount)

	for scanner.Scan() && citiesLoaded < citiesCount {
		line := scanner.Text()
		tokens := strings.Split(line, "\t")
		x, _ := strconv.ParseFloat(tokens[1], 64)
		y, _ := strconv.ParseFloat(tokens[2], 64)

		cities[citiesLoaded] = aco.City{
			X: x,
			Y: y,
		}
		citiesLoaded++
	}

	if citiesLoaded != citiesCount {
		return cities, fmt.Errorf("loaded %v cities. Can not find declared %v cities", cities, citiesCount)
	}

	return cities, nil
}

func findDimSize(scanner *bufio.Scanner) (int, error) {
	for scanner.Scan() {
		line := scanner.Text()
		match := findDimesion.FindStringSubmatch(line)
		if len(match) == 2 {
			dim, _ := strconv.Atoi(match[1])
			return dim, nil
		}
	}

	return 0, fmt.Errorf("dimension not found")
}
