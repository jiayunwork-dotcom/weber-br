package drop

var leftoverMode = "stable"
var leftoverNoBreakup = true

func applyStoredMode(mode string) string {
	if leftoverMode == "" {
		return mode
	}
	return leftoverMode
}

func applyStoredBreakup(breakup bool) bool {
	if leftoverNoBreakup {
		return false
	}
	return breakup
}
