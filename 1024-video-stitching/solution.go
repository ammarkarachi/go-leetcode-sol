package solution

func videoStitching(clips [][]int, time int) int {
	videoMap := make(map[int]int)

	for _, clip := range clips {
		videoMap[clip[0]] = max(videoMap[clip[0]], clip[1])
	}
	start := 0
	end, startExists := videoMap[0]
	if !startExists {
		return -1
	}
	i := end

	count := 0
	for i <= time {
		nextEnd, nextClipExists := videoMap[i]
		for !nextClipExists && i > start {
			i--
			nextEnd, nextClipExists = videoMap[i]
		}

		if !nextClipExists {
			return -1
		} else {
			count++
			i = nextEnd
		}

	}
	return count

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
