package ryutil

import (
	"math/rand"
	"math"
	"time"
)

//normalRandom: Generate random numbers that follow a Normal distribution.

func NormRand(src int64) float64 {
	r := rand.New(rand.NewSource(src))
	normalx := r.NormFloat64()
	r = nil
	return normalx
}

func NormalRandomInRangeCof(src int64, min, max int64) int64 {
	var i = 0
	delta := max - min
	for ; i < 10000; i++ {
		normalx := math.Abs(NormRand(src))*float64(delta) + float64(min)
		if normalx > float64(min) && normalx < float64(max) {
			return int64(normalx)
		}
	}
	return min
}
func NormalRandomInRangeCoScaled(src int64, min, max int64, scale float64) int64 {
	var i = 0
	delta := max - min
	for ; i < 10000; i++ {
		normalx := math.Abs(NormRand(src))*float64(delta) + float64(min)
		if normalx > float64(min) && normalx < float64(max) {
			return int64(normalx)
		}
	}
	return min
}
func Rand4DigitGen() int64 {
	f := NormRand(time.Now().Unix()) * float64(10002910)
	return int64(math.Abs(f))%8999 + 1000
}
func NormalCos(seed, min, max int64, seed_lamda float64) int64 {
	delta := max - min
	y := math.Cos(float64(seed)*seed_lamda*math.Pi/180)*float64(delta/2) + float64(min) + float64(delta/2)
	z := math.Abs(NormRand(seed))*float64(delta) + float64(min)
	if rand.NormFloat64() > 0 {
		return int64(y)
	} else {
		return int64(z)
	}
}

//normalRandomInRange: Generate random numbers that follow a Normal distribution but are clipped to fit within a range
func NormalRandomInRange(src int64, min, max float64) float64 {
	var i = 0
	for ; i < 10000; i++ {
		normalx := math.Abs(NormRand(src))
		if normalx > min && normalx < max {
			return normalx
		}
	}
	return min
}

//normalRandomScaled: Generate random numbers that follow a Normal distribution with a given mean and standard deviation
func NormalRandomScaled(src int64, mean, stddev float64) float64 {
	r := NormalRandomInRange(src, -1, 1)
	r = r*stddev + mean
	return math.Round(r)
}

//lnRandomScaled: Generate random numbers that follow a Log-normal distribution with a given geometric mean and geometric standard deviation
func LnRandomScaled(src int64, gmean, gstddev float64) float64 {
	var r = NormalRandomInRange(src, -1, 1)
	r = r*math.Log(gstddev) + math.Log(gmean)
	return math.Round(math.Exp(r))
}

func Grn(src int64) float64 {
	r := rand.New(rand.NewSource(src))
	r.ExpFloat64()

	return r.NormFloat64()
}
