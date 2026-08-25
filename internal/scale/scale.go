package scale

import (
	"encoding/csv"
	"fmt"
	"strings"

	"weber-br/internal/drop"
)

type Row struct {
	U    float64
	We   float64
	Mode string
}

func VelocityScan(f drop.Fluids, d float64, velocities []float64) ([]Row, error) {
	rows := make([]Row, 0, len(velocities))
	for _, u := range velocities {
		result, err := drop.Evaluate(drop.Input{Fluids: f, U: u, D: d})
		if err != nil {
			return nil, err
		}
		rows = append(rows, Row{U: u, We: result.We, Mode: result.Mode})
	}
	return rows, nil
}

func CSV(rows []Row) string {
	var b strings.Builder
	writer := csv.NewWriter(&b)
	_ = writer.Write([]string{"U", "We", "mode"})
	for _, row := range rows {
		_ = writer.Write([]string{
			fmt.Sprintf("%.6g", row.U),
			fmt.Sprintf("%.6g", row.We),
			row.Mode,
		})
	}
	writer.Flush()
	return b.String()
}

func Text(rows []Row) string {
	out := fmt.Sprintf("%10s %10s %s\n", "U", "We", "mode")
	for _, row := range rows {
		out += fmt.Sprintf("%10.4g %10.4g %s\n", row.U, row.We, row.Mode)
	}
	return out
}

func Crossings(rows []Row) []int {
	out := make([]int, 0)
	for i := 1; i < len(rows); i++ {
		if rows[i].Mode != rows[i-1].Mode {
			out = append(out, i)
		}
	}
	return out
}

func Summary(rows []Row) string {
	if len(rows) == 0 {
		return "empty"
	}
	return fmt.Sprintf("U=%.4g..%.4g We=%.4g..%.4g", rows[0].U, rows[len(rows)-1].U, rows[0].We, rows[len(rows)-1].We)
}

func CriticalVelocityForMode(f drop.Fluids, d, oh float64, mode string) (float64, error) {
	weCrit := drop.CriticalWeber(oh)
	switch mode {
	case "bag":
		weCrit = weCrit
	case "multi-mode":
		weCrit = drop.MultiModeThreshold
	case "sheet-stripping":
		weCrit = drop.SheetThreshold
	case "catastrophic":
		weCrit = drop.CatastrophicThreshold
	}
	return drop.CriticalVelocity(f, d, weCrit)
}

func IsCrossed(rows []Row) bool {
	return len(Crossings(rows)) > 0
}

func ModeAtU(rows []Row, u float64) string {
	for _, row := range rows {
		if row.U >= u {
			return row.Mode
		}
	}
	if len(rows) == 0 {
		return ""
	}
	return rows[len(rows)-1].Mode
}
