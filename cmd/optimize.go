package cmd

import (
	"fmt"
	"os"

	"github.com/MisterCodo/ngu/maps"
	"github.com/spf13/cobra"
)

var mapsIntToName = map[int]string{
	1: "TutorialIsland",
	2: "FleshWorld",
	3: "PlanetTronne",
	4: "CandyLand",
	5: "MansionsAndManagers",
}

var optimizationIntToName = map[int]string{
	1: "SpeedAndProduction",
	2: "Speed",
	3: "Production",
}

func init() {
	optimizeCmd.Flags().IntVarP(&mapToOptimize, "map", "m", 1, "map to optimize: (1)Tutorial Island, (2)Flesh World, (3)Planet Tronne, (4)Candy Land, (5)Mansions & Managers")
	optimizeCmd.Flags().IntVarP(&optimizationType, "type", "t", 1, "optimization type: (1)Speed&Production, (2)Speed, (3)Production")
	optimizeCmd.Flags().IntVarP(&optimizationSpread, "spread", "s", 3, "optimization modifies up to X tiles at once where X is the spread")
	optimizeCmd.Flags().IntVarP(&mapGoodCount, "cycle", "c", 100, "how many global optimization cycles to run")
	optimizeCmd.Flags().IntVarP(&mapRandomCount, "random", "r", 100, "how many random map to generate per cycle")
	optimizeCmd.Flags().IntVarP(&mapAdjustCount, "adjust", "a", 50000, "how many adjustments to perform on each random map")
	rootCmd.AddCommand(optimizeCmd)
}

var mapToOptimize, optimizationType, optimizationSpread, mapGoodCount, mapRandomCount, mapAdjustCount int
var optimizeCmd = &cobra.Command{
	Use:   "optimize",
	Short: "Optimize map beacons.",
	Long:  `Optimize placement of beacons on NGU Industries map.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		mapMaskName, ok := mapsIntToName[mapToOptimize]
		if !ok {
			return fmt.Errorf("provided map number invalid")
		}

		optimizationTypeName, ok := optimizationIntToName[optimizationType]
		if !ok {
			return fmt.Errorf("provided optimization type number invalid")
		}

		fmt.Printf("Running %s optimization of map %s with %d cycles, %d random map per cycle and %d adjustments for each random map\n\n", optimizationTypeName, mapMaskName, mapGoodCount, mapRandomCount, mapAdjustCount)
		optimize(mapMaskName, optimizationTypeName, optimizationSpread)
		return nil
	},
}

func optimize(mapMaskName string, optimizationType string, optimizationSpread int) {
	mask, ok := maps.MapMasks[mapMaskName]
	if !ok {
		fmt.Printf("could not find map mask %s\n", mapMaskName)
		os.Exit(-1)
	}

	bestMap := findBestMap(mask, optimizationType, optimizationSpread)
	bestMap.Print()
	fmt.Printf("\nFinal map for %s scored: %.2f\n", mapMaskName, bestMap.Score(optimizationType))
}

func findBestMap(mask maps.MapMask, optimizationType string, optimizationSpread int) *maps.Map {
	var bestMap *maps.Map
	highScore := -1.0

	// Find a very good map
	for i := 0; i < mapGoodCount; i++ {
		m := findGoodMap(mask, optimizationType, optimizationSpread)
		newScore := m.Score(optimizationType)
		fmt.Printf("Cycle i:%d map scored:%.2f\n", i, newScore)
		if newScore > highScore {
			bestMap = m
			highScore = newScore
		}
	}

	// Optimize this best candidate to get the best map
	fmt.Println("")
	fmt.Println("Found one best candidate, performing final optimization")
	bestMap = optimizeMap(bestMap, optimizationType, optimizationSpread)

	fmt.Println("")
	return bestMap
}

func findGoodMap(mask maps.MapMask, optimizationType string, optimizationSpread int) *maps.Map {
	// Generate random map
	bestMap := findRandomMap(mask, optimizationType)

	// Optimize map
	for i := 1; i < optimizationSpread+1; i++ {
		bestMap = optimizeMap(bestMap, optimizationType, i)
	}

	return bestMap
}

func optimizeMap(bestMap *maps.Map, optimizationType string, numChangedTiles int) *maps.Map {
	highScore := bestMap.Score(optimizationType)
	for i := 0; i < mapAdjustCount; i++ {
		m := bestMap.Copy()
		for j := 0; j < numChangedTiles; j++ {
			m.Adjust(optimizationType)
		}
		newScore := m.Score(optimizationType)
		if newScore > highScore {
			// if numChangedTiles != 1 {
			// 	fmt.Printf("new:%f old:%f numChangedTiles:%d\n", newScore, highScore, numChangedTiles)
			// }
			bestMap = m
			highScore = newScore
			// repeat cycle if a change was made
			if numChangedTiles != 1 {
				bestMap = optimizeMap(bestMap, optimizationType, numChangedTiles-1)
				i = 0
			}
		}
	}
	return bestMap
}

func findRandomMap(mask maps.MapMask, optimizationType string) *maps.Map {
	highScore := 0.0
	bestMap := maps.NewMap(mask)
	for i := 0; i < mapRandomCount; i++ {
		m := maps.NewMap(mask)
		m.Randomize(optimizationType)
		newScore := m.Score(optimizationType)
		if newScore > highScore {
			bestMap = m
			highScore = newScore
		}
	}
	return bestMap
}
