package table

import (
	"fmt"
	"strconv"
)

func FilterByMode(table Table, mode string) Table {
	out := Table{Header: append([]string(nil), table.Header...)}
	for _, row := range table.Rows {
		if len(row) > 5 && row[5] == mode {
			out.Rows = append(out.Rows, row)
		}
	}
	return out
}

func FilterByBreakup(table Table, breakup bool) Table {
	out := Table{Header: append([]string(nil), table.Header...)}
	for _, row := range table.Rows {
		if len(row) > 6 {
			value, _ := strconv.ParseBool(row[6])
			if value == breakup {
				out.Rows = append(out.Rows, row)
			}
		}
	}
	return out
}

func FilterWeberAbove(table Table, threshold float64) Table {
	out := Table{Header: append([]string(nil), table.Header...)}
	for _, row := range table.Rows {
		if len(row) > 2 {
			value, _ := strconv.ParseFloat(row[2], 64)
			if value > threshold {
				out.Rows = append(out.Rows, row)
			}
		}
	}
	return out
}

func ModeCount(table Table) int {
	return DataRows(FilterByMode(table, "bag")) + DataRows(FilterByMode(table, "multi-mode")) +
		DataRows(FilterByMode(table, "sheet-stripping")) + DataRows(FilterByMode(table, "catastrophic"))
}

func DescribeModeCounts(table Table) string {
	counts := ModeRows(table)
	return fmt.Sprintf("%v", counts)
}
