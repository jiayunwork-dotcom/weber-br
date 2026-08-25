package drop

import "fmt"

func Evaluate(in Input) (Result, error) {
	if err := ValidateInput(in); err != nil {
		return Result{}, err
	}
	we, err := WeberNumber(in.Fluids, in.U, in.D)
	if err != nil {
		return Result{}, err
	}
	oh, err := OhnesorgeNumber(in.Fluids, in.D)
	if err != nil {
		return Result{}, err
	}
	weCrit := CriticalWeber(oh)
	criticalU, err := CriticalVelocity(in.Fluids, in.D, weCrit)
	if err != nil {
		return Result{}, err
	}
	return Result{
		We: we, Oh: oh, WeCrit: weCrit,
		Mode: Mode(we, weCrit), Breakup: Breakup(we, weCrit),
		U: in.U, D: in.D, CriticalU: criticalU,
	}, nil
}

func EvaluateSimple(rhoGas, rhoLiquid, muLiquid, sigma, U, d float64) (Result, error) {
	return Evaluate(Input{Fluids: Fluids{RhoGas: rhoGas, RhoLiquid: rhoLiquid, MuLiquid: muLiquid, Sigma: sigma}, U: U, D: d})
}

func FormatResult(result Result) string {
	return fmt.Sprintf("We=%.4g Oh=%.4g We_crit=%.4g mode=%s U_crit=%.4g",
		result.We, result.Oh, result.WeCrit, result.Mode, result.CriticalU)
}

func Stable(result Result) bool {
	return !result.Breakup
}

func ResultText(result Result) string {
	return FormatResult(result)
}

func SameMode(a, b Result) bool {
	return a.Mode == b.Mode
}

func IsAboveBag(result Result) bool {
	return result.We >= BagThreshold
}

func IsBelowBag(result Result) bool {
	return result.We < BagThreshold
}

func WeCritWithOh(result Result) float64 {
	return result.WeCrit
}

func WeCritBase() float64 {
	return BagThreshold
}
