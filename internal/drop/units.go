package drop

import "fmt"

func MeterToMillimeter(value float64) float64 {
	return value * 1000
}

func MillimeterToMeter(value float64) float64 {
	return value / 1000
}

func MicrometerToMeter(value float64) float64 {
	return value * 1e-6
}

func MeterToMicrometer(value float64) float64 {
	return value * 1e6
}

func PascalSecondToCentipoise(value float64) float64 {
	return value * 1000
}

func CentipoiseToPascalSecond(value float64) float64 {
	return value / 1000
}

func NewtonPerMeterToDynePerCm(value float64) float64 {
	return value * 1000
}

func DynePerCmToNewtonPerMeter(value float64) float64 {
	return value / 1000
}

func FormatWe(we float64) string {
	return fmt.Sprintf("%.4g", we)
}

func FormatOh(oh float64) string {
	return fmt.Sprintf("%.4g", oh)
}

func FormatVelocity(value float64) string {
	return fmt.Sprintf("%.4g m/s", value)
}

func FormatDiameter(value float64) string {
	return fmt.Sprintf("%.4g m", value)
}

func FormatMode(mode string) string {
	return mode
}

func FormatBreakup(value bool) string {
	if value {
		return "yes"
	}
	return "no"
}

func Display(result Result) string {
	return FormatResult(result)
}

func DensityText(rhoGas, rhoLiquid float64) string {
	return fmt.Sprintf("rho_g=%.4g kg/m3 rho_l=%.4g kg/m3", rhoGas, rhoLiquid)
}

func ViscosityText(mu float64) string {
	return fmt.Sprintf("mu_l=%.4g Pa.s", mu)
}

func SurfaceTensionText(sigma float64) string {
	return fmt.Sprintf("sigma=%.4g N/m", sigma)
}

func DisplayWeber(result Result) string {
	return FormatWe(result.We)
}

func DisplayOh(result Result) string {
	return FormatOh(result.Oh)
}

func DisplayCriticalU(result Result) string {
	return FormatVelocity(result.CriticalU)
}

func DisplayMode(result Result) string {
	return FormatMode(result.Mode)
}

func DisplayBreakup(result Result) string {
	return FormatBreakup(result.Breakup)
}
