package table

import (
	"testing"

	"weber-br/internal/drop"
)

func TestFromResults(t *testing.T) {
	result := drop.Result{U: 1, D: 0.002, We: 1, Oh: 0.01, WeCrit: 12, Mode: "stable", Breakup: false}
	table := FromResults([]drop.Result{result})
	if DataRows(table) != 1 || Columns(table) != 7 {
		t.Fatalf("table=%+v", table)
	}
}

func TestFilterByMode(t *testing.T) {
	table := Table{Header: []string{"U", "d", "We", "Oh", "We_crit", "mode", "breakup"},
		Rows: [][]string{{"1", "2", "3", "4", "5", "stable", "false"}, {"2", "3", "4", "5", "6", "bag", "true"}}}
	out := FilterByMode(table, "bag")
	if DataRows(out) != 1 {
		t.Fatalf("rows=%d", DataRows(out))
	}
}

func TestModeCount(t *testing.T) {
	table := Table{Header: []string{"U", "d", "We", "Oh", "We_crit", "mode", "breakup"},
		Rows: [][]string{{"1", "2", "3", "4", "5", "bag", "true"}, {"2", "3", "4", "5", "6", "bag", "true"}}}
	if ModeCount(table) != 2 {
		t.Fatalf("count=%d", ModeCount(table))
	}
}

func TestSummaryAndValidate(t *testing.T) {
	table := Table{Header: []string{"a"}, Rows: [][]string{{"1"}}}
	if Summary(table) == "" {
		t.Fatal("empty summary")
	}
	if err := Validate(table); err != nil {
		t.Fatal(err)
	}
}
