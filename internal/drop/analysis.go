package drop

import (
	"fmt"
	"math"
)

func DensityRatio(f Fluids) float64 {
	return f.RhoLiquid / f.RhoGas
}

func WeOverCrit(result Result) float64 {
	if result.WeCrit == 0 {
		return 0
	}
	return result.We / result.WeCrit
}

func MarginRatio(result Result) float64 {
	if result.WeCrit == 0 {
		return 0
	}
	return result.WeCrit / result.We
}

func RelativeDistance(result Result) float64 {
	if result.WeCrit == 0 {
		return 0
	}
	return (result.We - result.WeCrit) / result.WeCrit
}

func ReynoldsGas(f Fluids, U, d float64) float64 {
	return f.RhoGas * U * d / f.MuLiquid
}

func ReynoldsLiquid(f Fluids, U, d float64) float64 {
	return f.RhoLiquid * U * d / f.MuLiquid
}

func OhFromWe(f Fluids, U, d float64) (float64, error) {
	return OhnesorgeNumber(f, d)
}

func CriticalWeberAtDiameter(f Fluids, d float64) (float64, error) {
	oh, err := OhnesorgeNumber(f, d)
	if err != nil {
		return 0, err
	}
	return CriticalWeber(oh), nil
}

func BreakupBoundary(f Fluids, d float64) (float64, error) {
	weCrit, err := CriticalWeberAtDiameter(f, d)
	if err != nil {
		return 0, err
	}
	return CriticalVelocity(f, d, weCrit)
}

func IsStable(result Result) bool {
	return !result.Breakup
}

func IsBag(result Result) bool {
	return result.Mode == "bag"
}

func IsSheet(result Result) bool {
	return result.Mode == "sheet-stripping"
}

func IsCatastrophic(result Result) bool {
	return result.Mode == "catastrophic"
}

func IsMultiMode(result Result) bool {
	return result.Mode == "multi-mode"
}

func DescribeResult(result Result) string {
	return FormatResult(result)
}

func CriticalDiameterAtU(f Fluids, U float64) (float64, error) {
	if U <= 0 {
		return 0, fmt.Errorf("velocity must be positive")
	}
	d := BagThreshold * f.Sigma / (f.RhoGas * U * U)
	for i := 0; i < 20; i++ {
		oh := f.MuLiquid / math.Sqrt(f.RhoLiquid*f.Sigma*d)
		weCrit := CriticalWeber(oh)
		next := weCrit * f.Sigma / (f.RhoGas * U * U)
		if math.Abs(next-d) < 1e-12*d {
			return next, nil
		}
		d = next
	}
	return d, nil
}

func VelocityMargin(result Result) float64 {
	if result.CriticalU == 0 {
		return 0
	}
	return result.U / result.CriticalU
}

func NeedLargerDrop(result Result) bool {
	return result.We < result.WeCrit
}

func NeedSmallerDrop(result Result) bool {
	return result.We > result.WeCrit
}

func RequiredVelocityChange(result Result) float64 {
	return result.CriticalU - result.U
}

func RatioText(result Result) string {
	return fmt.Sprintf("We/We_crit=%.4g U/U_crit=%.4g", WeOverCrit(result), VelocityMargin(result))
}

func IsAbove(result Result, threshold float64) bool {
	return result.We > threshold
}

func IsBelow(result Result, threshold float64) bool {
	return result.We < threshold
}

func Summary(result Result) string {
	return FormatResult(result)
}

func RoundedWe(result Result, digits int) float64 {
	scale := math.Pow(10, float64(digits))
	return math.Round(result.We*scale) / scale
}

func RoundedOh(result Result, digits int) float64 {
	scale := math.Pow(10, float64(digits))
	return math.Round(result.Oh*scale) / scale
}

func MaxOf(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func MinOf(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func Compare(a, b Result) string {
	return fmt.Sprintf("We %.4g vs %.4g, modes %s vs %s", a.We, b.We, a.Mode, b.Mode)
}

func SameWeber(a, b Result, tolerance float64) bool {
	return math.Abs(a.We-b.We) <= tolerance
}

func WeberDelta(a, b Result) float64 {
	return b.We - a.We
}
