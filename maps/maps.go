package maps

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/MisterCodo/ngu/plugins/beacons"
	_ "github.com/MisterCodo/ngu/plugins/beacons/all"
	"github.com/MisterCodo/ngu/plugins/locations"
	_ "github.com/MisterCodo/ngu/plugins/locations/all"
)

const (
	MapX = 20 // How many X tiles in a map
	MapY = 17 // How many Y tiles in a map
)

func init() {
	// change rand seed so that each run is different
	rand.Seed(time.Now().UnixNano())
}

// Map is a NGU Industries map layout consisting of tiles and a location.
type Map struct {
	Tiles    [MapY][MapX]Tile
	Location locations.Location
}

// NewMap generates a new map based on the provided location. Each tile which can be
// modified by the gamer is set to a regular resource tile.
func NewMap(location locations.Location) *Map {
	m := &Map{
		Tiles:    [MapY][MapX]Tile{},
		Location: location,
	}
	for y, row := range m.Location.Mask() {
		for x, val := range row {
			if val == 1 {
				m.Tiles[y][x] = Tile{Type: ProductionTile, ProductionMultiplier: 0.0, SpeedMultiplier: 0.0, EfficiencyMultiplier: 0.0}
				continue
			}
			m.Tiles[y][x] = Tile{Type: UnusableTile, ProductionMultiplier: 0.0, SpeedMultiplier: 0.0, EfficiencyMultiplier: 0.0}
		}
	}

	return m
}

// Randomize picks random tile types for the entire map.
func (m *Map) Randomize(t *TileRandomizer) {
	for y, row := range m.Tiles {
		for x := range row {
			if m.Location.Mask()[y][x] == 1 {
				m.Tiles[y][x].Type = t.randomTile()
			}
		}
	}
}

// Copy creates a new map with the same tiles as the original map.
func (m *Map) Copy() *Map {
	newMap := NewMap(m.Location)
	for y, row := range m.Tiles {
		for x := range row {
			newMap.Tiles[y][x] = m.Tiles[y][x]
		}
	}
	return newMap
}

// Score evaluates the score of the map.
func (m *Map) Score(goal OptimizationGoal) float64 {
	// Go through all map tiles and apply the effect of each beacon found.
	for y, row := range m.Tiles {
		for x, val := range row {
			// Skip production and unusable tiles
			if val.Type == ProductionTile || val.Type == UnusableTile {
				continue
			}

			// Speed beacons
			b := beacons.Beacons[val.Type]
			if b().Category() == beacons.Speed {
				effects := beacons.Beacons[val.Type]().Effect()
				for _, effect := range effects {
					impactedX := x + effect.X
					impactedY := y + effect.Y
					if impactedX >= 0 && impactedX < MapX && impactedY >= 0 && impactedY < MapY {
						m.Tiles[impactedY][impactedX].SpeedMultiplier += effect.Gain
					}
				}
				continue
			}

			// Production beacons
			if b().Category() == beacons.Production {
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
			if val.Type == ProductionTile {
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

	if goal == SpeedGoal {
		return speedScore
	}
	if goal == ProductionGoal {
		return productionScore
	}
	if goal == SpeedAndProductionGoal {
		return productionAndSpeedScore
	}
	return 0.0
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

// DrawMap draws the map image.
func (m *Map) Draw(goal OptimizationGoal, beaconTypes []beacons.BType) error {
	// Initialize output image
	img := m.Location.Image()
	outputImg := image.NewRGBA(image.Rect(0, 0, locations.ImgSizeX, locations.ImgSizeY))
	sr := img.Bounds()
	draw.Draw(outputImg, sr, img, image.Point{}, draw.Src)

	// Go through each tile and if it's a beacon, print on top of the loaded image
	for y, row := range m.Tiles {
		for x := range row {
			beaconType := m.Tiles[y][x].Type
			if beaconType == UnusableTile || beaconType == ProductionTile {
				continue
			}
			beaconImg := beacons.Beacons[beaconType]().Image()
			sr = beaconImg.Bounds()
			r := sr.Sub(sr.Min).Add(image.Point{x * beacons.ImgSize, y * beacons.ImgSize})
			draw.Draw(outputImg, r, beaconImg, image.Point{}, draw.Over)
		}
	}

	// Save image to disk
	score := m.Score(goal)
	var outName string
	if len(beaconTypes) == 0 || goal == -1 {
		outName = strings.Join([]string{m.Location.UglyName(), fmt.Sprintf("%d", time.Now().Unix())}, "_") + ".png"
	} else {

		outName = strings.Join([]string{m.Location.UglyName(), OptimizationGoal(goal).String(), beaconTypes[len(beaconTypes)-1].String(), fmt.Sprintf("%.0f", score*100)}, "_") + ".png"
	}

	out, err := os.Create(outName)
	if err != nil {
		return err
	}
	defer out.Close()

	err = png.Encode(out, outputImg)
	if err != nil {
		return err
	}

	fmt.Printf("Generated output image %s\n", outName)
	return nil
}

// Adjust changes one tile of the map to another type.
func (m *Map) Adjust(tr *TileRandomizer) {
	// Find a tile to adjust, it must be a valid spot based on the map mask.
	var impactedX, impactedY int
	for {
		impactedX = rand.Intn(MapX)
		impactedY = rand.Intn(MapY)
		if m.Location.Mask()[impactedY][impactedX] == 1 {
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
	return s.Maps[i].Score(SpeedGoal) > s.Maps[j].Score(SpeedGoal)
}

type ByProductionScore struct{ Maps }

func (s ByProductionScore) Less(i, j int) bool {
	return s.Maps[i].Score(ProductionGoal) > s.Maps[j].Score(ProductionGoal)
}

type BySpeedAndProductionScore struct{ Maps }

func (s BySpeedAndProductionScore) Less(i, j int) bool {
	return s.Maps[i].Score(SpeedAndProductionGoal) > s.Maps[j].Score(SpeedAndProductionGoal)
}
