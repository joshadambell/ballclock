package clock

import (
	"encoding/json"
	"fmt"
)

type BallClock struct {
	Min       *BallRung
	FiveMin   *BallRung
	Hour      *BallRung
	Main      *BallQueue
	minPassed int
}

func NewBallClock(balls int) *BallClock {
	return &BallClock{
		Min:     NewBallRung(4),
		FiveMin: NewBallRung(11),
		Hour:    NewBallRung(11),
		Main:    NewBallQueue(balls),
	}
}

func (bc BallClock) Print() {
	b, err := json.Marshal(bc)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", b)
}

// Update - alters clock rungs based on passed in ball
func (bc *BallClock) Update(b Ball) {
	// add one min for every ball
	bc.minPassed++

	// one minute rung
	emptiedBalls := bc.Min.Push(b)
	if len(emptiedBalls) == 0 {
		return
	}

	// add emptied balls to pool
	bc.Main.BatchPush(emptiedBalls)

	// five minute rung
	emptiedBalls = bc.FiveMin.Push(b)
	if len(emptiedBalls) == 0 {
		return
	}

	// add emptied balls to pool
	bc.Main.BatchPush(emptiedBalls)

	// five minute rung
	emptiedBalls = bc.Hour.Push(b)
	if len(emptiedBalls) == 0 {
		return
	}

	// add emptied balls to pool
	bc.Main.BatchPush(append(emptiedBalls, b))
}

// IsCycle - tests if main pool is in starting order
func (bc BallClock) IsCycle() bool {
	if bc.Main.Full() == false {
		return false
	}

	tmp := *bc.Main
	for i := 1; i < tmp.size+1; i++ {
		b, _ := tmp.Pop()
		if b.Value() != i {
			return false
		}
	}

	return true
}

func (bc BallClock) MinsPassed() int {
	return bc.minPassed
}
