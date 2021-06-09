package maps

import (
	"math/rand"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

// Tile consists of either a regular resource tile, a unusable tile or a beacon tile.
type Tile struct {
	Type                 string
	ProductionMultiplier float64
	SpeedMultiplier      float64
	EfficiencyMultiplier float64
}

// TileRandomizer provides a random tile based on available beacons
type TileRandomizer struct {
	Beacons []string // Available beacons (symbols only)
}

func NewTileRandomizer(category string, btypes []string) *TileRandomizer {
	// Find available beacons based on category and beacon types

	// This sucks, fix it
	var categories []string
	if category == "SpeedAndProduction" {
		categories = []string{"Production", "Speed"}
	} else {
		categories = []string{category}
	}

	beaconsAvailable := []string{}
	for beaconSymbol, beacon := range beacons.Beacons {
		for _, c := range categories {
			if beacon().Category().String() == c {
				for _, btype := range btypes {
					if btype == beacon().BType().String() {
						beaconsAvailable = append(beaconsAvailable, beaconSymbol)
						continue
					}
				}
			}
		}
	}

	return &TileRandomizer{Beacons: beaconsAvailable}
}

// randomTile returns either an available beacon (symbol) or a regular production tile symbol.
func (tr *TileRandomizer) randomTile() string {
	r := rand.Intn(len(tr.Beacons) + 2)
	if r < len(tr.Beacons) {
		return tr.Beacons[r]
	}
	return "."
}
