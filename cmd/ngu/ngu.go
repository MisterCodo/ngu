package main

import (
	"fmt"

	"github.com/MisterCodo/ngu"
)

func main() {
	// New map
	mask := ngu.TutorialIslandMask
	baseMap := ngu.NewMap(mask)

	// m.Print()
	highScore := baseMap.Score()
	fmt.Printf("Base score: %.2f\n", highScore)

	bestMap := ngu.NewMap(mask)
	for i := 0; i < 100000; i++ {
		m := ngu.NewMap(mask)
		m.Randomize()
		newScore := m.Score()
		if newScore > highScore {
			bestMap = m
			highScore = newScore
			fmt.Printf("Better randomized map i:%d score: %.2f\n", i, newScore)
		}
	}
	bestMap.Print()
	fmt.Printf("Best randomized map score: %.2f\n", bestMap.Score())

	for i := 0; i < 300000; i++ {
		m := bestMap.Copy()
		m.Adjust()
		newScore := m.Score()
		if newScore > highScore {
			bestMap = m
			highScore = newScore
			fmt.Printf("Better adjusted map i:%d score: %.2f\n", i, newScore)
		}
	}

	bestMap.Print()
	fmt.Printf("Final map of score: %.2f\n", bestMap.Score())
}
