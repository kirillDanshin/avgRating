// Package avgRating calculates average score and rating based on Wilson Score Equation
// Refer: https://en.wikipedia.org/wiki/Binomial_proportion_confidence_interval
//
// Port of https://github.com/ndaidong/average-rating
package avgRating

import (
	"math"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

// GetWilsonScore returns Wilson Score
func GetWilsonScore(p, n float64) float64 {
	if p == 0 && n == 0 {
		return 0
	}

	return toFixed(((p+1.9208)/(p+n)-1.96*math.Sqrt(p*n/(p+n)+0.9604)/(p+n))/(1+3.8416/(p+n)), 2)
}

// Rate calculates the rate based on given rating
func Rate(rating [5]int) float64 {
	n := 0.0
	p := 0.0
	n += float64(rating[0])

	n += float64(rating[1]) * 0.75
	p += float64(rating[1]) * 0.25

	n += float64(rating[2]) * 0.5
	p += float64(rating[2]) * 0.5

	n += float64(rating[3]) * 0.25
	p += float64(rating[3]) * 0.75

	p += float64(rating[4])

	return GetWilsonScore(p, n)
}

// Average rating
func Average(rating [5]int) float64 {
	total := 0.0
	sum := 0.0
	k := 1.0
	for _, item := range rating {
		total += float64(item)
		sum += float64(item) * k
		k++
	}

	if total == 0 {
		return 0
	}

	return toFixed(sum/total, 1)
}
