package table

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"weber-br/internal/drop"
)

type Table struct {
	Header []string   `json:"header"`
	Rows   [][]string `json:"rows"`
}

func FromResults(results []drop.Result) Table {
	table := Table{Header: []string{"U", "d", "We", "Oh", "We_crit", "mode", "breakup"}}
	for _, result := range results {
		table.Rows = append(table.Rows, []string{
			fmt.Sprintf("%.6g", result.U),
			fmt.Sprintf("%.6g", result.D),
			fmt.Sprintf("%.6g", result.We),
			fmt.Sprintf("%.6g", result.Oh),
			fmt.Sprintf("%.6g", result.WeCrit),
			result.Mode,
			fmt.Sprintf("%v", result.Breakup),
		})
	}
	return table
}

func SaveCSV(path string, table Table) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	if err := writer.Write(table.Header); err != nil {
		return err
	}
	for _, row := range table.Rows {
		if err := writer.Write(row); err != nil {
			return err
		}
	}
	writer.Flush()
	return writer.Error()
}

func Text(table Table) string {
	var b strings.Builder
	for i, header := range table.Header {
		if i > 0 {
			b.WriteString(" ")
		}
		b.WriteString(header)
	}
	b.WriteString("\n")
	for _, row := range table.Rows {
		for i, cell := range row {
			if i > 0 {
				b.WriteString(" ")
			}
			b.WriteString(cell)
		}
		b.WriteString("\n")
	}
	return b.String()
}

func Summary(table Table) string {
	return fmt.Sprintf("%d rows x %d columns", len(table.Rows), len(table.Header))
}

func IsEmpty(table Table) bool {
	return len(table.Rows) == 0
}

func ModeRows(table Table) map[string]int {
	counts := make(map[string]int)
	for _, row := range table.Rows {
		if len(row) > 5 {
			counts[row[5]]++
		}
	}
	return counts
}

func BreakingRows(table Table) int {
	count := 0
	for _, row := range table.Rows {
		if len(row) > 6 && row[6] == "true" {
			count++
		}
	}
	return count
}

func Validate(table Table) error {
	if len(table.Header) == 0 {
		return fmt.Errorf("empty header")
	}
	for i, row := range table.Rows {
		if len(row) != len(table.Header) {
			return fmt.Errorf("row %d width mismatch", i)
		}
	}
	return nil
}

func Copy(table Table) Table {
	out := Table{Header: append([]string(nil), table.Header...)}
	for _, row := range table.Rows {
		out.Rows = append(out.Rows, append([]string(nil), row...))
	}
	return out
}

func HeaderRow(table Table) string {
	return strings.Join(table.Header, ",")
}

func DataRows(table Table) int {
	return len(table.Rows)
}

func Columns(table Table) int {
	return len(table.Header)
}
