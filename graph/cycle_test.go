package graph

import "testing"

type cycleTest struct {
	m        []int
	size     int
	hasCycle bool
}

var cycleTests = []cycleTest{
	{
		[]int{
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
		}, 4, false,
	},
	{
		[]int{
			0, 0, 0, 0,
			1, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
		}, 4, false,
	},
	{
		[]int{
			0, 1, 0, 0,
			1, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0,
		}, 4, true,
	},
	{
		[]int{
			0, 0, 0, 0,
			1, 0, 1, 1,
			1, 0, 0, 1,
			1, 0, 0, 0,
		}, 4, false,
	},
}

func TestCycleDetect(t *testing.T) {
	for _, test := range cycleTests {
		if DetectCycle(test.size, test.m) != test.hasCycle {
			t.Errorf("mismatch")
		}
	}
}
