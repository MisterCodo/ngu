package cmd

import (
	"fmt"

	"github.com/MisterCodo/ngu/maps"
	"github.com/spf13/cobra"
)

var mapsIntToName = map[int]string{
	1: "Tutorial Island",
	2: "Flesh World",
	3: "Planet Tronne",
	4: "Candy Land",
	5: "Mansions & Managers",
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
		maps.Optimize(mapMaskName, optimizationTypeName, optimizationSpread, mapGoodCount, mapRandomCount, mapAdjustCount)
		return nil
	},
}
