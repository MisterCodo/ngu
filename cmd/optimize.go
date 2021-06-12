package cmd

import (
	"fmt"

	"github.com/MisterCodo/ngu/maps"
	"github.com/MisterCodo/ngu/plugins/beacons"
	"github.com/spf13/cobra"
)

var locationsMapping = map[int]string{
	1: "TutorialIsland",
	2: "FleshWorld",
	3: "PlanetTronne",
	4: "CandyLand",
	5: "MansionsAndManagers",
}

var goalMapping = map[int]maps.OptimizationGoal{
	1: maps.SpeedAndProductionGoal,
	2: maps.SpeedGoal,
	3: maps.ProductionGoal,
}

var beaconTypesMapping = map[int][]beacons.BType{
	1: {beacons.Box},
	2: {beacons.Box, beacons.Knight},
	3: {beacons.Box, beacons.Knight, beacons.Arrow},
	4: {beacons.Box, beacons.Knight, beacons.Arrow, beacons.Wall},
	5: {beacons.Box, beacons.Knight, beacons.Arrow, beacons.Wall, beacons.Donut},
}

func init() {
	optimizeCmd.Flags().IntVarP(&optimizeCmdMap, "map", "m", 1, "map to optimize: (1)Tutorial Island, (2)Flesh World, (3)Planet Tronne, (4)Candy Land, (5)Mansions & Managers")
	optimizeCmd.Flags().IntVarP(&optimizeCmdGoal, "goal", "g", 1, "optimization goal: (1)Speed&Production, (2)Speed, (3)Production")
	optimizeCmd.Flags().IntVarP(&optimizeCmdBeacons, "beacons", "b", 5, "optimization beacon types available: (1)Box, (2)Box & Knight, (3)Box, Knight & Arrow, (4) Box, Knight, Arrow & Wall, (5)All")
	optimizeCmd.Flags().IntVarP(&optimizeCmdSpread, "spread", "s", 3, "optimization modifies up to X tiles at once during randomised hill climbing where X is the spread")
	optimizeCmd.Flags().IntVarP(&optimizeCmdCycles, "cycle", "c", 100, "how many global optimization cycles to run. Set to -1 to run forever.")
	optimizeCmd.Flags().IntVarP(&optimizeCmdRandomCycles, "random", "r", 100, "how many random map to generate per cycle")
	optimizeCmd.Flags().IntVarP(&optimizeCmdAdjustCycles, "adjust", "a", 10000, "how many adjustments to perform on each random map")
	rootCmd.AddCommand(optimizeCmd)
}

var optimizeCmdMap, optimizeCmdGoal, optimizeCmdBeacons, optimizeCmdSpread, optimizeCmdCycles, optimizeCmdRandomCycles, optimizeCmdAdjustCycles int
var optimizeCmd = &cobra.Command{
	Use:   "optimize",
	Short: "Optimize map beacons.",
	Long:  `Optimize placement of beacons on NGU Industries map.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		locationName, ok := locationsMapping[optimizeCmdMap]
		if !ok {
			return fmt.Errorf("provided map number invalid")
		}

		goal, ok := goalMapping[optimizeCmdGoal]
		if !ok {
			return fmt.Errorf("provided optimization goal number invalid")
		}

		beaconTypes, ok := beaconTypesMapping[optimizeCmdBeacons]
		if !ok {
			return fmt.Errorf("provided optimization beacons number invalid")
		}

		optimizer, err := maps.NewOptimizer(goal, beaconTypes, locationName, optimizeCmdSpread, optimizeCmdCycles, optimizeCmdRandomCycles, optimizeCmdAdjustCycles)
		if err != nil {
			return fmt.Errorf("could not start optimization: %s", err.Error())
		}

		fmt.Printf("Running %s optimization of map %s with %d cycles\n\n", goal.String(), optimizer.Location.PrettyName(), optimizeCmdCycles)

		err = optimizer.Run()
		if err != nil {
			return fmt.Errorf("could not run optimization: %s", err.Error())
		}

		return nil
	},
}
