package maps

import (
	"fmt"
	"math/rand"
	"os"
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

// Tile consists of either a regular resource tile, a unusable tile or a beacon tile.
type Tile struct {
	Type                 string
	ProductionMultiplier float64
	SpeedMultiplier      float64
	EfficiencyMultiplier float64
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
func (m *Map) Randomize(optimizationType string) {
	for y, row := range m.Tiles {
		for x := range row {
			if m.Mask[y][x] == 1 {
				m.Tiles[y][x].Type = randTileType(optimizationType)
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
func (m *Map) Adjust(optimizationType string) {
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
		newType = randTileType(optimizationType)
		if m.Tiles[impactedY][impactedX].Type != newType {
			break
		}
	}

	// Apply change
	m.Tiles[impactedY][impactedX].Type = newType
}

// randTileType returns a random tile type based on which beacons are available.
func randTileType(optimizationType string) string {
	if optimizationType == "Speed" {
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
		case r == 8:
			return "o"
		default:
			return "."
		}
	}

	if optimizationType == "Production" {
		r := rand.Intn(10)
		switch {
		case r == 0:
			return "b"
		case r == 1:
			return "l"
		case r == 2:
			return "r"
		case r == 3:
			return "d"
		case r == 4:
			return "u"
		case r == 5:
			return "&"
		case r == 6:
			return "h"
		case r == 7:
			return "w"
		case r == 8:
			return "O"
		default:
			return "."
		}
	}

	// Production&Speed
	r := rand.Intn(20)
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
	case r == 8:
		return "o"
	case r == 10:
		return "b"
	case r == 11:
		return "l"
	case r == 12:
		return "r"
	case r == 13:
		return "d"
	case r == 14:
		return "u"
	case r == 15:
		return "&"
	case r == 16:
		return "h"
	case r == 17:
		return "w"
	case r == 18:
		return "O"
	default:
		return "."
	}
}

// Optimize attempts to find the best map possible for a specific optimization type, be it speed, production or a combination of speed and production.
func Optimize(mapMaskName string, optimizationType string, optimizationSpread int, mapGoodCount int, mapRandomCount int, mapAdjustCount int) {
	mask, ok := MapMasks[mapMaskName]
	if !ok {
		fmt.Printf("could not find map mask %s\n", mapMaskName)
		os.Exit(-1)
	}

	bestMap := findBestMap(mask, optimizationType, optimizationSpread, mapGoodCount, mapRandomCount, mapAdjustCount)
	bestMap.Print()
	fmt.Printf("\nScore: %.2f\n", bestMap.Score(optimizationType))
}

// findBestMap returns the best map from all good map candidates generated.
func findBestMap(mask MapMask, optimizationType string, optimizationSpread int, mapGoodCount int, mapRandomCount int, mapAdjustCount int) *Map {
	var bestMap *Map
	highScore := -1.0

	// Find a very good map
	for i := 0; i < mapGoodCount; i++ {
		m := findGoodMap(mask, optimizationType, optimizationSpread, mapRandomCount, mapAdjustCount)
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
	bestMap = optimizeMap(bestMap, optimizationType, optimizationSpread, mapAdjustCount)

	fmt.Println("")
	return bestMap
}

// findGoodMap finds a good map candidate.
func findGoodMap(mask MapMask, optimizationType string, optimizationSpread int, mapRandomCount int, mapAdjustCount int) *Map {
	// Generate random map
	bestMap := generateGoodRandomMap(mask, optimizationType, mapRandomCount)

	// Optimize map
	for i := 1; i < optimizationSpread+1; i++ {
		bestMap = optimizeMap(bestMap, optimizationType, i, mapAdjustCount)
	}

	return bestMap
}

// generateGoodRandomMap tries to find a good starting random map.
func generateGoodRandomMap(mask MapMask, optimizationType string, mapRandomCount int) *Map {
	highScore := 0.0
	bestMap := NewMap(mask)
	for i := 0; i < mapRandomCount; i++ {
		m := NewMap(mask)
		m.Randomize(optimizationType)
		newScore := m.Score(optimizationType)
		if newScore > highScore {
			bestMap = m
			highScore = newScore
		}
	}
	return bestMap
}

// optimizeMap performs adjustments on provided map and slowly makes it better.
func optimizeMap(bestMap *Map, optimizationType string, numChangedTiles int, mapAdjustCount int) *Map {
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
				bestMap = optimizeMap(bestMap, optimizationType, numChangedTiles-1, mapAdjustCount)
				i = 0
			}
		}
	}
	return bestMap
}
