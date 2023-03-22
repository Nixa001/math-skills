package functions

import (
	"math"
	"sort"
)

func Average(nombres []float64) float64 {
	sum := 0.0
	for _, num := range nombres {
		sum += num
	}
	average := sum / float64(len(nombres))

	return average
}

func Median(nombres []float64) float64 {
	sort.Float64s(nombres)
	// fmt.Println((nombres))
	milieu := (len(nombres)) / 2
	if (len(nombres))%2 == 0 {
		return (nombres[milieu-1] + nombres[milieu]) / 2
	}

	return nombres[milieu]
}

func Variance(nombres []float64) float64 {
	variance := 0.0
	average := Average(nombres)

	for _, num := range nombres {
		variance += math.Pow(num-average, 2)
	}
	variance = variance / float64(len(nombres))

	return variance
}

func StandardDeviation(nombres []float64) float64 {
	variance := Variance(nombres)
	ecartType := math.Sqrt(variance)
	return ecartType
}
