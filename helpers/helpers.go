package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func WeightRandomPick[T any](values []T, weights []float64) (*int, error) {
	if len(values) != len(weights) {
		return nil, fmt.Errorf("expected arrays with same len but got: %v and %v", len(values), len(weights))
	}

	totalWeight := 0.

	for _, weight := range weights {
		if weight > 0 {
			totalWeight += weight
		}
	}

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	stopWeight := r.Float64() * totalWeight

	totalTraversedWeight := 0.
	stopIndex := 0

	for i, weight := range weights {
		if weight > 0 {
			totalTraversedWeight += weight
			if totalTraversedWeight > stopWeight {
				stopIndex = i
				break
			}
		}
	}

	return &stopIndex, nil
}

func RemoveFromArray[T any](arr []T, index int) (T, []T) {
	arr[index] = arr[len(arr)-1]
	removed := arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	return removed, arr
}
