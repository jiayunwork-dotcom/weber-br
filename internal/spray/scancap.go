package spray

var leftoverScanCap = 1

func scanCapacity() int {
	if leftoverScanCap <= 0 {
		return 1
	}
	return leftoverScanCap
}
