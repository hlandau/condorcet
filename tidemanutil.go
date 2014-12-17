package condorcet

// Votes: Slice of votes.
// Vote: A>B>C=D>E>F becomes [][]string{ {"A"}, {"B"}, {"C","D"}, {"E"}, {"F"}, }
func TidemanWinner(candidates []string, votes [][][]string) (winners []string, err error) {
	counter := 0
	cmap := map[string]int{}
	for _, c := range candidates {
		cmap[c] = counter
		counter++
	}

	numCandidates := len(candidates)
	matrix := make([]int, numCandidates*numCandidates)
	var touched map[int]struct{}

	oneForI := func(x, y int) {
		_, ok1 := touched[x]
		_, ok2 := touched[y]
		if ok1 && ok2 {
			return
		}

		touched[x] = struct{}{}
		touched[y] = struct{}{}
		matrix[x*numCandidates+y]++
	}

	oneFor := func(x, y string) { // X > Y
		xnum := cmap[x]
		ynum := cmap[y]

		oneForI(xnum, ynum)
	}

	oneForA := func(x, y []string) { // X1=X2 > Y1=Y2
		for _, x1 := range x {
			for _, y1 := range y {
				oneFor(x1, y1)
			}
		}
	}

	for _, v := range votes {
		touched = map[int]struct{}{}

		for x := 0; x < len(v); x++ {
			for y := x + 1; y < len(v); y++ {
				oneForA(v[x], v[y])
			}
		}

		// add implicit items
		for x := 0; x < len(v); x++ {
			for i := 0; i < numCandidates; i++ {
				if i == x {
					continue
				}
				if _, ok := touched[i]; !ok {
					oneForI(x, i)
				}
			}
		}
	}

	iwinners := TidemanWinnerFromMatrix(numCandidates, matrix)
	for _, iw := range iwinners {
		ws := candidates[iw]
		winners = append(winners, ws)
	}

	return
}
