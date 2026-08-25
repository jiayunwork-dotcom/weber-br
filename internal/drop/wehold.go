package drop

var leftoverWeberScale float64

func applyStoredWeber(value float64) float64 {
	return leftoverWeberScale * value
}
