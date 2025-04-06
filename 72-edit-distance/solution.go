package solution

func minDistance(word1 string, word2 string) int {
	memo := make([][]int, len(word1)+1)
	for i, _ := range memo {
		memo[i] = make([]int, len(word2)+1)
		for j := range len(word2) + 1 {
			memo[i][j] = -1
		}

	}

	var getMindistance func(w1, w2 string, id1, id2 int) int
	getMindistance = func(w1, w2 string, id1, id2 int) int {
		if id1 == 0 {
			return id2
		}
		if id2 == 0 {
			return id1
		}
		if memo[id1][id2] != -1 {
			return memo[id1][id2]
		}
		var minD int
		if w1[id1-1] == w2[id2-1] {
			minD = getMindistance(
				w1,
				w2,
				id1-1,
				id2-1,
			)
		} else {
			deleteIdx := getMindistance(w1, w2, id1, id2-1)
			insertIdx := getMindistance(w1, w2, id1-1, id2)
			replace := getMindistance(w1, w2, id1-1, id2-1)

			minD = min(deleteIdx, min(insertIdx, replace))
		}
		memo[id1][id2] = minD
		return minD
	}
	return getMindistance(word1, word2, len(word1), len(word2))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
