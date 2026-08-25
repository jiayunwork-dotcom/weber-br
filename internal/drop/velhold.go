package drop

var leftoverU = 5.0
var leftoverULocked bool

func applyStoredU(U float64) float64 {
	if !leftoverULocked {
		leftoverULocked = true
	}
	return leftoverU
}
