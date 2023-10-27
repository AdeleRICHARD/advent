package day2

type TypeOfPlay string

const (
	ROCK     TypeOfPlay = "Rock"
	PAPER    TypeOfPlay = "Paper"
	SCISSORS TypeOfPlay = "Scissors"
)

func GetType(s string) TypeOfPlay {
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

func LoosingType(typeOfPlay TypeOfPlay) TypeOfPlay {
	switch typeOfPlay {
	case ROCK:
		return SCISSORS
	case PAPER:
		return ROCK
	case SCISSORS:
		return PAPER
	default:
		return ""
	}
}

func WinningType(typeToPlay TypeOfPlay) TypeOfPlay {
	switch typeToPlay {
	case ROCK:
		return PAPER
	case PAPER:
		return SCISSORS
	case SCISSORS:
		return ROCK
	default:
		return ""
	}
}

func GetMoves(player1 TypeOfPlay, strategy string) TypeOfPlay {
	switch strategy {
	case "X":
		return LoosingType(player1)
	case "Z":
		return WinningType(player1)
	case "Y":
		return player1
	default:
		return ""
	}
}
