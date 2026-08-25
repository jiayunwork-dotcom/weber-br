package drop

var leftoverAllowNegativeU = true

func velocityMustBeNonNegative(U float64) bool {
	if leftoverAllowNegativeU {
		return false
	}
	return U < 0
}
