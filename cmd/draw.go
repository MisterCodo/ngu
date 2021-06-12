package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MisterCodo/ngu/maps"
	"github.com/MisterCodo/ngu/plugins/locations"
	"github.com/spf13/cobra"
)

func init() {
	drawCmd.Flags().IntVarP(&drawCmdMap, "map", "m", 1, "map to optimize: (1)Tutorial Island, (2)Flesh World, (3)Planet Tronne, (4)Candy Land, (5)Mansions & Managers")
	drawCmd.Flags().StringVarP(&drawCmdFile, "file", "f", "", "file consisting of characters defining placement of beacons. These beacons will be drawn on top of selected map.")
	drawCmd.MarkFlagRequired("file")
	rootCmd.AddCommand(drawCmd)
}

var drawCmdMap int
var drawCmdFile string
var drawCmd = &cobra.Command{
	Use:   "draw",
	Short: "Draw map.",
	Long:  `Draw map to disk for choosen location according to beacons file provided.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Load the location
		locationName, ok := locationsMapping[drawCmdMap]
		if !ok {
			return fmt.Errorf("provided map number invalid")
		}
		location, ok := locations.Locations[locationName]
		if !ok {
			return fmt.Errorf("could not find map location %s", locationName)
		}

		// Convert file to a map
		m := maps.NewMap(location())

		f, err := os.Open(drawCmdFile)
		if err != nil {
			return err
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanLines)
		y := 0
		for scanner.Scan() {
			if scanner.Text() == "" {
				continue
			}
			if y >= maps.MapY {
				return fmt.Errorf("height provided in beacons file is larger than map allows")
			}
			for x, char := range strings.Split(scanner.Text(), "") {
				if x >= maps.MapX {
					return fmt.Errorf("width provided in beacons file is larger than map allows")
				}
				m.Tiles[y][x].Type = char
			}
			y += 1
		}
		if err := scanner.Err(); err != nil {
			return fmt.Errorf("scan of file failed: %w", err)
		}

		// Draw it
		m.Draw()

		return nil
	},
}
