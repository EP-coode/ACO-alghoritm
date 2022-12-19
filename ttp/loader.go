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

func LoadPorblem(filePath string) ([]aco.City, error) {
	file, err1 := os.Open(filePath)

	if err1 != nil {
		return nil, err1
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	dim, err2 := findDimSize(fileScanner)

	if err2 != nil {
		return nil, err2
	}

	cities, err3 := loadCities(fileScanner, dim)

	
	if err2 != nil {
		return nil, err3
	}

	file.Close()

	return cities, nil
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

		x, err1 := strconv.ParseFloat(tokens[1], 64)

		if err1 != nil {
			return nil, fmt.Errorf("failed to parse cities data: %v", err1)
		}

		y, err2 := strconv.ParseFloat(tokens[2], 64)

		if err2 != nil {
			return nil, fmt.Errorf("failed to parse cities data: %v", err2)
		}

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

	return 0, fmt.Errorf("dimension not found in header of file")
}
