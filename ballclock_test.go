package main

import (
	"fmt"
	"testing"
)

func TestBCFFull(t *testing.T) {
	q := BallRung{balls: []Ball{Ball(1)}, cap: 11}

	if q.Full() {
		t.Fatalf("rung is not full")
	}

	for i := 0; i < 10; i++ {
		q.balls = append(q.balls, Ball(0))
	}
	fmt.Println(q.Len())

	if q.Full() == false {
		t.Fatalf("rung is full")
	}
}
