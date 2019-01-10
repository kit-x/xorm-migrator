package migrator

import (
	"sort"
	"testing"
)

func TestSortStageByID(t *testing.T) {
	stages := []Stage{
		{
			ID: "v002",
		},
		{
			ID: "v001",
		},
	}

	sort.Sort(ByID(stages))

	if stages[0].ID != "v001" {
		t.Error("sorted failed stages[0] should be v001")
	}

	if stages[1].ID != "v002" {
		t.Error("sorted failed stages[1] should be v002")
	}
}
