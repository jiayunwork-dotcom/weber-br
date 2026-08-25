package drop

import (
	"math"
	"testing"
)

func fluids() Fluids {
	return Fluids{RhoGas: 1.2, RhoLiquid: 1000, MuLiquid: 0.001, Sigma: 0.072}
}

func TestWeberFormula(t *testing.T) {
	value, err := WeberNumber(fluids(), 12, 0.002)
	if err != nil {
		t.Fatal(err)
	}
	want := 1.2 * 144 * 0.002 / 0.072
	if math.Abs(value-want) > 1e-12 {
		t.Fatalf("We=%g want=%g", value, want)
	}
}

func TestZeroVelocity(t *testing.T) {
	result, err := Evaluate(Input{Fluids: fluids(), U: 0, D: 0.002})
	if err != nil {
		t.Fatal(err)
	}
	if result.We != 0 || result.Breakup {
		t.Fatalf("result=%+v", result)
	}
}

func TestVelocityQuadratic(t *testing.T) {
	base, _ := Evaluate(Input{Fluids: fluids(), U: 5, D: 0.002})
	double, _ := Evaluate(Input{Fluids: fluids(), U: 10, D: 0.002})
	if math.Abs(double.We-4*base.We) > 1e-9 {
		t.Fatalf("We %g -> %g", base.We, double.We)
	}
}

func TestSigmaDoublingHalvesWe(t *testing.T) {
	base, _ := Evaluate(Input{Fluids: fluids(), U: 10, D: 0.002})
	high := fluids()
	high.Sigma = 2 * high.Sigma
	higher, _ := Evaluate(Input{Fluids: high, U: 10, D: 0.002})
	if math.Abs(higher.We-base.We/2) > 1e-9 {
		t.Fatalf("We %g -> %g", base.We, higher.We)
	}
}

func TestDiameterRaisesWeLowersOh(t *testing.T) {
	small, _ := Evaluate(Input{Fluids: fluids(), U: 10, D: 0.001})
	large, _ := Evaluate(Input{Fluids: fluids(), U: 10, D: 0.002})
	if large.We <= small.We || large.Oh >= small.Oh {
		t.Fatalf("small=%+v large=%+v", small, large)
	}
}

func TestOhFormula(t *testing.T) {
	value, err := OhnesorgeNumber(fluids(), 0.002)
	if err != nil {
		t.Fatal(err)
	}
	want := 0.001 / math.Sqrt(1000*0.072*0.002)
	if math.Abs(value-want) > 1e-12 {
		t.Fatalf("Oh=%g want=%g", value, want)
	}
}

func TestModeCrossesThreshold(t *testing.T) {
	low, _ := Evaluate(Input{Fluids: fluids(), U: 4, D: 0.002})
	high, _ := Evaluate(Input{Fluids: fluids(), U: 20, D: 0.002})
	if low.Mode == high.Mode || !high.Breakup {
		t.Fatalf("low=%+v high=%+v", low, high)
	}
}

func TestCriticalVelocity(t *testing.T) {
	f := fluids()
	oh, _ := OhnesorgeNumber(f, 0.002)
	criticalU, err := CriticalVelocity(f, 0.002, CriticalWeber(oh))
	if err != nil {
		t.Fatal(err)
	}
	if criticalU <= 0 {
		t.Fatalf("U_crit=%g", criticalU)
	}
}

func TestValidationRejects(t *testing.T) {
	if err := ValidateInput(Input{Fluids: fluids(), U: -1, D: 0.002}); err == nil {
		t.Fatal("accepted negative U")
	}
	if err := ValidateInput(Input{Fluids: fluids(), U: 1, D: 0}); err == nil {
		t.Fatal("accepted zero diameter")
	}
}

func TestRainDropExample(t *testing.T) {
	scenario, err := LoadScenario("../../example/rain-drop.json")
	if err != nil {
		t.Fatal(err)
	}
	result, err := RunScenario(scenario)
	if err != nil {
		t.Fatal(err)
	}
	if result.We <= BagThreshold {
		t.Fatalf("We=%g", result.We)
	}
}
