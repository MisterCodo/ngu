package main

import (
	"fmt"
	"os"

	"github.com/MisterCodo/ngu"
)

const (
	mapGoodCount   = 10000
	mapRandomCount = 100
	mapAdjustCount = 50000
)

func main() {
	mapMaskName := "FleshWorld"
	mask, ok := ngu.MapMasks[mapMaskName]
	if !ok {
		fmt.Printf("could not find map mask %s\n", mapMaskName)
		os.Exit(-1)
	}

	bestMap := findBestMap(mask)
	bestMap.Print()
	fmt.Printf("Final map of score: %.2f\n", bestMap.Score())
}

func findBestMap(mask ngu.MapMask) *ngu.Map {
	var bestMap *ngu.Map
	highScore := -1.0

	for i := 0; i < mapGoodCount; i++ {
		m := findGoodMap(mask)
		newScore := m.Score()
		fmt.Printf("Good map i:%d scored:%.2f\n", i, newScore)
		if newScore > highScore {
			bestMap = m
			highScore = newScore
		}
	}

	return bestMap
}

func findGoodMap(mask ngu.MapMask) *ngu.Map {
	baseMap := ngu.NewMap(mask)
	highScore := baseMap.Score()

	bestMap := ngu.NewMap(mask)
	for i := 0; i < mapRandomCount; i++ {
		m := ngu.NewMap(mask)
		m.Randomize()
		newScore := m.Score()
		if newScore > highScore {
			bestMap = m
			highScore = newScore
		}
	}

	for i := 0; i < mapAdjustCount; i++ {
		m := bestMap.Copy()
		m.Adjust()
		newScore := m.Score()
		if newScore > highScore {
			bestMap = m
			highScore = newScore
		}
	}

	return bestMap
}
