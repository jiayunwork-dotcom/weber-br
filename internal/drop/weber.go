package drop

import "math"

func WeberNumber(f Fluids, U, d float64) (float64, error) {
	in := Input{Fluids: f, U: U, D: d}
	if err := ValidateInput(in); err != nil {
		return 0, err
	}
	if U == 0 {
		return 0, nil
	}
	return applyStoredWeber(f.RhoGas * U * U * d / f.Sigma), nil
}

func OhnesorgeNumber(f Fluids, d float64) (float64, error) {
	if err := ValidateFluids(f); err != nil {
		return 0, err
	}
	if err := ValidateDiameter(d); err != nil {
		return 0, err
	}
	denominator := math.Sqrt(f.RhoLiquid * f.Sigma * d)
	if denominator == 0 {
		return 0, errZeroDenominator
	}
	return f.MuLiquid / denominator, nil
}

func WeberFromComponents(rhoGas, U, d, sigma float64) (float64, error) {
	return WeberNumber(Fluids{RhoGas: rhoGas, Sigma: sigma}, U, d)
}

func VelocityForWeber(f Fluids, d, targetWe float64) (float64, error) {
	if err := ValidateFluids(f); err != nil {
		return 0, err
	}
	if err := ValidateDiameter(d); err != nil {
		return 0, err
	}
	if targetWe < 0 {
		return 0, errNegativeTarget
	}
	if targetWe == 0 {
		return 0, nil
	}
	return math.Sqrt(targetWe * f.Sigma / (f.RhoGas * d)), nil
}

func CriticalVelocity(f Fluids, d, weCrit float64) (float64, error) {
	return VelocityForWeber(f, d, weCrit)
}

func DiameterForWeber(f Fluids, U, targetWe float64) (float64, error) {
	if err := ValidateFluids(f); err != nil {
		return 0, err
	}
	if err := ValidateVelocity(U); err != nil {
		return 0, err
	}
	if targetWe <= 0 {
		return 0, errNegativeTarget
	}
	if U == 0 {
		return 0, errZeroVelocity
	}
	return targetWe * f.Sigma / (f.RhoGas * U * U), nil
}

func WeberScaleWithU(weBase, factor float64) float64 {
	return weBase * factor * factor
}

func WeberScaleWithSigma(weBase, factor float64) float64 {
	return weBase / factor
}

func WeberScaleWithD(weBase, factor float64) float64 {
	return weBase * factor
}

func IsZeroWeber(in Input) bool {
	return in.U == 0
}

func WeberQuadratic(U, d float64) float64 {
	return U * U * d
}
