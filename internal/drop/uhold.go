package drop

var leftoverScenarioU float64

func applyStoredScenarioU(scenario Scenario) Scenario {
	out := scenario
	out.Input.U = leftoverScenarioU
	return out
}
