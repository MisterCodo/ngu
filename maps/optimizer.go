package maps

import (
	"fmt"
	"sort"
)

// Optimizer performs map optimization with randomised hill climbing and beam search.
type Optimizer struct {
	Goal           string          // Either Speed, Production or Production&Speed
	Beacons        []string        // Available beacons for the optmization
	Mask           MapMask         // Map mask
	TileRandomizer *TileRandomizer // Allows switching tiles randomly

	CandidatesCount int // How many map candidates to generate during optimization
	RandomMapCount  int // How many random map to generate for each candidate map
	AdjustCycle     int // How many optimization cycles during randomised hill climbing
	Spread          int // How much spread during randomised hill climbing
}

// NewOptimizer returns a map optimizer for a specific map, specific goal and using a list of available beacons.
func NewOptimizer(goal string, beacons []string, mapMaskName string, spread int, candidateCount int, randomMapCount int, adjustCycle int) (*Optimizer, error) {
	mask, ok := MapMasks[mapMaskName]
	if !ok {
		return nil, fmt.Errorf("could not find map mask %s", mapMaskName)
	}

	o := &Optimizer{
		Goal:            goal,
		Beacons:         beacons,
		Mask:            mask,
		TileRandomizer:  NewTileRandomizer(goal, beacons),
		CandidatesCount: candidateCount,
		RandomMapCount:  randomMapCount,
		AdjustCycle:     adjustCycle,
		Spread:          spread,
	}

	return o, nil
}

// Optimize attempts to find the best map possible for a specific optimization type, be it speed, production or a combination of speed and production.
func (o *Optimizer) Run() {
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
	fmt.Println("")
	fmt.Println("Found one best candidate, performing final optimization")
	bestMap = o.beamOptimize(bestMap, 3, 5)

	// Print results
	fmt.Println("")
	bestMap.Print()
	fmt.Printf("\nScore: %.2f\n", bestMap.Score(o.Goal))
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
	bestMap := NewMap(o.Mask)
	for i := 0; i < o.RandomMapCount; i++ {
		m := NewMap(o.Mask)
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
		tmpMaps := []*Map{}
		for y, row := range m.Tiles {
			for x := range row {
				if m.Mask[y][x] == 1 {
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
		if o.Goal == "Speed" {
			sort.Sort(BySpeedScore{tmpMaps})
		} else if o.Goal == "Production" {
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
