package graph

func DetermineSources(numNodes int, m []int) (sources []int) {
	sources = make([]int, 0, numNodes)
	touched := map[int]struct{}{}

loop:
	for y := 0; y < numNodes; y++ {
		// make sure node in the graph is a source,
		// i.e. it has no incoming edges
		for x := 0; x < numNodes; x++ {
			if m[x*numNodes+y] != 0 {
				touched[y] = struct{}{}
				continue loop
			}
		}
	}

	for i := 0; i < numNodes; i++ {
		if _, ok := touched[i]; !ok {
			// found winner
			sources = append(sources, i)
		}
	}

	return
}
