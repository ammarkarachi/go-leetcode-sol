package solution

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestValue(root *TreeNode, target float64) int {
	closest := root.Val
	for root != nil {
		previousDiff := math.Abs(float64(closest) - target)
		currentDiff := math.Abs(float64(root.Val) - target)
		if previousDiff >= currentDiff {
			if previousDiff == currentDiff {
				if root.Val < closest {
					closest = root.Val
				}

			} else {
				closest = root.Val
			}
		}

		if target > float64(root.Val) {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return closest
}
