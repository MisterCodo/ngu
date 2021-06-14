package maps

import (
	"math/rand"
	"testing"

	"github.com/MisterCodo/ngu/plugins/beacons"
	"github.com/stretchr/testify/assert"
)

func TestOptimizationGoal(t *testing.T) {
	assert.Equal(t, "SpeedAndProduction", OptimizationGoal(SpeedAndProductionGoal).String())
	assert.Equal(t, "Speed", OptimizationGoal(SpeedGoal).String())
	assert.Equal(t, "Production", OptimizationGoal(ProductionGoal).String())
}

func TestOptimizerRun(t *testing.T) {
	o, err := NewOptimizer(SpeedGoal, []beacons.BType{beacons.Box, beacons.Knight, beacons.Arrow, beacons.Wall, beacons.Donut}, "CandyLand")
	assert.NoError(t, err)
	o.Infinite = false

	// Force always the same random seed so test results are static. This doesn't work, likely a map order problem
	rand.Seed(7)

	_, err = o.Run(false)
	assert.NoError(t, err)
}
