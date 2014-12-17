package condorcet_test

import "reflect"
import "testing"
import "github.com/hlandau/condorcet"

type test struct {
	numOptions int
	winners    []int
	matrix     []int
}

var tests = []test{
	{4, []int{1}, []int{
		00, 42, 42, 42,
		58, 00, 68, 68,
		58, 32, 00, 83,
		58, 32, 17, 00,
	}},
}

func TestTideman(t *testing.T) {
	for i, tst := range tests {
		winners := condorcet.TidemanWinnerFromMatrix(tst.numOptions, tst.matrix)
		if !reflect.DeepEqual(winners, tst.winners) {
			t.Errorf("wrong winners on item %d: got %v, expected %v", i, winners, tst.winners)
		}
	}
}

type strtest struct {
	candidates []string
	votes      [][][]string
	winners    []string
}

var strtests = []strtest{
	{[]string{"A", "B", "C", "D"}, [][][]string{
		{{"B"}, {"A"}},
	}, []string{"B"}},

	{[]string{"A", "B", "C", "D"}, [][][]string{
		{{"B", "C"}, {"A"}},
	}, []string{"B", "C"}},

	{[]string{"A", "B", "C", "D"}, [][][]string{}, []string{"A", "B", "C", "D"}},
}

func TestTidemanStr(t *testing.T) {
	for i, tst := range strtests {
		winners, err := condorcet.TidemanWinner(tst.candidates, tst.votes)
		if err != nil {
			t.Errorf("got error: %v", err)
			continue
		}

		if !reflect.DeepEqual(winners, tst.winners) {
			t.Errorf("wrong winners on item %d: got %v, expected %v", i, winners, tst.winners)
		}
	}
}
