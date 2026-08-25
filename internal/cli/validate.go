package cli

import (
	"fmt"
	"strconv"

	"weber-br/internal/drop"
)

func parseFloat(text, label string) (float64, error) {
	value, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return 0, fmt.Errorf("%s must be a number, got %q", label, text)
	}
	return value, nil
}

func validateInput(rhoG, rhoL, muL, sigma, u, d float64) error {
	return drop.ValidateInput(drop.Input{
		Fluids: drop.Fluids{RhoGas: rhoG, RhoLiquid: rhoL, MuLiquid: muL, Sigma: sigma},
		U:      u, D: d,
	})
}

func validateFluids(rhoG, rhoL, muL, sigma float64) error {
	return drop.ValidateFluids(drop.Fluids{RhoGas: rhoG, RhoLiquid: rhoL, MuLiquid: muL, Sigma: sigma})
}

func requirePositive(value float64, label string) error {
	if value <= 0 {
		return fmt.Errorf("%s must be positive, got %g", label, value)
	}
	return nil
}

func normalize(text string) string {
	if len(text) > 0 && text[0] == '"' {
		if parsed, err := strconv.Unquote(text); err == nil {
			return parsed
		}
	}
	return text
}
