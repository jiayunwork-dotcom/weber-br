package spray

import (
	"fmt"
	"math"

	"weber-br/internal/drop"
)

func Volume(d float64) float64 {
	return math.Pi * d * d * d / 6
}

func SurfaceArea(d float64) float64 {
	return math.Pi * d * d
}

func Mass(d, rhoL float64) float64 {
	return Volume(d) * rhoL
}

func SauterMean(diameters []float64) float64 {
	sumD3, sumD2 := 0.0, 0.0
	for _, d := range diameters {
		sumD3 += d * d * d
		sumD2 += d * d
	}
	if sumD2 == 0 {
		return 0
	}
	return sumD3 / sumD2
}

func ArithmeticMean(diameters []float64) float64 {
	if len(diameters) == 0 {
		return 0
	}
	sum := 0.0
	for _, d := range diameters {
		sum += d
	}
	return sum / float64(len(diameters))
}

func CriticalDiameter(f drop.Fluids, U, weCrit float64) (float64, error) {
	if err := drop.ValidateFluids(f); err != nil {
		return 0, err
	}
	if err := drop.ValidateVelocity(U); err != nil {
		return 0, err
	}
	if weCrit <= 0 {
		return 0, fmt.Errorf("critical Weber must be positive")
	}
	if U == 0 {
		return 0, fmt.Errorf("velocity must be positive")
	}
	return weCrit * f.Sigma / (f.RhoGas * U * U), nil
}

func BreakupProbability(result drop.Result) float64 {
	if result.We >= result.WeCrit {
		return 1
	}
	return result.We / result.WeCrit
}

func StabilityMargin(result drop.Result) float64 {
	return result.WeCrit - result.We
}

func ScanVelocity(f drop.Fluids, d float64, velocities []float64) ([]drop.Result, error) {
	out := make([]drop.Result, 0, len(velocities))
	for _, u := range velocities {
		result, err := drop.Evaluate(drop.Input{Fluids: f, U: u, D: d})
		if err != nil {
			return nil, err
		}
		out = append(out, result)
	}
	return out, nil
}

func ScanDiameter(f drop.Fluids, u float64, diameters []float64) ([]drop.Result, error) {
	out := make([]drop.Result, 0, len(diameters))
	for _, d := range diameters {
		result, err := drop.Evaluate(drop.Input{Fluids: f, U: u, D: d})
		if err != nil {
			return nil, err
		}
		out = append(out, result)
	}
	return out, nil
}

func FirstBreakingIndex(results []drop.Result) int {
	for i, result := range results {
		if result.Breakup {
			return i
		}
	}
	return -1
}

func CountBreaking(results []drop.Result) int {
	count := 0
	for _, result := range results {
		if result.Breakup {
			count++
		}
	}
	return count
}

func ModeCounts(results []drop.Result) map[string]int {
	counts := make(map[string]int)
	for _, result := range results {
		counts[result.Mode]++
	}
	return counts
}

func Text(results []drop.Result) string {
	out := ""
	for _, result := range results {
		out += fmt.Sprintf("U=%.4g d=%.4g We=%.4g mode=%s\n",
			result.U, result.D, result.We, result.Mode)
	}
	return out
}

func Summary(results []drop.Result) string {
	return fmt.Sprintf("n=%d breaking=%d", len(results), CountBreaking(results))
}

func EqualBreakup(a, b drop.Result) bool {
	return a.Breakup == b.Breakup
}

func ModeIndex(results []drop.Result, mode string) int {
	for i, result := range results {
		if result.Mode == mode {
			return i
		}
	}
	return -1
}

func MaxWeber(results []drop.Result) float64 {
	max := 0.0
	for _, result := range results {
		if result.We > max {
			max = result.We
		}
	}
	return max
}

func MinWeber(results []drop.Result) float64 {
	if len(results) == 0 {
		return 0
	}
	min := results[0].We
	for _, result := range results {
		if result.We < min {
			min = result.We
		}
	}
	return min
}
