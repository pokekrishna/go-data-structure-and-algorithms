package selection_test

import (
	"github.com/pokekrishna/dsa/algorithms/sort/selection"
	"testing"
)


func TestSort(t *testing.T) {
	var sourceList = []int{
		1,9,2,4,10,0,
	}

	expected := []int{
		0,1,2,4,9,10,
	}

	t.Logf("Source list %v, len %v\n", sourceList, len(sourceList))
	selection.Sort(&sourceList)
	t.Logf("Sorted Source  list %v, len %v\n", sourceList, len(sourceList))
	for i := range expected {
		if expected[i] != sourceList[i] {
			t.Errorf("expected value '%d' at index %d but found '%d'",
				expected[i], i, sourceList[i])
		}
	}


	// TODO: try with repeated values
}

func TestSwapValues(t *testing.T) {
	var sourceList = []int{
		1,9,2,4,10,0,
	}

	expected := []int{
		1,9,10,4,2,0,
	}

	selection.SwapValues(&sourceList, 2, 4)

	for i := range expected {
		if expected[i] != sourceList[i] {
			t.Errorf("expected value '%d' at index %d but found '%d'",
				expected[i], i, sourceList[i])
		}
	}

	t.Log("Swapped List.", sourceList)


}