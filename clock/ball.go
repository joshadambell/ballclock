package clock

type Ball int

func (b Ball) Value() int {
	return int(b)
}