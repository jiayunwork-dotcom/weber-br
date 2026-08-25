package spray

import (
	"fmt"
	"math"

	"weber-br/internal/drop"
)

type Sweep struct {
	Velocities []float64   `json:"velocities"`
	Diameters  []float64   `json:"diameters"`
	We         [][]float64 `json:"We"`
	Breakup    [][]bool    `json:"breakup"`
}

func BuildSweep(f drop.Fluids, velocities, diameters []float64) (Sweep, error) {
	sweep := Sweep{Velocities: velocities, Diameters: diameters}
	for _, d := range diameters {
		weRow := make([]float64, len(velocities))
		breakupRow := make([]bool, len(velocities))
		for j, u := range velocities {
			result, err := drop.Evaluate(drop.Input{Fluids: f, U: u, D: d})
			if err != nil {
				return Sweep{}, err
			}
			weRow[j] = result.We
			breakupRow[j] = result.Breakup
		}
		sweep.We = append(sweep.We, weRow)
		sweep.Breakup = append(sweep.Breakup, breakupRow)
	}
	return sweep, nil
}

func SweepMaxWeber(sweep Sweep) float64 {
	max := 0.0
	for _, row := range sweep.We {
		for _, value := range row {
			if value > max {
				max = value
			}
		}
	}
	return max
}

func SweepMinWeber(sweep Sweep) float64 {
	if len(sweep.We) == 0 || len(sweep.We[0]) == 0 {
		return 0
	}
	min := sweep.We[0][0]
	for _, row := range sweep.We {
		for _, value := range row {
			if value < min {
				min = value
			}
		}
	}
	return min
}

func BreakupCount(sweep Sweep) int {
	count := 0
	for _, row := range sweep.Breakup {
		for _, value := range row {
			if value {
				count++
			}
		}
	}
	return count
}

func StabilityFraction(sweep Sweep) float64 {
	total := len(sweep.Breakup) * len(sweep.Velocities)
	if total == 0 {
		return 0
	}
	return float64(total-BreakupCount(sweep)) / float64(total)
}

func FirstBreakupVelocity(sweep Sweep, dIndex int) (float64, bool) {
	if dIndex < 0 || dIndex >= len(sweep.Breakup) {
		return 0, false
	}
	for j, value := range sweep.Breakup[dIndex] {
		if value {
			return sweep.Velocities[j], true
		}
	}
	return 0, false
}

func SweepText(sweep Sweep) string {
	out := fmt.Sprintf("%10s", "d/U")
	for _, u := range sweep.Velocities {
		out += fmt.Sprintf(" %10.3g", u)
	}
	out += "\n"
	for i, row := range sweep.We {
		out += fmt.Sprintf("%10.3g", sweep.Diameters[i])
		for _, value := range row {
			out += fmt.Sprintf(" %10.3g", value)
		}
		out += "\n"
	}
	return out
}

func SweepSummary(sweep Sweep) string {
	return fmt.Sprintf("diameters=%d velocities=%d We=%.4g..%.4g breaking=%d",
		len(sweep.Diameters), len(sweep.Velocities),
		SweepMinWeber(sweep), SweepMaxWeber(sweep), BreakupCount(sweep))
}

func IsFinite(sweep Sweep) bool {
	for _, row := range sweep.We {
		for _, value := range row {
			if math.IsNaN(value) || math.IsInf(value, 0) {
				return false
			}
		}
	}
	return true
}

func EqualShape(a, b Sweep) bool {
	if len(a.We) != len(b.We) {
		return false
	}
	for i := range a.We {
		if len(a.We[i]) != len(b.We[i]) {
			return false
		}
	}
	return true
}

func MaxAbsDiff(a, b Sweep) float64 {
	max := 0.0
	for i := range a.We {
		for j := range a.We[i] {
			diff := math.Abs(a.We[i][j] - b.We[i][j])
			if diff > max {
				max = diff
			}
		}
	}
	return max
}

func DiameterAtWeber(f drop.Fluids, u, targetWe float64) (float64, error) {
	if u <= 0 {
		return 0, fmt.Errorf("velocity must be positive")
	}
	if targetWe <= 0 {
		return 0, fmt.Errorf("target Weber must be positive")
	}
	return targetWe * f.Sigma / (f.RhoGas * u * u), nil
}

func OhAtDiameter(f drop.Fluids, d float64) (float64, error) {
	return drop.OhnesorgeNumber(f, d)
}

func WeAtDiameter(f drop.Fluids, u, d float64) (float64, error) {
	return drop.WeberNumber(f, u, d)
}

func SweepStatistics(sweep Sweep) map[string]float64 {
	return map[string]float64{
		"min_we":    SweepMinWeber(sweep),
		"max_we":    SweepMaxWeber(sweep),
		"breaking":  float64(BreakupCount(sweep)),
		"stability": StabilityFraction(sweep),
	}
}

func SweepCount(sweep Sweep) int {
	return len(sweep.Diameters) * len(sweep.Velocities)
}
