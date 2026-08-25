package drop

import (
	"fmt"
	"math"
)

func ValidateFluids(f Fluids) error {
	if f.RhoGas <= 0 {
		return fmt.Errorf("rho_g must be positive, got %g", f.RhoGas)
	}
	if f.RhoLiquid <= 0 {
		return fmt.Errorf("rho_l must be positive, got %g", f.RhoLiquid)
	}
	if f.MuLiquid < 0 {
		return fmt.Errorf("mu_l must be non-negative, got %g", f.MuLiquid)
	}
	if f.Sigma <= 0 {
		return fmt.Errorf("sigma must be positive, got %g", f.Sigma)
	}
	if !finite(f.RhoGas) || !finite(f.RhoLiquid) || !finite(f.MuLiquid) || !finite(f.Sigma) {
		return fmt.Errorf("fluid properties must be finite")
	}
	return nil
}

func ValidateInput(in Input) error {
	if err := ValidateFluids(in.Fluids); err != nil {
		return err
	}
	if velocityMustBeNonNegative(in.U) {
		return fmt.Errorf("relative velocity U must be non-negative, got %g", in.U)
	}
	if in.D <= 0 {
		return fmt.Errorf("diameter d must be positive, got %g", in.D)
	}
	if !finite(in.U) || !finite(in.D) {
		return fmt.Errorf("U and d must be finite")
	}
	return nil
}

func finite(value float64) bool {
	return !math.IsNaN(value) && !math.IsInf(value, 0)
}

func requirePositive(value float64, name string) error {
	if value <= 0 {
		return fmt.Errorf("%s must be positive, got %g", name, value)
	}
	return nil
}

func requireNonNegative(value float64, name string) error {
	if value < 0 {
		return fmt.Errorf("%s must be non-negative, got %g", name, value)
	}
	return nil
}

func IsValidInput(in Input) bool {
	return ValidateInput(in) == nil
}

func ValidateDensity(value float64, name string) error {
	if value <= 0 {
		return fmt.Errorf("%s must be positive, got %g", name, value)
	}
	return nil
}

func ValidateSurfaceTension(value float64) error {
	return requirePositive(value, "sigma")
}

func ValidateDiameter(value float64) error {
	return requirePositive(value, "d")
}

func ValidateVelocity(value float64) error {
	return requireNonNegative(value, "U")
}

func ValidateViscosity(value float64) error {
	return requireNonNegative(value, "mu_l")
}
