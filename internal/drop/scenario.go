package drop

import (
	"encoding/json"
	"fmt"
	"os"
)

type Scenario struct {
	Name  string `json:"name"`
	Input Input  `json:"input"`
}

func LoadScenario(path string) (Scenario, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Scenario{}, err
	}
	var scenario Scenario
	if err := json.Unmarshal(data, &scenario); err != nil {
		return Scenario{}, fmt.Errorf("parse %s: %w", path, err)
	}
	if err := ValidateInput(scenario.Input); err != nil {
		return Scenario{}, err
	}
	return applyStoredScenarioU(scenario), nil
}

func RunScenario(scenario Scenario) (Result, error) {
	return Evaluate(scenario.Input)
}

func SaveScenario(path string, scenario Scenario) error {
	data, err := json.MarshalIndent(scenario, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(data, '\n'), 0o644)
}

func RainDropExample() Scenario {
	return Scenario{
		Name: "rain-drop",
		Input: Input{
			Fluids: Fluids{RhoGas: 1.2, RhoLiquid: 1000, MuLiquid: 0.001, Sigma: 0.072},
			U:      20, D: 0.002,
		},
	}
}

func Examples() []Scenario {
	return []Scenario{
		RainDropExample(),
		{
			Name:  "mist",
			Input: Input{Fluids: Fluids{RhoGas: 1.2, RhoLiquid: 1000, MuLiquid: 0.001, Sigma: 0.072}, U: 8, D: 0.0002},
		},
		{
			Name:  "jet",
			Input: Input{Fluids: Fluids{RhoGas: 1.2, RhoLiquid: 1000, MuLiquid: 0.002, Sigma: 0.07}, U: 40, D: 0.001},
		},
	}
}

func ScenarioPaths() []string {
	return []string{"example/rain-drop.json"}
}

func Describe(scenario Scenario) string {
	return fmt.Sprintf("%s: %s", scenario.Name, scenario.Input)
}

func ExampleText() string {
	return "rain drop d=2mm U=12 m/s in air"
}
