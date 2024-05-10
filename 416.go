// 分割等和子集 动态规划
// 时间复杂度O(n^2) 空间复杂度O(n)
func canPartition(nums []int) bool {
    length := len(nums)

    sum := 0
    for _, v := range nums {
        sum += v
    }

    if sum % 2 != 0 { // 如果不能被 2 整除，则表明不可能平分为两个子集
        return false
    }

    target := sum / 2

    dp := make([]int, target + 1) // dp 数组长度 = 背包容量 + 1（因为从 0 开始所以 +1）

    // 在这个题目中，物品的 价值 = 重量 = 数字
    // 背包的容量为 target
    values := nums
    weight := nums

    // 初始化 dp 数组
    dp[0] = 0 // 当背包容量为 0 时不能容纳任何物品所以最大价值为 0
    for i := 1; i < target; i++ {
        dp[i] = 0 // 递推公式有 max 函数，所以递推初始值需从非负整数的最小值开始
    }

    // dp[j] 的含义是：最大重量（容量）为 j 的背包放入物品的最大价值为 dp[j]

    // 遍历顺序
    for i := 0; i < length; i++ { // 物品编号（for i j 循环嵌套不能交换）
        for j := target; j >= nums[i]; j-- { // 背包容量（当背包容量小于物品重量时停止迭代）
            // 为什么要从后面往前遍历，因为一维状态压缩后，从前往后会出现背包重复放入物品的问题

            // 状态转移方程
            // dp[j] 相当于二维版本中的 dp[i - 1][j]，这里进行了状态压缩，利用了上一个值（in-place）
            dp[j] = max(dp[j], dp[j - weight[i]] + values[i])
        }
    }

    // 把数字放入容量为 target 背包后最大价值是否与 target 相等
    // 因为 重量 = 价值，所以如果 dp[target] < target [ (价值 => 重量) < (容量) ] 则不能装满背包
    // 则返回 false，不能分割
    return dp[target] == target
}