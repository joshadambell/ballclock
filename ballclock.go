package main

import (
	"flag"
	"fmt"

	"github.com/joshadambell/ballclock/clock"
)

func main() {
	balls := flag.Int("balls", 30, "number of balls to use in clock")
	min := flag.Int("min", 0, "number of mintes to run")

	flag.Parse()

	if *min == 0 {
		fmt.Printf("days: %d\n", numberOfDaysUntilCycle(*balls))
		return
	}

	stateAfterMins(*balls, *min)
}

func numberOfDaysUntilCycle(balls int) int {
	// create new clock with number of balls
	bc := clock.NewBallClock(balls)

	for {
		// remove first ball from queue
		b, err := bc.Main.Pop()
		if err != nil {
			fmt.Println(err)
			break
		}

		// update ball clock with remove ball
		bc.Update(b)

		// Test for cycle
		if bc.IsCycle() {
			break
		}
	}

	return convertMinToDays(bc.MinsPassed())
}

func stateAfterMins(balls, min int) {
	// create new clock with number of balls
	bc := clock.NewBallClock(balls)

	for {
		// remove first ball from queue
		b, err := bc.Main.Pop()
		if err != nil {
			fmt.Println(err)
			break
		}

		// update ball clock with remove ball
		bc.Update(b)

		// Test for cycle
		if bc.MinsPassed() == min {
			break
		}
	}

	bc.Print()
}

func convertMinToDays(min int) int {
	return min / 1440
}
