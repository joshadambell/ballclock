package clock

import "encoding/json"

type BallRung struct {
	balls []Ball
	cap   int
}

func NewBallRung(cap int) *BallRung {
	return &BallRung{balls: make([]Ball, 0), cap: cap}
}

func (br *BallRung) MarshalJSON() ([]byte, error) {
	return json.Marshal(br.balls)
}

func (br BallRung) Full() bool {
	return br.Len() == br.cap
}

func (br BallRung) Len() int {
	return len(br.balls)
}

// Push - attempts to add a ball to the rung. If rung becomes full, return balls emptied from rung
func (br *BallRung) Push(b Ball) []Ball {
	emptiedBalls := make([]Ball, 0)

	if br.Full() {
		// remove balls in reverse from rung
		for i := br.Len() - 1; i >= 0; i-- {
			emptiedBalls = append(emptiedBalls, br.balls[i])
		}

		br.balls = make([]Ball, 0)
		return emptiedBalls
	}

	br.balls = append(br.balls, b)

	return emptiedBalls

}
