func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1]) // 只有两间屋子则只能抢其中一间，因为不能抢隔壁的
	}

	// 递推公式分析：因为不能抢隔壁的屋子，所以有两种情况：
	// 	1. 抢了当前的屋子 dp[i] 就只能抢 dp[i-2] 的屋子，所以总共抢 dp[i-2] + nums[i] 的钱
	//  2. 没抢当前的屋子就是 dp[i-1] 的钱
	// 所以得出递推公式 dp[i] = max(dp[i-1], dp[i-2] + nums[i])
	// 这里用了 max 比较函数，因为题目要求得出最多抢多少钱，所以用到贪心思想（应该算吧?），每次迭代得到最多的钱

	// 初始化数组，通过分析得知 dp[n] 的后续结果是基于 dp[0] 和 dp[1] 递推得到
	// 当抢第 0 间屋子时，dp[0] = nums[0]；抢 0~1 号屋子时，dp[1] = max(nums[0], nums[1]) (因为不能抢隔壁的屋子，二选一)

	// 分析 dp[i] 的含义：抢第 i 间屋子或不抢第 i 间屋子最多能抢到的钱为 dp[i]
	//	 (下标i（包括i）以内的房屋，最多可以偷窃的金额为dp[i])

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1]) // 两间屋子的话，最多可以偷两者选一种

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i - 2] + nums[i], dp[i-1])
	}

	return dp[len(nums)-1]
}