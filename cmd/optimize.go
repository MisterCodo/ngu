package cmd

import (
	"fmt"

	"github.com/MisterCodo/ngu/maps"
	"github.com/spf13/cobra"
)

var mapsMapping = map[int]string{
	1: "Tutorial Island",
	2: "Flesh World",
	3: "Planet Tronne",
	4: "Candy Land",
	5: "Mansions & Managers",
}

var goalMapping = map[int]maps.OptimizationGoal{
	1: maps.SpeedAndProductionGoal,
	2: maps.SpeedGoal,
	3: maps.ProductionGoal,
}

var beaconTypesMapping = map[int][]string{
	1: {"Box"},
	2: {"Box", "Knight"},
	3: {"Box", "Knight", "Arrow"},
	4: {"Box", "Knight", "Arrow", "Wall"},
	5: {"Box", "Knight", "Arrow", "Wall", "Donut"},
}

func init() {
	optimizeCmd.Flags().IntVarP(&mapToOptimize, "map", "m", 1, "map to optimize: (1)Tutorial Island, (2)Flesh World, (3)Planet Tronne, (4)Candy Land, (5)Mansions & Managers")
	optimizeCmd.Flags().IntVarP(&optimizationGoal, "goal", "g", 1, "optimization goal: (1)Speed&Production, (2)Speed, (3)Production")
	optimizeCmd.Flags().IntVarP(&optimizationBeacons, "beacons", "b", 5, "optimization beacon types available: (1)Box, (2)Box & Knight, (3)Box, Knight & Arrow, (4) Box, Knight, Arrow & Wall, (5)All")
	optimizeCmd.Flags().IntVarP(&optimizationSpread, "spread", "s", 3, "optimization modifies up to X tiles at once during randomised hill climbing where X is the spread")
	optimizeCmd.Flags().IntVarP(&mapGoodCount, "cycle", "c", 100, "how many global optimization cycles to run")
	optimizeCmd.Flags().IntVarP(&mapRandomCount, "random", "r", 100, "how many random map to generate per cycle")
	optimizeCmd.Flags().IntVarP(&mapAdjustCount, "adjust", "a", 10000, "how many adjustments to perform on each random map")
	rootCmd.AddCommand(optimizeCmd)
}

var mapToOptimize, optimizationGoal, optimizationBeacons, optimizationSpread, mapGoodCount, mapRandomCount, mapAdjustCount int
var optimizeCmd = &cobra.Command{
	Use:   "optimize",
	Short: "Optimize map beacons.",
	Long:  `Optimize placement of beacons on NGU Industries map.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		mapMaskName, ok := mapsMapping[mapToOptimize]
		if !ok {
			return fmt.Errorf("provided map number invalid")
		}

		goal, ok := goalMapping[optimizationGoal]
		if !ok {
			return fmt.Errorf("provided optimization goal number invalid")
		}

		beaconTypes, ok := beaconTypesMapping[optimizationBeacons]
		if !ok {
			return fmt.Errorf("provided optimization beacons number invalid")
		}

		fmt.Printf("Running %s optimization of map %s with %d cycles, %d random map per cycle and %d adjustments for each random map\n\n", goal.String(), mapMaskName, mapGoodCount, mapRandomCount, mapAdjustCount)
		optimizer, err := maps.NewOptimizer(goal, beaconTypes, mapMaskName, optimizationSpread, mapGoodCount, mapRandomCount, mapAdjustCount)
		if err != nil {
			return fmt.Errorf("could not start optimization: %s", err.Error())
		}
		optimizer.Run()

		return nil
	},
}
