package main

import (
	"log"
	"math"
)

func main() {
	earliestDeparture := 1000655
	bs := []int{17, 37, 571, 13, 23, 29, 401, 41, 19}

	var nextBus int
	nd := math.MaxInt64
	for _, b := range bs {
		n := b * int(math.Ceil(float64(earliestDeparture)/float64(b)))
		if n < nd {
			nd = n
			nextBus = b
		}
	}

	log.Printf("Next available departure on bus %d at %d (hash %d)", nextBus, nd, (nd-earliestDeparture)*nextBus)
}
