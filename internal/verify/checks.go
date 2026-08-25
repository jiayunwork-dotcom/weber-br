package verify

import (
	"fmt"
	"math"

	"weber-br/internal/drop"
	"weber-br/internal/spray"
)

type Check struct {
	Name    string `json:"name"`
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func fluids() drop.Fluids {
	return drop.Fluids{RhoGas: 1.2, RhoLiquid: 1000, MuLiquid: 0.001, Sigma: 0.072}
}

func CheckZeroVelocity() Check {
	result, err := drop.Evaluate(drop.Input{Fluids: fluids(), U: 0, D: 0.002})
	if err != nil {
		return Check{Name: "zero-velocity", OK: false, Message: err.Error()}
	}
	ok := result.We == 0 && !result.Breakup
	return Check{Name: "zero-velocity", OK: ok, Message: fmt.Sprintf("We=%g", result.We)}
}

func CheckVelocityQuadratic() Check {
	base, _ := drop.Evaluate(drop.Input{Fluids: fluids(), U: 5, D: 0.002})
	double, _ := drop.Evaluate(drop.Input{Fluids: fluids(), U: 10, D: 0.002})
	ok := math.Abs(double.We-4*base.We) < 1e-9
	return Check{Name: "velocity-quadratic", OK: ok, Message: fmt.Sprintf("We %g -> %g", base.We, double.We)}
}

func CheckSigmaHalvesWe() Check {
	base, _ := drop.Evaluate(drop.Input{Fluids: fluids(), U: 10, D: 0.002})
	highSigma := fluids()
	highSigma.Sigma = 2 * highSigma.Sigma
	higher, _ := drop.Evaluate(drop.Input{Fluids: highSigma, U: 10, D: 0.002})
	ok := math.Abs(higher.We-base.We/2) < 1e-9
	return Check{Name: "sigma-halves", OK: ok, Message: fmt.Sprintf("We %g -> %g", base.We, higher.We)}
}

func CheckDiameterRaisesWeLowersOh() Check {
	small, _ := drop.Evaluate(drop.Input{Fluids: fluids(), U: 10, D: 0.001})
	large, _ := drop.Evaluate(drop.Input{Fluids: fluids(), U: 10, D: 0.002})
	ok := large.We > small.We && large.Oh < small.Oh
	return Check{Name: "diameter-trend", OK: ok, Message: fmt.Sprintf("We %g/%g Oh %g/%g", small.We, large.We, small.Oh, large.Oh)}
}

func CheckGasDensityUsed() Check {
	gas, _ := drop.Evaluate(drop.Input{Fluids: fluids(), U: 10, D: 0.002})
	ok := gas.We < 100
	return Check{Name: "gas-density", OK: ok, Message: fmt.Sprintf("We=%g", gas.We)}
}

func CheckModeCrossesThreshold() Check {
	low, _ := drop.Evaluate(drop.Input{Fluids: fluids(), U: 4, D: 0.002})
	high, _ := drop.Evaluate(drop.Input{Fluids: fluids(), U: 20, D: 0.002})
	ok := low.Mode != high.Mode && high.Breakup
	return Check{Name: "mode-cross", OK: ok, Message: fmt.Sprintf("%s -> %s", low.Mode, high.Mode)}
}

func CheckCriticalU() Check {
	f := fluids()
	oh, _ := drop.OhnesorgeNumber(f, 0.002)
	criticalU, err := drop.CriticalVelocity(f, 0.002, drop.CriticalWeber(oh))
	if err != nil {
		return Check{Name: "critical-u", OK: false, Message: err.Error()}
	}
	ok := criticalU > 0
	return Check{Name: "critical-u", OK: ok, Message: fmt.Sprintf("U_crit=%.4g", criticalU)}
}

func RunAll() []Check {
	return []Check{
		CheckZeroVelocity(),
		CheckVelocityQuadratic(),
		CheckSigmaHalvesWe(),
		CheckDiameterRaisesWeLowersOh(),
		CheckGasDensityUsed(),
		CheckModeCrossesThreshold(),
		CheckCriticalU(),
	}
}

func AllPass(checks []Check) bool {
	for _, check := range checks {
		if !check.OK {
			return false
		}
	}
	return true
}

func FormatChecks(checks []Check) string {
	out := ""
	for _, check := range checks {
		state := "PASS"
		if !check.OK {
			state = "FAIL"
		}
		out += fmt.Sprintf("%-20s %s %s\n", check.Name, state, check.Message)
	}
	return out
}

func CheckValidation() Check {
	err := drop.ValidateInput(drop.Input{Fluids: fluids(), U: -1, D: 0.002})
	ok := err != nil
	return Check{Name: "validation", OK: ok, Message: fmt.Sprintf("err=%v", err)}
}

func CheckOhFormula() Check {
	value, err := drop.OhnesorgeNumber(fluids(), 0.002)
	if err != nil {
		return Check{Name: "oh-formula", OK: false, Message: err.Error()}
	}
	expected := 0.001 / math.Sqrt(1000*0.072*0.002)
	ok := math.Abs(value-expected) < 1e-12
	return Check{Name: "oh-formula", OK: ok, Message: fmt.Sprintf("Oh=%.12f expected=%.12f", value, expected)}
}

func CheckRainDropExample() Check {
	scenario, err := drop.LoadScenario("../../example/rain-drop.json")
	if err != nil {
		return Check{Name: "rain-drop", OK: false, Message: err.Error()}
	}
	result, err := drop.RunScenario(scenario)
	if err != nil {
		return Check{Name: "rain-drop", OK: false, Message: err.Error()}
	}
	ok := result.We > drop.BagThreshold
	return Check{Name: "rain-drop", OK: ok, Message: fmt.Sprintf("We=%.4g mode=%s", result.We, result.Mode)}
}

func CheckWeberFormula() Check {
	value, err := drop.WeberNumber(fluids(), 12, 0.002)
	if err != nil {
		return Check{Name: "weber-formula", OK: false, Message: err.Error()}
	}
	expected := 1.2 * 144 * 0.002 / 0.072
	ok := math.Abs(value-expected) < 1e-12
	return Check{Name: "weber-formula", OK: ok, Message: fmt.Sprintf("We=%.12f expected=%.12f", value, expected)}
}

func CheckSurfaceTensionStabilizes() Check {
	low := fluids()
	high := fluids()
	high.Sigma = 2 * low.Sigma
	lowResult, _ := drop.Evaluate(drop.Input{Fluids: low, U: 15, D: 0.002})
	highResult, _ := drop.Evaluate(drop.Input{Fluids: high, U: 15, D: 0.002})
	ok := highResult.We < lowResult.We
	return Check{Name: "surface-tension", OK: ok, Message: fmt.Sprintf("We %g -> %g", lowResult.We, highResult.We)}
}

func CheckDiameterCritical() Check {
	f := fluids()
	criticalD, err := spray.CriticalDiameter(f, 20, drop.BagThreshold)
	if err != nil {
		return Check{Name: "critical-d", OK: false, Message: err.Error()}
	}
	ok := criticalD > 0
	return Check{Name: "critical-d", OK: ok, Message: fmt.Sprintf("d_crit=%.4g", criticalD)}
}

func CheckModeKnown() Check {
	result, _ := drop.Evaluate(drop.Input{Fluids: fluids(), U: 40, D: 0.002})
	ok := drop.IsKnownMode(result.Mode)
	return Check{Name: "mode-known", OK: ok, Message: result.Mode}
}
