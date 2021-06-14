package maps

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

// Optimizer performs map optimization with randomised hill climbing and beam search.
type Optimizer struct {
	Goal           OptimizationGoal // Either Speed, Production or Production&Speed
	Location       string           // Location (e.g. tutorial island)
	BeaconTypes    []beacons.BType
	TileRandomizer *TileRandomizer // Allows switching tiles randomly

	CandidatesCount int // How many map candidates to generate during optimization
	RandomMapCount  int // How many random map to generate for each candidate map
	AdjustCycle     int // How many optimization cycles during randomised hill climbing
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
func NewOptimizer(goal OptimizationGoal, beaconTypes []beacons.BType, location string, candidateCount int, randomMapCount int, adjustCycle int) (*Optimizer, error) {
	var beaconCategories []beacons.Category
	if goal == SpeedAndProductionGoal {
		beaconCategories = []beacons.Category{beacons.Speed, beacons.Production}
	} else if goal == SpeedGoal {
		beaconCategories = []beacons.Category{beacons.Speed}
	} else if goal == ProductionGoal {
		beaconCategories = []beacons.Category{beacons.Production}
	}

	o := &Optimizer{
		Goal:            goal,
		Location:        location,
		BeaconTypes:     beaconTypes,
		TileRandomizer:  NewTileRandomizer(beaconCategories, beaconTypes),
		CandidatesCount: candidateCount,
		RandomMapCount:  randomMapCount,
		AdjustCycle:     adjustCycle,
	}

	return o, nil
}

// Optimize attempts to find the best map possible for a specific optimization type, be it speed, production or a combination of speed and production.
func (o *Optimizer) Run(drawMap bool) (*Map, error) {
	var bestMap *Map
	highScore := -1.0

	infinite := false
	if o.CandidatesCount == -1 {
		infinite = true
		o.CandidatesCount = 3
	}

	bestScore := 0.0

	for {
		// Find a very good map
		for i := 0; i < o.CandidatesCount; i++ {
			m := o.generateGoodMapCandidate()
			newScore := m.Score(o.Goal)
			if newScore > highScore {
				bestMap = m
				highScore = newScore
			}
		}

		// Optimize this best candidate to get the best map
		bestMap = o.beamOptimize(bestMap, 3, 5)

		// Print results
		score := bestMap.Score(o.Goal)
		if score > bestScore {
			fmt.Printf("=====\nNew high score of %.2f\n", score)
			bestScore = score
			bestMap.Print()

			// Generate map image
			if drawMap {
				err := bestMap.Draw(o.Goal, o.BeaconTypes)
				if err != nil {
					return nil, err
				}
			}
		}

		if infinite {
			continue
		}
		break
	}

	return bestMap, nil
}

// generateGoodMapCandidate finds a good map candidate by first generating a random map, hill climbing and then beam searching.
func (o *Optimizer) generateGoodMapCandidate() *Map {
	m := o.generateGoodRandomMap()

	m = o.hillClimbMap(m)

	m = o.beamOptimize(m, 2, 3)

	return m
}

// generateGoodRandomMap tries to find a good random map.
func (o *Optimizer) generateGoodRandomMap() *Map {
	highScore := 0.0
	bestMap := NewMap(o.Location)
	for i := 0; i < o.RandomMapCount; i++ {
		m := NewMap(o.Location)
		m.Randomize(o.TileRandomizer)
		newScore := m.Score(o.Goal)
		if newScore > highScore {
			bestMap = m
			highScore = newScore
		}
	}
	return bestMap
}

// hillClimbMap performs adjustments on provided map and slowly makes it better.
func (o *Optimizer) hillClimbMap(m *Map) *Map {
	highScore := m.Score(o.Goal)
	for i := 0; i < o.AdjustCycle; i++ {
		impactedX, impactedY, oldType := m.Adjust(o.TileRandomizer)
		newScore := m.Score(o.Goal)
		if newScore > highScore {
			highScore = newScore
		} else {
			//reset move
			m.Tiles[impactedY][impactedX].Type = oldType
		}
	}
	return m
}

// beamOptimize beam searches the map until it can't improve the map further.
func (o *Optimizer) beamOptimize(m *Map, beamSize int, beamKeep int) *Map {
	for {
		highScore := m.Score(o.Goal)

		maps := []*Map{m}
		for b := 0; b < beamSize; b++ {
			maps = o.beam(maps, beamKeep)
		}
		sort.Sort(BySpeedScore{maps})

		if maps[0].Score(o.Goal) > highScore {
			m = maps[0]
			continue
		}
		break
	}
	return m
}

// beam performs a beam search on the map
func (o *Optimizer) beam(maps Maps, beamKeep int) Maps {
	returnMaps := []*Map{}
	for _, m := range maps {
		// Generate all possible maps 1 change away from map m
		tmpMaps := []*Map{}
		for y, row := range m.Tiles {
			for x := range row {
				if m.Tiles[y][x].Type == ProductionTile {
					for _, bt := range o.TileRandomizer.Beacons {
						if m.Tiles[y][x].Type == bt {
							continue
						}
						tmpMap := m.Copy()
						tmpMap.Tiles[y][x].Type = bt
						tmpMaps = append(tmpMaps, tmpMap)
					}
				}
			}
		}

		// Shuffle the maps so that two maps with the same score have an equal chance of being in the top scoring maps
		rand.Shuffle(len(tmpMaps), func(i, j int) { tmpMaps[i], tmpMaps[j] = tmpMaps[j], tmpMaps[i] })

		// Keep the best X maps only
		if o.Goal == SpeedGoal {
			sort.Sort(BySpeedScore{tmpMaps})
		} else if o.Goal == ProductionGoal {
			sort.Sort(ByProductionScore{tmpMaps})
		} else {
			sort.Sort(BySpeedAndProductionScore{tmpMaps})
		}
		howMany := beamKeep
		if len(tmpMaps) < howMany {
			howMany = len(tmpMaps)
		}

		returnMaps = append(returnMaps, tmpMaps[0:howMany-1]...)
	}
	return returnMaps
}
