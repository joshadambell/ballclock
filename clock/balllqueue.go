package clock

import (
	"encoding/json"
	"fmt"
)

type BallQueue struct {
	items []Ball
	size  int
}

func NewBallQueue(balls int) *BallQueue {
	q := &BallQueue{
		items: make([]Ball, 0),
		size:  balls,
	}

	for i := 1; i < balls+1; i++ {
		q.Push(Ball(i))
	}

	return q
}

func (br *BallQueue) MarshalJSON() ([]byte, error) {
	return json.Marshal(br.items)
}

func (q *BallQueue) Push(n Ball) {
	q.items = append(q.items, n)
}

func (q *BallQueue) BatchPush(balls []Ball) {
	for _, b := range balls {
		q.items = append(q.items, b)
	}
}

func (q *BallQueue) Pop() (Ball, error) {
	if q.Len() < 1 {
		return Ball(0), fmt.Errorf("Queue is empty")
	}

	n := (q.items)[0]
	q.items = (q.items)[1:]
	return n, nil
}

func (q *BallQueue) Peek() (Ball, bool) {
	if q.Len() < 1 {
		return Ball(0), false
	}

	return (q.items)[0], true
}

func (q *BallQueue) Len() int {
	return len(q.items)
}

func (q *BallQueue) Full() bool {
	return q.size == q.Len()
}
