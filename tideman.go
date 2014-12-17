// Warning, insufficiently tested alpha software.
package condorcet

import "sort"
import "github.com/hlandau/condorcet/graph"

// Matrix shall be of dimensionality [numOptions][numOptions]int.
// For each z = matrix[x][y], z is the number of votes x has over y.
//
//   matrix[row][column]
//
//      0  1  2  3
//   0  -  d  g  j
//   1  a  -  h  k
//   2  b  e  -  l
//   3  c  f  i  -
//
//   a: 1>0
//   b: 2>0
//   c: 3>0
//   d: 0>1
//   e: 2>1
//   f: 3>1
//   g: 0>2
//   h: 1>2
//   i: 3>2
//   j: 0>3
//   k: 1>3
//
// To read the matrix:
//   Let f(x,y) represent the number of votes for x>y.
//   f(x,y) = M[x][y]   } where x ≠ y
//
// The x=y cells are not used and may have any values.
//
func TidemanWinnerFromMatrix(numOptions int, voteMatrix []int) (winners []int) {
	victories := determineVictories(numOptions, voteMatrix)

	sortVictories(victories, voteMatrix)

	// This matrix represents a directed graph.
	// Each row and column represents a node.
	// Each column is set to 1 iff the row node has an outgoing edge to the column node.
	// i.e. row x, column y = 1 means that node x has an outgoing edge to node y.
	m2 := make([]int, numOptions*numOptions)
	for _, v := range victories {
		graph.SetIfDoesNotCreateCycle(numOptions, m2, v.winningOption, v.losingOption)
	}

	winners = graph.DetermineSources(numOptions, m2)
	return
}

type victory struct {
	winningOption, losingOption int
	winningVotes, losingVotes   int
}

func determineVictories(numOptions int, voteMatrix []int) (victories []victory) {
	victories = make([]victory, 0, numOptions)

	for x := 0; x < numOptions; x++ {
		for y := 0; y < numOptions; y++ {
			if x == y {
				continue
			}

			v1 := voteMatrix[x*numOptions+y]
			v2 := voteMatrix[y*numOptions+x]
			vict := victory{}
			d := v1 - v2
			if d > 0 {
				vict.winningOption = x
				vict.losingOption = y
				vict.winningVotes = v1
				vict.losingVotes = v2
				victories = append(victories, vict)
			}
		}
	}

	return
}

func sortVictories(victories []victory, matrix []int) {
	sorter := &sorter{
		lenFunc: func() int { return len(victories) },
		lessFunc: func(i, j int) bool {
			Vxy, Vzw := &victories[i], &victories[j]
			dVxy := Vxy.winningVotes // - Vxy.losingVotes
			dVzw := Vzw.winningVotes // - Vzw.losingVotes
			dVyx := Vxy.losingVotes  // - Vxy.winningVotes
			dVwz := Vzw.losingVotes  // - Vzw.winningVotes

			if dVxy > dVzw {
				return true
			}

			if dVxy == dVzw && Vxy.winningOption == Vzw.winningOption && dVyx < dVwz {
				return true
			}

			return false
		},
		swapFunc: func(i, j int) {
			victories[i], victories[j] = victories[j], victories[i]
		},
	}

	sort.Sort(sorter)
}
