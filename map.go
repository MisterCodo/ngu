package ngu

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
	rand.Seed(time.Now().UnixNano())
}

type Map struct {
	Tiles [MapY][MapX]Tile
	Mask  MapMask
}

type Tile struct {
	Type                 string
	ProductionMultiplier float64
	SpeedMultiplier      float64
	EfficiencyMultiplier float64
}

type MapMask [MapY][MapX]int

var MapMasks = map[string]MapMask{
	"TutorialIsland": {
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 0, 1, 0, 1, 1, 0},
		{0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0},
		{0, 0, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0},
		{0, 0, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
}

func NewMap(mask MapMask) *Map {
	m := &Map{
		Tiles: [MapY][MapX]Tile{},
		Mask:  mask,
	}

	// Default to all normal production tiles based on map mask.
	for y, row := range m.Mask {
		for x, val := range row {
			if val == 1 {
				m.Tiles[y][x] = Tile{Type: ".", ProductionMultiplier: 0.0, SpeedMultiplier: 0.0, EfficiencyMultiplier: 0.0}
				continue
			}
			m.Tiles[y][x] = Tile{Type: "x", ProductionMultiplier: 0.0, SpeedMultiplier: 0.0, EfficiencyMultiplier: 0.0}
		}
	}

	return m
}

func (m *Map) Score() float64 {
	// Apply beacons. Go through all map tiles and apply the effect of each beacon found.
	for y, row := range m.Tiles {
		for x, val := range row {
			// TODO: Fix this once not just speed beacons are used.
			if val.Type == "*" || val.Type == "<" || val.Type == ">" || val.Type == "v" || val.Type == "^" || val.Type == "k" || val.Type == "-" || val.Type == "|" {
				effects := beacons.Beacons[val.Type]().Effect()
				for _, effect := range effects {
					impactedX := x + effect.X
					impactedY := y + effect.Y
					if impactedX >= 0 && impactedX < MapX && impactedY >= 0 && impactedY < MapY {
						m.Tiles[impactedY][impactedX].SpeedMultiplier += effect.Gain
					}
				}
			}
		}
	}

	// Measure score
	speedScore := 0.0
	for _, row := range m.Tiles {
		for _, val := range row {
			if val.Type == "." {
				speedScore += 1.0 * ((100.0 + val.SpeedMultiplier) / 100)
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

	return speedScore
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
func (m *Map) Adjust() {
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
		newType = randTileType()
		if m.Tiles[impactedY][impactedX].Type != newType {
			break
		}
	}

	// Apply change
	m.Tiles[impactedY][impactedX].Type = newType
}

// Randomize picks random tile types for the entire map.
func (m *Map) Randomize() {
	// For a change on only a subset of tiles at the start
	for y, row := range m.Tiles {
		for x := range row {
			if m.Mask[y][x] == 1 && rand.Intn(5) == 0 {
				m.Tiles[y][x].Type = randTileType()
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

func randTileType() string {
	r := rand.Intn(10)
	switch {
	case r == 0:
		return "*"
	case r == 1:
		return "<"
	case r == 2:
		return ">"
	case r == 3:
		return "v"
	case r == 4:
		return "^"
	case r == 5:
		return "k"
	case r == 6:
		return "-"
	case r == 7:
		return "|"
	default:
		return "."
	}
}
