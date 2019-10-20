/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
<<<<<<< HEAD
package leetCode

=======
>>>>>>> b3d0e63b69b69f1a19d98744e68907bc93d2ad01
func twoSum(nums []int, target int) []int {
	ret := make([]int, 0, 2)

	if len(nums) < 2 {
		return ret
	}

	maps := make(map[int]int, len(nums))

	for index, value := range nums {
		needValue := target - value

		if _, exist := maps[needValue]; exist {
			ret = append(ret, maps[needValue])
			ret = append(ret, index)
			break
		}

		maps[value] = index
	}

	return ret
}

// @lc code=end
