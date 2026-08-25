package spray

import (
	"testing"

	"weber-br/internal/drop"
)

func fluids() drop.Fluids {
	return drop.Fluids{RhoGas: 1.2, RhoLiquid: 1000, MuLiquid: 0.001, Sigma: 0.072}
}

func TestVolumeAndArea(t *testing.T) {
	if Volume(0.002) <= 0 || SurfaceArea(0.002) <= 0 {
		t.Fatal("volume/area invalid")
	}
}

func TestSauterMean(t *testing.T) {
	value := SauterMean([]float64{0.001, 0.002})
	if value <= 0 {
		t.Fatal("SMD invalid")
	}
}

func TestScanVelocity(t *testing.T) {
	results, err := ScanVelocity(fluids(), 0.002, []float64{1, 10, 20})
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 3 {
		t.Fatalf("results=%d", len(results))
	}
}

func TestCountBreaking(t *testing.T) {
	results, _ := ScanVelocity(fluids(), 0.002, []float64{1, 10, 20})
	if CountBreaking(results) == 0 {
		t.Fatal("expected breaking cases")
	}
}

func TestCriticalDiameter(t *testing.T) {
	value, err := CriticalDiameter(fluids(), 20, 12)
	if err != nil {
		t.Fatal(err)
	}
	if value <= 0 {
		t.Fatalf("d=%g", value)
	}
}

func TestBuildSweep(t *testing.T) {
	sweep, err := BuildSweep(fluids(), []float64{1, 10}, []float64{0.001, 0.002})
	if err != nil {
		t.Fatal(err)
	}
	if SweepCount(sweep) != 4 {
		t.Fatalf("count=%d", SweepCount(sweep))
	}
	if !IsFinite(sweep) {
		t.Fatal("sweep not finite")
	}
}

func TestSweepSummary(t *testing.T) {
	sweep, _ := BuildSweep(fluids(), []float64{1, 10}, []float64{0.001, 0.002})
	if SweepSummary(sweep) == "" {
		t.Fatal("empty summary")
	}
}

func TestBreakupProbability(t *testing.T) {
	result, _ := drop.Evaluate(drop.Input{Fluids: fluids(), U: 20, D: 0.002})
	if BreakupProbability(result) != 1 {
		t.Fatal("probability not 1")
	}
}
