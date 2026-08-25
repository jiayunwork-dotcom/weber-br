package cli

import (
	"testing"

	"weber-br/internal/drop"
)

func TestValidateInput(t *testing.T) {
	if err := validateInput(1.2, 1000, 0.001, 0.072, 12, 0.002); err != nil {
		t.Fatal(err)
	}
	if err := validateInput(1.2, 1000, 0.001, 0.072, -1, 0.002); err == nil {
		t.Fatal("accepted negative U")
	}
}

func TestValidateFluids(t *testing.T) {
	if err := validateFluids(1.2, 1000, 0.001, 0.072); err != nil {
		t.Fatal(err)
	}
	if err := validateFluids(0, 1000, 0.001, 0.072); err == nil {
		t.Fatal("accepted zero gas density")
	}
}

func TestLoadExample(t *testing.T) {
	scenario, err := loadExample("../../example/rain-drop.json")
	if err != nil {
		t.Fatal(err)
	}
	if scenario.Input.U != 20 {
		t.Fatalf("scenario=%+v", scenario)
	}
}

func TestExamplePaths(t *testing.T) {
	if len(ExamplePaths()) != 1 {
		t.Fatal("expected one example path")
	}
}

func TestEvaluateExample(t *testing.T) {
	scenario, _ := loadExample("../../example/rain-drop.json")
	result, err := drop.RunScenario(scenario)
	if err != nil {
		t.Fatal(err)
	}
	if result.We <= drop.BagThreshold {
		t.Fatalf("We=%g", result.We)
	}
}
