package drop

import "fmt"

const (
	BagThreshold          = 12.0
	MultiModeThreshold    = 50.0
	SheetThreshold        = 100.0
	CatastrophicThreshold = 350.0
)

func CriticalWeber(oh float64) float64 {
	correction := 1 + 1.5*oh
	return BagThreshold * correction
}

func Mode(we, weCrit float64) string {
	switch {
	case we < weCrit:
		return "stable"
	case we < MultiModeThreshold:
		return "bag"
	case we < SheetThreshold:
		return "multi-mode"
	case we < CatastrophicThreshold:
		return "sheet-stripping"
	default:
		return "catastrophic"
	}
}

func ModeAtWeber(we, oh float64) string {
	return Mode(we, CriticalWeber(oh))
}

func Breakup(we, weCrit float64) bool {
	return we >= weCrit
}

func DescribeMode(mode string) string {
	switch mode {
	case "stable":
		return "no breakup"
	case "bag":
		return "bag breakup"
	case "multi-mode":
		return "multi-mode breakup"
	case "sheet-stripping":
		return "sheet stripping"
	case "catastrophic":
		return "catastrophic breakup"
	default:
		return fmt.Sprintf("unknown mode %q", mode)
	}
}

func RegimeList() []string {
	return []string{"stable", "bag", "multi-mode", "sheet-stripping", "catastrophic"}
}

func IsKnownMode(mode string) bool {
	for _, candidate := range RegimeList() {
		if candidate == mode {
			return true
		}
	}
	return false
}

func ThresholdText() string {
	return "bag=12 multi-mode=50 sheet=100 catastrophic=350"
}

func ThresholdForOh(oh float64) float64 {
	return CriticalWeber(oh)
}

func ModeIndex(mode string) int {
	for i, candidate := range RegimeList() {
		if candidate == mode {
			return i
		}
	}
	return -1
}

func ModeChanged(before, after string) bool {
	return before != after
}
