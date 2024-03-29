package maps

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"sync"
	"time"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

// Optimizer performs map optimization with randomised hill climbing and beam search.
type Optimizer struct {
	Goal           OptimizationGoal // Either Speed, Production or Production&Speed
	LocationName   string           // Location (e.g. tutorial island)
	BeaconTypes    []beacons.BType
	TileRandomizer *TileRandomizer // Allows switching tiles randomly
	BlockedTiles   []int

	AdjustCycle int // How many optimization cycles during randomised hill climbing

	beamMapPool sync.Pool
}

// OptimizationGoal represents the optimization goal.
type OptimizationGoal int

const (
	SpeedAndProductionGoal = iota
	SpeedGoal
	ProductionGoal
)

func (og OptimizationGoal) String() string {
	return [...]string{"SpeedAndProduction", "Speed", "Production"}[og]
}

// NewOptimizer returns a map optimizer for a specific map location, specific goal and using a list of available beacons.
func NewOptimizer(goal OptimizationGoal, beaconTypes []beacons.BType, locationName string, blockedTiles []int) (*Optimizer, error) {
	var beaconCategories []beacons.Category
	if goal == SpeedAndProductionGoal {
		beaconCategories = []beacons.Category{beacons.Speed, beacons.Production}
	} else if goal == SpeedGoal {
		beaconCategories = []beacons.Category{beacons.Speed}
	} else if goal == ProductionGoal {
		beaconCategories = []beacons.Category{beacons.Production}
	}

	// Adjust number of cycles based on number of beacons. Numbers below are based on some quick stats from running the tool.
	var adjustCycle int
	switch len(beaconTypes) {
	case 1:
		adjustCycle = 5000 * len(beaconCategories)
	case 2:
		adjustCycle = 15000 * len(beaconCategories)
	case 3:
		adjustCycle = 40000 * len(beaconCategories)
	case 4:
		adjustCycle = 50000 * len(beaconCategories)
	case 5:
		adjustCycle = 60000 * len(beaconCategories)
	default:
		adjustCycle = 125000
	}

	o := &Optimizer{
		Goal:           goal,
		LocationName:   locationName,
		BeaconTypes:    beaconTypes,
		TileRandomizer: NewTileRandomizer(beaconCategories, beaconTypes),
		AdjustCycle:    adjustCycle,
		BlockedTiles:   blockedTiles,
		beamMapPool: sync.Pool{
			New: func() interface{} {
				return NewMap(locationName, blockedTiles)
			},
		},
	}

	return o, nil
}

// Optimize attempts to find the best map possible for a specific optimization type, be it speed, production or a combination of speed and production.
func (o *Optimizer) Run(drawMap bool, howLong time.Duration) (*Map, error) {
	// Initialize empty map
	bestMap := NewMap(o.LocationName, o.BlockedTiles)
	bestMap.UpdateScore(o.Goal)

	start := time.Now()

	for {
		// Find a very good map
		m := o.generateGoodMapCandidate()

		// Print results
		if m.Score > bestMap.Score {
			fmt.Printf("=====\nNew high score of %.2f\n", m.Score)
			bestMap = m
			bestMap.Print()

			// Generate map image
			if drawMap {
				err := bestMap.Draw(o.Goal, o.BeaconTypes)
				if err != nil {
					return nil, err
				}
			}
		}

		if time.Since(start) > howLong {
			break
		}
	}

	return bestMap, nil
}

// generateGoodMapCandidate finds a good map candidate by first generating a random map, hill climbing and then beam searching.
func (o *Optimizer) generateGoodMapCandidate() *Map {
	m := o.generateGoodRandomMap()

	m = o.hillClimbMap(m)

	// beam search is slow, fix it then reactivate it
	// m = o.beamOptimize(m, 2, 3)

	return m
}

// generateGoodRandomMap generates a random map.
func (o *Optimizer) generateGoodRandomMap() *Map {
	m := NewMap(o.LocationName, o.BlockedTiles)
	m.Randomize(o.TileRandomizer)
	// Todo: Randomize does not update score, fix this. For now just call UpdateScore(.
	m.UpdateScore(o.Goal)

	return m
}

// hillClimbMap performs adjustments on provided map and slowly makes it better.
func (o *Optimizer) hillClimbMap(m *Map) *Map {
	for i := 0; i < o.AdjustCycle; i++ {
		impactedX, impactedY, _, newType, impactedTiles, impactedScore := m.Adjust(o.TileRandomizer, o.Goal)

		if impactedScore > 0.0 {
			// Apply adjust change
			m.Tiles[impactedY][impactedX].Type = newType
			for _, impactedTile := range impactedTiles {
				m.Tiles[impactedTile.Y][impactedTile.X].SpeedMultiplier = impactedTile.NewSpeedMultiplier
				m.Tiles[impactedTile.Y][impactedTile.X].ProductionMultiplier = impactedTile.NewProductionMultiplier
			}

			// Update map score
			m.Score += impactedScore

			// Debugging code block
			// fmt.Printf("new calculation %.2f", m.Score)
			// m.UpdateScore(o.Goal)
			// fmt.Printf(" updated %.2f oldType %s newType %s impactedX %d impactedY %d impactedScore %.2f\n", m.Score, oldType, newType, impactedX, impactedY, impactedScore)
		}
	}
	return m
}

// beamOptimize beam searches the map until it can't improve the map further.
func (o *Optimizer) beamOptimize(m *Map, beamSize int, beamKeep int) *Map {
	for {
		// m.UpdateScore(o.Goal)
		maps := []*Map{m}
		for b := 0; b < beamSize; b++ {
			maps = o.beam(maps, beamKeep)
		}
		sort.Sort(Maps(maps))

		// Check if the beam search offers a better map. If so, use it and do another beam search. Else beam optimize is done.
		if maps[0].Score > m.Score {
			m = maps[0]
			continue
		}
		break
	}
	return m
}

// TODO: this needs work but I'm mostly certain it works properly
func (o *Optimizer) applyBeamImpact(m *Map, x int, y int, beacon string) {
	scoreImpact := 0.0
	oldBeaconType := m.Tiles[y][x].Type

	// If beacon we are replacing was a production tile, then removing the production tile lowers the map score
	if oldBeaconType == ProductionTile {
		if o.Goal == SpeedGoal {
			scoreImpact -= 1.0 + m.Tiles[y][x].SpeedMultiplier/100
		} else if o.Goal == ProductionGoal {
			scoreImpact -= 1.0 + m.Tiles[y][x].ProductionMultiplier/100
		} else if o.Goal == SpeedAndProductionGoal {
			scoreImpact -= (1.0 + m.Tiles[y][x].SpeedMultiplier/100) * (1.0 + m.Tiles[y][x].ProductionMultiplier/100)
		}
	} else { // If beacon we are replacing was a beacon, then we need to remove it's beacon effects and all production tiles affected lowers the map score
		b, ok := beacons.Beacons[oldBeaconType]
		if !ok {
			log.Fatalf("apply beam impact func could not find beacon %s", oldBeaconType)
		}
		for _, effect := range b.Effect() {
			impactedX := x + effect.X
			impactedY := y + effect.Y
			if impactedX >= 0 && impactedX < MapX && impactedY >= 0 && impactedY < MapY {
				oldSpeedProdScore := (1.0 + m.Tiles[impactedY][impactedX].SpeedMultiplier/100) * (1.0 + m.Tiles[impactedY][impactedX].ProductionMultiplier/100)
				if b.Category() == beacons.Speed {
					m.Tiles[impactedY][impactedX].SpeedMultiplier -= effect.Gain
				} else if b.Category() == beacons.Production {
					m.Tiles[impactedY][impactedX].ProductionMultiplier -= effect.Gain
				}
				if m.Tiles[impactedY][impactedX].Type == ProductionTile {
					if o.Goal == SpeedGoal || o.Goal == ProductionGoal {
						scoreImpact -= effect.Gain / 100
					} else if o.Goal == SpeedAndProductionGoal {
						newSpeedProdScore := (1.0 + m.Tiles[impactedY][impactedX].SpeedMultiplier/100) * (1.0 + m.Tiles[impactedY][impactedX].ProductionMultiplier/100)
						scoreImpact += newSpeedProdScore - oldSpeedProdScore
					}
				}
			}
		}
	}

	// Update map beacon
	m.Tiles[y][x].Type = beacon

	// If new beacon is a production tile, it increases the map score
	if beacon == ProductionTile {
		if o.Goal == SpeedGoal {
			scoreImpact += 1.0 + m.Tiles[y][x].SpeedMultiplier/100
		} else if o.Goal == ProductionGoal {
			scoreImpact += 1.0 + m.Tiles[y][x].ProductionMultiplier/100
		} else if o.Goal == SpeedAndProductionGoal {
			scoreImpact += (1.0 + m.Tiles[y][x].SpeedMultiplier/100) * (1.0 + m.Tiles[y][x].ProductionMultiplier/100)
		}
	} else { // If new beacon is a beacon, then apply it's effects to other tiles and if those are production tiles the map score increases
		b, ok := beacons.Beacons[beacon]
		if !ok {
			log.Fatalf("apply beam impact func could not find beacon %s", beacon)
		}
		for _, effect := range b.Effect() {
			impactedX := x + effect.X
			impactedY := y + effect.Y
			if impactedX >= 0 && impactedX < MapX && impactedY >= 0 && impactedY < MapY {
				oldSpeedProdScore := (1.0 + m.Tiles[impactedY][impactedX].SpeedMultiplier/100) * (1.0 + m.Tiles[impactedY][impactedX].ProductionMultiplier/100)
				if b.Category() == beacons.Speed {
					m.Tiles[impactedY][impactedX].SpeedMultiplier += effect.Gain
				} else if b.Category() == beacons.Production {
					m.Tiles[impactedY][impactedX].ProductionMultiplier += effect.Gain
				}
				if m.Tiles[impactedY][impactedX].Type == ProductionTile {
					if o.Goal == SpeedGoal || o.Goal == ProductionGoal {
						scoreImpact += effect.Gain / 100
					} else if o.Goal == SpeedAndProductionGoal {
						newSpeedProdScore := (1.0 + m.Tiles[impactedY][impactedX].SpeedMultiplier/100) * (1.0 + m.Tiles[impactedY][impactedX].ProductionMultiplier/100)
						scoreImpact += newSpeedProdScore - oldSpeedProdScore
					}
				}
			}
		}
	}

	// Update map score
	m.Score += scoreImpact
}

// beam performs a beam search on the map
func (o *Optimizer) beam(maps Maps, beamKeep int) Maps {
	returnMaps := []*Map{}
	for _, m := range maps {
		// Generate all possible maps 1 change away from map m
		tmpMaps := []*Map{}
		for y, row := range m.Tiles {
			for x := range row {
				if m.Tiles[y][x].Type != UnusableTile {
					for _, bt := range o.TileRandomizer.Beacons {
						if m.Tiles[y][x].Type == bt {
							continue
						}
						tmpMap := m.CopyUsing(o.beamMapPool.Get().(*Map))
						o.applyBeamImpact(tmpMap, x, y, bt)
						// this commented block helps validate the applyBeamImpact func accuracy
						// asdf := tmpMap.Score
						// tmpMap.UpdateScore(o.Goal)
						// fmt.Printf("%.2f %.2f\n", asdf, tmpMap.Score)
						tmpMaps = append(tmpMaps, tmpMap)
					}
				}
			}
		}

		// Shuffle the maps so that two maps with the same score have an equal chance of being in the top scoring maps
		rand.Shuffle(len(tmpMaps), func(i, j int) { tmpMaps[i], tmpMaps[j] = tmpMaps[j], tmpMaps[i] })

		// Keep the best X maps only
		sort.Sort(Maps(tmpMaps))
		howMany := beamKeep
		if len(tmpMaps) < howMany {
			howMany = len(tmpMaps)
		}

		// return unused maps to the pool
		for _, tmpM := range tmpMaps[howMany:] {
			o.beamMapPool.Put(tmpM)
		}

		returnMaps = append(returnMaps, tmpMaps[0:howMany-1]...)
	}
	return returnMaps
}
