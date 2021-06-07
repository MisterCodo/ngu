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
	optimizeCmd.Flags().IntVarP(&mapToOptimize, "map", "m", 1, "map to optimize. (1)Tutorial Island, (2)Flesh World, (3)Planet Tronne, (4)Candy Land, (5)Mansions & Managers")
	optimizeCmd.Flags().IntVarP(&optimizationType, "type", "t", 1, "optimization type. (1)Speed&Production, (2)Speed, (3)Production")
	optimizeCmd.Flags().IntVarP(&mapGoodCount, "cycle", "c", 100, "how many global optimization cycles to run")
	optimizeCmd.Flags().IntVarP(&mapRandomCount, "random", "r", 100, "how many random map to generate per cycle")
	optimizeCmd.Flags().IntVarP(&mapAdjustCount, "adjust", "a", 50000, "how many adjustments to perform on each random map")
	rootCmd.AddCommand(optimizeCmd)
}

var mapToOptimize, optimizationType, mapGoodCount, mapRandomCount, mapAdjustCount int
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
		optimize(mapMaskName, optimizationTypeName)
		return nil
	},
}

func optimize(mapMaskName string, optimizationType string) {
	mask, ok := maps.MapMasks[mapMaskName]
	if !ok {
		fmt.Printf("could not find map mask %s\n", mapMaskName)
		os.Exit(-1)
	}

	bestMap := findBestMap(mask, optimizationType)
	bestMap.Print()
	fmt.Printf("\nFinal map for %s scored: %.2f\n", mapMaskName, bestMap.Score(optimizationType))
}

func findBestMap(mask maps.MapMask, optimizationType string) *maps.Map {
	var bestMap *maps.Map
	highScore := -1.0

	for i := 0; i < mapGoodCount; i++ {
		m := findGoodMap(mask, optimizationType)
		newScore := m.Score(optimizationType)
		fmt.Printf("Cycle i:%d map scored:%.2f\n", i, newScore)
		if newScore > highScore {
			bestMap = m
			highScore = newScore
		}
	}
	fmt.Println("")

	return bestMap
}

func findGoodMap(mask maps.MapMask, optimizationType string) *maps.Map {
	baseMap := maps.NewMap(mask)
	highScore := baseMap.Score(optimizationType)

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

	for i := 0; i < mapAdjustCount; i++ {
		m := bestMap.Copy()
		m.Adjust(optimizationType)
		newScore := m.Score(optimizationType)
		if newScore > highScore {
			bestMap = m
			highScore = newScore
		}
	}

	return bestMap
}
