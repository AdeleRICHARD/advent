package day10

import (
	"strings"

	"github.com/spf13/cast"
	"main.go/util"
)

// Keep track of cycles
// addx [nb] = 2 complete cycles -> So at nb 3 add the nb to X
// noop do nothing
// Sum all the numbers after each end of the cycles
// Multiply the result by the total nb of cycles we've done
// Need to check only at 20th, 60th, 100th, 140th and 180th cycle

type CPU struct {
	IndexInstruction int
	X                int
}

func (c *CPU) AddX(nb int) {
	c.X += nb
}

func Day10() {
	datas := util.ReadFile("day10/data.txt")
	dataSplit := strings.Split(string(datas), "\n")
	cpuCycles := CPU{IndexInstruction: 0, X: 1}
	totalSum := 0
	instructions := parseInput(dataSplit)

	for cycle := 1; cycle <= 220; cycle++ {
		if (cycle-20)%40 == 0 {
			totalSum += cpuCycles.X * cycle
		}
		switch instructions[cpuCycles.IndexInstruction].name {
		case "addx":
			instructions[cpuCycles.IndexInstruction].cycles--
			if instructions[cpuCycles.IndexInstruction].cycles == 0 {
				cpuCycles.AddX(instructions[cpuCycles.IndexInstruction].val)
				cpuCycles.IndexInstruction++
			}
		case "noop":
			cpuCycles.IndexInstruction++
		}
	}

	//part1(string(datas))
	// 13520 Good one
	println(totalSum)
}

type instruction struct {
	name   string
	val    int
	cycles int
}

func parseInput(input []string) (ans []instruction) {
	for _, l := range input {
		switch strings.Split(l, " ")[0] {
		case "addx":
			ans = append(ans, instruction{
				name:   "addx",
				val:    cast.ToInt(l[1]),
				cycles: 2,
			})
		case "noop":
			ans = append(ans, instruction{
				name:   "noop",
				cycles: 1,
			})
		default:
			panic("input line: " + l)
		}
	}
	return ans
}
