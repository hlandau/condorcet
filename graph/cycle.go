package graph

func DetectCycle(numNodes int, m []int) bool {
	for x := 0; x < numNodes; x++ {
		seen := map[int]struct{}{}
		seen[x] = struct{}{}
		if detectCycle(m, numNodes, x, seen) {
			return true
		}
	}
	return false
}

func detectCycle(m []int, numNodes, startNode int, seen map[int]struct{}) bool {
	for i := 0; i < numNodes; i++ {
		idx := startNode*numNodes + i
		if m[idx] != 0 {
			if _, ok := seen[i]; ok {
				return true
			}
			seen2 := map[int]struct{}{}
			for k, v := range seen {
				seen2[k] = v
			}
			seen2[i] = struct{}{}
			if detectCycle(m, numNodes, i, seen2) {
				return true
			}
		}
	}

	return false
}

func SetIfDoesNotCreateCycle(numNodes int, m []int, from, to int) {
	idx := from*numNodes + to
	if m[idx] == 0 {
		m[idx] = 1
		if DetectCycle(numNodes, m) {
			m[idx] = 0
		}
	}
}
