package maps

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/MisterCodo/ngu/plugins/beacons"
	"github.com/MisterCodo/ngu/plugins/locations"
)

// Optimizer performs map optimization with randomised hill climbing and beam search.
type Optimizer struct {
	Goal           OptimizationGoal   // Either Speed, Production or Production&Speed
	Location       locations.Location // Location (e.g. tutorial island)
	TileRandomizer *TileRandomizer    // Allows switching tiles randomly

	CandidatesCount int // How many map candidates to generate during optimization
	RandomMapCount  int // How many random map to generate for each candidate map
	AdjustCycle     int // How many optimization cycles during randomised hill climbing
	Spread          int // How much spread during randomised hill climbing
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
func NewOptimizer(goal OptimizationGoal, beaconTypes []beacons.BType, locationName string, spread int, candidateCount int, randomMapCount int, adjustCycle int) (*Optimizer, error) {
	location, ok := locations.Locations[locationName]
	if !ok {
		return nil, fmt.Errorf("could not find map location %s", locationName)
	}

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
		Location:        location(),
		TileRandomizer:  NewTileRandomizer(beaconCategories, beaconTypes),
		CandidatesCount: candidateCount,
		RandomMapCount:  randomMapCount,
		AdjustCycle:     adjustCycle,
		Spread:          spread,
	}

	return o, nil
}

// Optimize attempts to find the best map possible for a specific optimization type, be it speed, production or a combination of speed and production.
func (o *Optimizer) Run() error {
	var bestMap *Map
	highScore := -1.0

	// Find a very good map
	for i := 0; i < o.CandidatesCount; i++ {
		m := o.generateGoodMapCandidate()
		newScore := m.Score(o.Goal)
		m.Print()
		fmt.Printf("Cycle i:%d map scored:%.2f\n\n", i, newScore)
		if newScore > highScore {
			bestMap = m
			highScore = newScore
		}
	}

	// Optimize this best candidate to get the best map
	fmt.Println("Found one best candidate, performing final optimization")
	bestMap = o.beamOptimize(bestMap, 3, 5)

	// Print results
	score := bestMap.Score(o.Goal)
	fmt.Printf("\n%s (%.2f)\n", o.Location.PrettyName(), score)
	bestMap.Print()
	fmt.Println("")

	// Generate map image
	err := bestMap.Draw()
	if err != nil {
		return err
	}

	return nil
}

// generateGoodMapCandidate finds a good map candidate by first generating a random map, hill climbing and then beam searching.
func (o *Optimizer) generateGoodMapCandidate() *Map {
	m := o.generateGoodRandomMap()

	for i := 1; i < o.Spread+1; i++ {
		m = o.hillClimbMap(m, i)
	}

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
func (o *Optimizer) hillClimbMap(bestMap *Map, numChangedTiles int) *Map {
	highScore := bestMap.Score(o.Goal)
	for i := 0; i < o.AdjustCycle; i++ {
		m := bestMap.Copy()
		for j := 0; j < numChangedTiles; j++ {
			m.Adjust(o.TileRandomizer)
		}
		newScore := m.Score(o.Goal)
		if newScore > highScore {
			bestMap = m
			highScore = newScore
			// repeat cycle if a change was made
			if numChangedTiles != 1 {
				bestMap = o.hillClimbMap(bestMap, numChangedTiles-1)
				i = 0
			}
		}
	}
	return bestMap
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
