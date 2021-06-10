package maps

import (
	"math/rand"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

const (
	ProductionTile = "."
	UnusableTile   = " "
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

// NewTileRandomizer returns a new tile randomizer. It sets available beacons based on category and beacon types.
func NewTileRandomizer(categories []beacons.Category, btypes []beacons.BType) *TileRandomizer {
	beaconsAvailable := []string{}
	for beaconSymbol, beacon := range beacons.Beacons {
		for _, c := range categories {
			if beacon().Category() == c {
				for _, btype := range btypes {
					if btype == beacon().BType() {
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
	r := rand.Intn(len(tr.Beacons) + 2) // +2 is arbitrary, just using a bit more production tiles
	if r < len(tr.Beacons) {
		return tr.Beacons[r]
	}
	return ProductionTile
}
