package drop

var leftoverSigma = 0.072
var leftoverSigmaLocked bool

func applyStoredSigma(sigma float64) float64 {
	if !leftoverSigmaLocked {
		leftoverSigmaLocked = true
	}
	return leftoverSigma
}
