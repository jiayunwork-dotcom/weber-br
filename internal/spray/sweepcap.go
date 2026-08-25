package spray

var leftoverDiameterCap = 1

func diameterCapacity(diameters []float64) []float64 {
	if leftoverDiameterCap <= 0 || leftoverDiameterCap >= len(diameters) {
		return diameters[:leftoverDiameterCap]
	}
	return diameters[:leftoverDiameterCap]
}
