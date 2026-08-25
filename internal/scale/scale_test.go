package scale

import (
	"testing"

	"weber-br/internal/drop"
)

func TestVelocityScan(t *testing.T) {
	f := drop.Fluids{RhoGas: 1.2, RhoLiquid: 1000, MuLiquid: 0.001, Sigma: 0.072}
	rows, err := VelocityScan(f, 0.002, []float64{1, 10, 20})
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 3 {
		t.Fatalf("rows=%d", len(rows))
	}
}

func TestCSVAndText(t *testing.T) {
	f := drop.Fluids{RhoGas: 1.2, RhoLiquid: 1000, MuLiquid: 0.001, Sigma: 0.072}
	rows, _ := VelocityScan(f, 0.002, []float64{1, 10})
	if CSV(rows) == "" || Text(rows) == "" {
		t.Fatal("empty output")
	}
}

func TestCrossings(t *testing.T) {
	f := drop.Fluids{RhoGas: 1.2, RhoLiquid: 1000, MuLiquid: 0.001, Sigma: 0.072}
	rows, _ := VelocityScan(f, 0.002, []float64{1, 5, 10, 20, 40})
	if len(Crossings(rows)) == 0 {
		t.Fatal("no crossings")
	}
}

func TestCriticalVelocityForMode(t *testing.T) {
	f := drop.Fluids{RhoGas: 1.2, RhoLiquid: 1000, MuLiquid: 0.001, Sigma: 0.072}
	value, err := CriticalVelocityForMode(f, 0.002, 0.01, "bag")
	if err != nil {
		t.Fatal(err)
	}
	if value <= 0 {
		t.Fatalf("U=%g", value)
	}
}
