/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
<<<<<<< HEAD
package leetCode

=======
>>>>>>> b3d0e63b69b69f1a19d98744e68907bc93d2ad01
func maxArea(height []int) int {
	var (
		maxCapacity = 0
		leftIndex   = 0
		rightIndex  = len(height) - 1
	)

	for leftIndex < rightIndex {
		leftHegiht, rightHeight := height[leftIndex], height[rightIndex]

		minHeight := leftHegiht
		if minHeight > rightHeight {
			minHeight = rightHeight
		}

		capacity := minHeight * (rightIndex - leftIndex)
		if capacity > maxCapacity {
			maxCapacity = capacity
		}

		if minHeight == leftHegiht {
			leftIndex++
		} else {
			rightIndex--
		}
	}

	return maxCapacity
}

// @lc code=end
