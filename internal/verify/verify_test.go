package verify

import "testing"

func TestRunAllChecks(t *testing.T) {
	checks := RunAll()
	if !AllPass(checks) {
		for _, check := range checks {
			if !check.OK {
				t.Logf("%s: %s", check.Name, check.Message)
			}
		}
		t.Fatal("not all checks pass")
	}
}

func TestCheckZeroVelocity(t *testing.T) {
	if !CheckZeroVelocity().OK {
		t.Fatal("zero velocity check failed")
	}
}

func TestCheckVelocityQuadratic(t *testing.T) {
	if !CheckVelocityQuadratic().OK {
		t.Fatal("quadratic check failed")
	}
}

func TestCheckSigmaHalvesWe(t *testing.T) {
	if !CheckSigmaHalvesWe().OK {
		t.Fatal("sigma check failed")
	}
}

func TestCheckDiameterTrend(t *testing.T) {
	if !CheckDiameterRaisesWeLowersOh().OK {
		t.Fatal("diameter check failed")
	}
}

func TestCheckGasDensityUsed(t *testing.T) {
	if !CheckGasDensityUsed().OK {
		t.Fatal("gas density check failed")
	}
}

func TestCheckModeCrossesThreshold(t *testing.T) {
	if !CheckModeCrossesThreshold().OK {
		t.Fatal("mode crossing check failed")
	}
}

func TestCheckCriticalU(t *testing.T) {
	if !CheckCriticalU().OK {
		t.Fatal("critical U check failed")
	}
}

func TestCheckValidation(t *testing.T) {
	if !CheckValidation().OK {
		t.Fatal("validation check failed")
	}
}

func TestCheckWeberFormula(t *testing.T) {
	if !CheckWeberFormula().OK {
		t.Fatal("weber formula check failed")
	}
}

func TestCheckRainDropExample(t *testing.T) {
	if !CheckRainDropExample().OK {
		t.Fatal("rain drop check failed")
	}
}
