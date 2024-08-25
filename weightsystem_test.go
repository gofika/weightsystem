package weightsystem

import (
	"math"
	"testing"
)

const epsilon = 1e-6

func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}

func TestWeightSystem(t *testing.T) {
	ws := New[string]()

	ws.AddItems([]string{"a", "b", "c"})

	expectedInitialWeight := 500.5
	for _, item := range []string{"a", "b", "c"} {
		if !floatEquals(ws.weights[item], expectedInitialWeight) {
			t.Errorf("Expected initial weight of %s to be close to %.6f, got %.6f", item, expectedInitialWeight, ws.weights[item])
		}
	}

	expectedTotalWeight := 1501.5
	if !floatEquals(ws.totalWeight, expectedTotalWeight) {
		t.Errorf("Expected total weight to be close to %.6f, got %.6f", expectedTotalWeight, ws.totalWeight)
	}

	expectedAvgWeight := 500.5
	if !floatEquals(ws.avgWeight, expectedAvgWeight) {
		t.Errorf("Expected average weight to be close to %.6f, got %.6f", expectedAvgWeight, ws.avgWeight)
	}

	ws.AdjustWeight("a", true)

	expectedAdjustedWeight := 550.55
	if !floatEquals(ws.weights["a"], expectedAdjustedWeight) {
		t.Errorf("Expected weight of item 'a' to be close to %.6f, got %.6f", expectedAdjustedWeight, ws.weights["a"])
	}

	expectedNewTotalWeight := 1551.55
	if !floatEquals(ws.totalWeight, expectedNewTotalWeight) {
		t.Errorf("Expected total weight to be close to %.6f, got %.6f", expectedNewTotalWeight, ws.totalWeight)
	}

	expectedNewAvgWeight := 517.1833333333334
	if !floatEquals(ws.avgWeight, expectedNewAvgWeight) {
		t.Errorf("Expected average weight to be close to %.6f, got %.6f", expectedNewAvgWeight, ws.avgWeight)
	}

	ws.AddItem("d", WithWeight(1000))

	expectedDWeight := 1000.0
	if !floatEquals(ws.weights["d"], expectedDWeight) {
		t.Errorf("Expected weight of item 'd' to be %.6f, got %.6f", expectedDWeight, ws.weights["d"])
	}

	expectedFinalTotalWeight := 2551.55
	if !floatEquals(ws.totalWeight, expectedFinalTotalWeight) {
		t.Errorf("Expected total weight to be close to %.6f, got %.6f", expectedFinalTotalWeight, ws.totalWeight)
	}

	expectedFinalAvgWeight := 637.8875
	if !floatEquals(ws.avgWeight, expectedFinalAvgWeight) {
		t.Errorf("Expected average weight to be close to %.6f, got %.6f", expectedFinalAvgWeight, ws.avgWeight)
	}
}
