package aoc

import "math"

// Bearing returns the degrees clockwise from the +Y axis to a given point.
func Bearing(x, y int) float64 {
	return math.Mod(360-math.Atan2(float64(-x), float64(-y))*(180/math.Pi), 360)
}

// Max returns the maximum value of two given integers.
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Min returns the minimum value of two given integers.
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Abs returns the absolute value of a given integer.
func Abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

// Lcm returns the least common multiple of two integers.
func Lcm(x, y int) int {
	return x * y / Gcd(x, y)
}

// Gcd returns the greatest common denominator of two integers.
func Gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// Ceil wraps math.Ceil() in a cast to int.
func Ceil(x float64) int {
	return int(math.Ceil(x))
}

// IntSliceMax returns the largest integer in a []int.
func IntSliceMax(ints []int) int {
	var max int
	for _, i := range ints {
		if i > max {
			max = i
		}
	}
	return max
}

// IntSliceMin returns the smallest integer in a []int.
func IntSliceMin(ints []int) int {
	min := math.MaxInt64
	for _, i := range ints {
		if i < min {
			min = i
		}
	}
	return min
}

// InitIntSlice returns a new []int with all integers between min and max, inclusive.
func InitIntSlice(min, max int) []int {
	ints := []int{}

	for i := min; i <= max; i++ {
		ints = append(ints, i)
	}

	return ints
}
