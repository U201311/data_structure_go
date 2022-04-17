package dynamicprogramming

/*动态规划
leetcode 300  最长递增子序列
动态规划解法
时间复杂度n^2
*/
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	temp := -1
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j <= i-1; j++ {
			dp[i] = max(dp[j]+1, dp[i])
		}
		temp = max(temp, dp[i])
	}
	return temp
}

/*时间复杂度 nlogn解法
step1:建立递升序列数组，
step2:不断查找当前元素在递升序列中的位置（二分查找）
step3:返回数组长度
*/
func lengthOfLIS2(nums []int) int {
	res := []int{}
	n := len(nums)
	if n == 1 {
		return 1
	}
	res = append(res, nums[0])
	for i := 0; i < n; i++ {
		left := 0
		right := len(res) - 1
		for left <= right {
			mid := (left + right) / 2
			if res[mid] >= nums[i] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		if left == len(res) {
			res = append(res, nums[i])
		} else {
			res[left] = nums[i]
		}
	}
	return len(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
leetcode 322
*/

func coinChange(coins []int, amount int) int {
	n := len(coins)
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for j := 0; j < n; j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
