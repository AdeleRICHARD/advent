package day2

type TypeOfPlay string

const (
	ROCK     TypeOfPlay = "Rock"
	PAPER    TypeOfPlay = "Paper"
	SCISSORS TypeOfPlay = "Scissors"
)

func getType(s string) TypeOfPlay {
	switch s {
	case "A":
		return ROCK
	case "B":
		return PAPER
	case "C":
		return SCISSORS
	case "X":
		return ROCK
	case "Y":
		return PAPER
	case "Z":
		return SCISSORS
	default:
		return ""
	}
}

func (t TypeOfPlay) String() string {
	return string(t)
}

func (t TypeOfPlay) Number() int {
	switch t {
	case ROCK:
		return 1
	case PAPER:
		return 2
	case SCISSORS:
		return 3
	default:
		return 0
	}
}
