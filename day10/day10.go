package day10

import (
	"strconv"
	"strings"

	"main.go/util"
)

// Keep track of cycles
// addx [nb] = 2 complete cycles -> So at nb 3 add the nb to X
// noop do nothing
// Sum all the numbers after each end of the cycles
// Multiply the result by the total nb of cycles we've done
// Need to check only at 20th, 60th, 100th, 140th and 180th cycle

func Day10() {
	datas := util.ReadFile("day10/datas.txt")
	cpuCycles := CPU{TotalCycle: 0, X: 0}
	totalSum := 0
	for _, data := range datas {
		instructions := strings.Fields(string(data))
		if instructions[0] == "noop" {
			cpuCycles.TotalCycle++
			continue
		}
		if instructions[0] == "addx" {
			cpuCycles.TotalCycle++
			if isCycleToMultiply(&cpuCycles) {
				totalSum += cpuCycles.TotalCycle * cpuCycles.X
			}
			cpuCycles.TotalCycle++
			x, _ := strconv.Atoi(instructions[1])
			cpuCycles.AddX(x)
		}
	}
}

type CPU struct {
	TotalCycle int
	X          int
}

func (c *CPU) AddX(nb int) {
	c.X += nb
}

func isCycleToMultiply(c *CPU) bool {
	return c.TotalCycle == 20 || c.TotalCycle == 60 || c.TotalCycle == 100 || c.TotalCycle == 140 || c.TotalCycle == 180
}
