package maps

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
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
	Location string
	Score    float64
}

// NewMap generates a new map based on the provided location. Each tile which can be
// modified by the gamer is set to a regular resource tile.
func NewMap(location string) *Map {
	m := &Map{
		Tiles:    [MapY][MapX]Tile{},
		Location: location,
	}
	l, ok := locations.Locations[location]
	if !ok {
		log.Fatalf("NewMap could not find location %s", location)
	}
	for y, row := range l.Mask() {
		for x, val := range row {
			if val == 1 {
				m.Tiles[y][x] = Tile{Type: ProductionTile, ProductionMultiplier: 0.0, SpeedMultiplier: 0.0, EfficiencyMultiplier: 0.0}
				m.Score++
				continue
			}
			m.Tiles[y][x] = Tile{Type: UnusableTile, ProductionMultiplier: 0.0, SpeedMultiplier: 0.0, EfficiencyMultiplier: 0.0}
		}
	}

	return m
}

// Randomize picks random tile types for the entire map. Careful, it does not update the map score.
func (m *Map) Randomize(t *TileRandomizer) {
	for y, row := range m.Tiles {
		for x := range row {
			if m.Tiles[y][x].Type != UnusableTile {
				m.Tiles[y][x].Type = t.randomTile()
			}
		}
	}
}

// Copy creates a new map with the same tiles as the original map.
func (m *Map) Copy() *Map {
	newMap := &Map{
		Location: m.Location,
		Tiles:    [MapY][MapX]Tile{},
		Score:    m.Score,
	}
	for y, row := range m.Tiles {
		for x := range row {
			newMap.Tiles[y][x] = m.Tiles[y][x]
		}
	}
	return newMap
}

// CopyUsing creates a new map with the same tiles as the original map using an old map reference.
func (m *Map) CopyUsing(oldMap *Map) *Map {
	oldMap.Location = m.Location
	oldMap.Score = m.Score
	for y, row := range m.Tiles {
		for x := range row {
			oldMap.Tiles[y][x] = m.Tiles[y][x]
		}
	}
	return oldMap
}

// Score evaluates multipliers for each tile of the map and evaluates the score of the map.
func (m *Map) UpdateScore(goal OptimizationGoal) {
	// Reset Multipliers
	for y, row := range m.Tiles {
		for x := range row {
			m.Tiles[y][x].SpeedMultiplier = 0.0
			m.Tiles[y][x].ProductionMultiplier = 0.0
			m.Tiles[y][x].EfficiencyMultiplier = 0.0
		}
	}

	// Go through all map tiles and apply the effect of each beacon found.
	for y, row := range m.Tiles {
		for x, val := range row {
			// Skip production and unusable tiles
			if val.Type == ProductionTile || val.Type == UnusableTile {
				continue
			}

			// Speed beacons
			b, ok := beacons.Beacons[val.Type]
			if !ok {
				log.Fatalf("score func could not find beacon %s", val.Type)
			}
			if b.Category() == beacons.Speed {
				effects := b.Effect()
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
			if b.Category() == beacons.Production {
				effects := b.Effect()
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

	// Update map score
	if goal == SpeedGoal {
		m.Score = speedScore
	} else if goal == ProductionGoal {
		m.Score = productionScore
	} else if goal == SpeedAndProductionGoal {
		m.Score = productionAndSpeedScore
	} else {
		m.Score = 0.0
	}
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
	l, ok := locations.Locations[m.Location]
	if !ok {
		log.Fatalf("NewMap could not find location %s", m.Location)
	}
	img := l.Image()
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
			b, ok := beacons.Beacons[beaconType]
			if !ok {
				log.Fatalf("draw func could not find beacon %s", beaconType)
			}
			beaconImg := b.Image()
			sr = beaconImg.Bounds()
			r := sr.Sub(sr.Min).Add(image.Point{x * beacons.ImgSize, y * beacons.ImgSize})
			draw.Draw(outputImg, r, beaconImg, image.Point{}, draw.Over)
		}
	}

	// Save image to disk
	m.UpdateScore(goal)
	var outName string
	if len(beaconTypes) == 0 || goal == -1 {
		outName = strings.Join([]string{l.UglyName(), fmt.Sprintf("%d", time.Now().Unix())}, "_") + ".png"
	} else {
		outName = strings.Join([]string{l.UglyName(), OptimizationGoal(goal).String(), beaconTypes[len(beaconTypes)-1].String(), fmt.Sprintf("%.0f", m.Score*100)}, "_") + ".png"
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

// Adjust changes one tile of the map to another type. It sends details about which tile got modified and
// what was the previous tile type.
func (m *Map) Adjust(tr *TileRandomizer, goal OptimizationGoal) (impactedX int, impactedY int, oldType string, newType string, impactedTiles []impactedTile, impactedScore float64) {
	// Find a tile to adjust, it must be a valid spot.
	for {
		impactedX = rand.Intn(MapX)
		impactedY = rand.Intn(MapY)
		if m.Tiles[impactedY][impactedX].Type != UnusableTile {
			break
		}
	}
	oldType = m.Tiles[impactedY][impactedX].Type

	// Find a new type for the tile, it must be different than the current type.
	for {
		newType = tr.randomTile()
		if oldType != newType {
			break
		}
	}

	// Evaluate tile multipliers and score impact
	impactedTiles, scoreImpact := m.evaluateTileScoreImpact(impactedX, impactedY, oldType, newType, goal)

	return impactedX, impactedY, oldType, newType, impactedTiles, scoreImpact
}

type impactedTile struct {
	X                       int
	Y                       int
	OldSpeedMultiplier      float64
	NewSpeedMultiplier      float64
	OldProductionMultiplier float64
	NewProductionMultiplier float64
	ScoreImpact             float64
}

func (m *Map) evaluateTileScoreImpact(x int, y int, oldType string, newType string, goal OptimizationGoal) (impactedTiles []impactedTile, scoreImpact float64) {
	impactedTiles = []impactedTile{}
	scoreImpact = 0.0

	// If beacon we are replacing was a production tile, then removing the production tile lowers the map score
	if oldType == ProductionTile {
		if goal == SpeedGoal {
			scoreImpact -= 1.0 + m.Tiles[y][x].SpeedMultiplier/100
		} else if goal == ProductionGoal {
			scoreImpact -= 1.0 + m.Tiles[y][x].ProductionMultiplier/100
		} else { // SpeedAndProductionGoal
			scoreImpact -= (1.0 + m.Tiles[y][x].SpeedMultiplier/100) * (1.0 + m.Tiles[y][x].ProductionMultiplier/100)
		}
	} else {
		// If beacon we are replacing was a beacon, then we need to remove it's beacon effects and all production tiles affected lowers the map score
		b, ok := beacons.Beacons[oldType]
		if !ok {
			log.Fatalf("apply beam impact func could not find beacon %s", oldType)
		}
		for _, effect := range b.Effect() {
			impactedX := x + effect.X
			impactedY := y + effect.Y
			if impactedX >= 0 && impactedX < MapX && impactedY >= 0 && impactedY < MapY {
				impactedTile := impactedTile{
					X:                       impactedX,
					Y:                       impactedY,
					OldSpeedMultiplier:      m.Tiles[impactedY][impactedX].SpeedMultiplier,
					OldProductionMultiplier: m.Tiles[impactedY][impactedX].ProductionMultiplier,
					ScoreImpact:             0.0,
				}

				// measure the old combined speed and production score
				oldSpeedProdScore := (1.0 + impactedTile.OldSpeedMultiplier/100) * (1.0 + impactedTile.OldProductionMultiplier/100)

				// new beacon impacts either speed or production
				if b.Category() == beacons.Speed {
					impactedTile.NewSpeedMultiplier = impactedTile.OldSpeedMultiplier - effect.Gain
				} else if b.Category() == beacons.Production {
					impactedTile.NewProductionMultiplier = impactedTile.OldProductionMultiplier - effect.Gain
				}

				// If impacted tile was a production tile, then score needs to be adjusted
				if m.Tiles[impactedY][impactedX].Type == ProductionTile {
					if goal == SpeedGoal || goal == ProductionGoal {
						// fmt.Printf("negative impactedY %d impactedX %d effect %.2f\n", impactedY, impactedX, effect.Gain)
						impactedTile.ScoreImpact -= effect.Gain / 100
					} else if goal == SpeedAndProductionGoal {
						newSpeedProdScore := (1.0 + impactedTile.NewSpeedMultiplier/100) * (1.0 + impactedTile.NewProductionMultiplier/100)
						impactedTile.ScoreImpact += newSpeedProdScore - oldSpeedProdScore
					}
				}

				impactedTiles = append(impactedTiles, impactedTile)
			}
		}
	}

	// If new type is a production tile, it increases the map score
	if newType == ProductionTile {
		if goal == SpeedGoal {
			scoreImpact += 1.0 + m.Tiles[y][x].SpeedMultiplier/100
			// fmt.Printf("added %.2f\n", scoreImpact)
		} else if goal == ProductionGoal {
			scoreImpact += 1.0 + m.Tiles[y][x].ProductionMultiplier/100
		} else { // SpeedAndProductionGoal
			scoreImpact += (1.0 + m.Tiles[y][x].SpeedMultiplier/100) * (1.0 + m.Tiles[y][x].ProductionMultiplier/100)
		}
	} else {
		// If new beacon is a beacon, then apply it's effects to other tiles and if those are production tiles the map score increases
		b, ok := beacons.Beacons[newType]
		if !ok {
			log.Fatalf("apply beam impact func could not find beacon %s", newType)
		}
		for _, effect := range b.Effect() {
			impactedX := x + effect.X
			impactedY := y + effect.Y
			if impactedX >= 0 && impactedX < MapX && impactedY >= 0 && impactedY < MapY {
				impactedTile := impactedTile{
					X:                       impactedX,
					Y:                       impactedY,
					OldSpeedMultiplier:      m.Tiles[impactedY][impactedX].SpeedMultiplier,
					OldProductionMultiplier: m.Tiles[impactedY][impactedX].ProductionMultiplier,
					ScoreImpact:             0.0,
				}

				// check if this tile was already impacted by removal
				doubleImpacted := false
				itID := 0
				for id, it := range impactedTiles {
					if it.X == impactedX && it.Y == impactedY {
						impactedTile.OldSpeedMultiplier = it.NewSpeedMultiplier
						impactedTile.OldProductionMultiplier = it.NewProductionMultiplier
						impactedTile.ScoreImpact = it.ScoreImpact
						doubleImpacted = true
						itID = id
					}
				}

				// measure the old combined speed and production score
				oldSpeedProdScore := (1.0 + impactedTile.OldSpeedMultiplier/100) * (1.0 + impactedTile.OldProductionMultiplier/100)

				// new beacon impacts either speed or production
				if b.Category() == beacons.Speed {
					impactedTile.NewSpeedMultiplier = impactedTile.OldSpeedMultiplier + effect.Gain
				} else if b.Category() == beacons.Production {
					impactedTile.NewProductionMultiplier = impactedTile.OldProductionMultiplier + effect.Gain
				}

				// If impacted tile was a production tile, then score needs to be adjusted
				if m.Tiles[impactedY][impactedX].Type == ProductionTile {
					if goal == SpeedGoal || goal == ProductionGoal {
						impactedTile.ScoreImpact += effect.Gain / 100
					} else if goal == SpeedAndProductionGoal {
						newSpeedProdScore := (1.0 + impactedTile.NewSpeedMultiplier/100) * (1.0 + impactedTile.NewProductionMultiplier/100)
						impactedTile.ScoreImpact += newSpeedProdScore - oldSpeedProdScore
					}
				}

				// Update impacted tiles
				if doubleImpacted {
					impactedTiles[itID].NewSpeedMultiplier = impactedTile.NewSpeedMultiplier
					impactedTiles[itID].NewProductionMultiplier = impactedTile.NewProductionMultiplier
					impactedTiles[itID].ScoreImpact = impactedTile.ScoreImpact
				} else {
					impactedTiles = append(impactedTiles, impactedTile)
				}
			}
		}
	}

	for _, it := range impactedTiles {
		scoreImpact += it.ScoreImpact
	}

	return impactedTiles, scoreImpact
}

type Maps []*Map

func (m Maps) Len() int           { return len(m) }
func (m Maps) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m Maps) Less(i, j int) bool { return m[i].Score > m[j].Score }
