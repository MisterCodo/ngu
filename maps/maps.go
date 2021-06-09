package maps

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/MisterCodo/ngu/plugins/beacons"
	_ "github.com/MisterCodo/ngu/plugins/beacons/all"
)

const (
	MapX = 20
	MapY = 17
)

func init() {
	// change rand seed so that each run is different
	rand.Seed(time.Now().UnixNano())
}

// Map is a NGU Industries map layout consisting of tiles and a map mask.
type Map struct {
	Tiles [MapY][MapX]Tile
	Mask  MapMask
}

// NewMap generates a new map based on the provided mask. Each tile which can be
// modified by the gamer is set to a regular resource tile.
func NewMap(mask MapMask) *Map {
	m := &Map{
		Tiles: [MapY][MapX]Tile{},
		Mask:  mask,
	}
	for y, row := range m.Mask {
		for x, val := range row {
			if val == 1 {
				m.Tiles[y][x] = Tile{Type: ".", ProductionMultiplier: 0.0, SpeedMultiplier: 0.0, EfficiencyMultiplier: 0.0}
				continue
			}
			m.Tiles[y][x] = Tile{Type: " ", ProductionMultiplier: 0.0, SpeedMultiplier: 0.0, EfficiencyMultiplier: 0.0}
		}
	}

	return m
}

// Randomize picks random tile types for the entire map.
func (m *Map) Randomize(t *TileRandomizer) {
	for y, row := range m.Tiles {
		for x := range row {
			if m.Mask[y][x] == 1 {
				m.Tiles[y][x].Type = t.randomTile()
			}
		}
	}
}

// Copy creates a new map with the same tiles as the original map.
func (m *Map) Copy() *Map {
	newMap := NewMap(m.Mask)
	for y, row := range m.Tiles {
		for x := range row {
			newMap.Tiles[y][x] = m.Tiles[y][x]
		}
	}
	return newMap
}

// Score evaluates the score of the map.
func (m *Map) Score(optimizationType string) float64 {
	// Go through all map tiles and apply the effect of each beacon found.
	for y, row := range m.Tiles {
		for x, val := range row {
			// Speed beacons
			if val.Type == "*" || val.Type == "<" || val.Type == ">" || val.Type == "v" || val.Type == "^" || val.Type == "k" || val.Type == "-" || val.Type == "|" || val.Type == "o" {
				effects := beacons.Beacons[val.Type]().Effect()
				for _, effect := range effects {
					impactedX := x + effect.X
					impactedY := y + effect.Y
					if impactedX >= 0 && impactedX < MapX && impactedY >= 0 && impactedY < MapY {
						m.Tiles[impactedY][impactedX].SpeedMultiplier += effect.Gain
					}
				}
			}

			// Production beacons
			if val.Type == "b" || val.Type == "l" || val.Type == "r" || val.Type == "d" || val.Type == "u" || val.Type == "&" || val.Type == "h" || val.Type == "w" || val.Type == "O" {
				effects := beacons.Beacons[val.Type]().Effect()
				for _, effect := range effects {
					impactedX := x + effect.X
					impactedY := y + effect.Y
					if impactedX >= 0 && impactedX < MapX && impactedY >= 0 && impactedY < MapY {
						m.Tiles[impactedY][impactedX].ProductionMultiplier += effect.Gain
					}
				}
			}
		}
	}

	// Measure score
	speedScore := 0.0
	productionScore := 0.0
	productionAndSpeedScore := 0.0
	for _, row := range m.Tiles {
		for _, val := range row {
			if val.Type == "." {
				tSpeedScore := (100.0 + val.SpeedMultiplier) / 100
				speedScore += tSpeedScore
				tProductionScore := (100.0 + val.ProductionMultiplier) / 100
				productionScore += tProductionScore
				productionAndSpeedScore += tSpeedScore * tProductionScore
			}
		}
	}

	// Reset Multipliers, this is crappy
	for y, row := range m.Tiles {
		for x := range row {
			m.Tiles[y][x].SpeedMultiplier = 0.0
			m.Tiles[y][x].ProductionMultiplier = 0.0
			m.Tiles[y][x].EfficiencyMultiplier = 0.0
		}
	}

	if optimizationType == "Speed" {
		return speedScore
	}
	if optimizationType == "Production" {
		return productionScore
	}
	return productionAndSpeedScore
}

// Print displays the map layout.
func (m *Map) Print() {
	for _, row := range m.Tiles {
		for _, val := range row {
			fmt.Printf("%s", val.Type)
		}
		fmt.Println("")
	}
}

// Adjust changes one tile of the map to another type.
func (m *Map) Adjust(tr *TileRandomizer) {
	// Find a tile to adjust, it must be a valid spot based on the map mask.
	var impactedX, impactedY int
	for {
		impactedX = rand.Intn(MapX)
		impactedY = rand.Intn(MapY)
		if m.Mask[impactedY][impactedX] == 1 {
			break
		}
	}

	// Find a new type for the tile, it must be different than the current type.
	var newType string
	for {
		newType = tr.randomTile()
		if m.Tiles[impactedY][impactedX].Type != newType {
			break
		}
	}

	// Apply change
	m.Tiles[impactedY][impactedX].Type = newType
}

type Maps []*Map

func (m Maps) Len() int      { return len(m) }
func (m Maps) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

type BySpeedScore struct{ Maps }

func (s BySpeedScore) Less(i, j int) bool {
	return s.Maps[i].Score("Speed") > s.Maps[j].Score("Speed")
}

type ByProductionScore struct{ Maps }

func (s ByProductionScore) Less(i, j int) bool {
	return s.Maps[i].Score("Production") > s.Maps[j].Score("Production")
}

type BySpeedAndProductionScore struct{ Maps }

func (s BySpeedAndProductionScore) Less(i, j int) bool {
	return s.Maps[i].Score("SpeedAndProduction") > s.Maps[j].Score("SpeedAndProduction")
}
