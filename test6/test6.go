package main

/**
  给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
  考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：

  更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
  返回 k 。
  判题标准:

  示例 1：

  输入：nums = [1,1,2]
  输出：2, nums = [1,2,_]
  解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。

  示例 2：
  输入：nums = [0,0,1,1,1,2,2,3,3,4]
  输出：5, nums = [0,1,2,3,4]
  解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。

*/

import "fmt"

func removeDuplicates(nums []int) int {
	// 边界情况：空数组直接返回0
	if len(nums) == 0 {
		return 0
	}

	// 慢指针初始指向第一个元素
	slow := 0

	// 快指针从第二个元素开始遍历
	for fast := 1; fast < len(nums); fast++ {
		// 当快指针元素与慢指针元素不同时，更新慢指针并复制元素
		if nums[fast] != nums[slow] {
			slow++                  // 慢指针后移一位，准备存储新元素
			nums[slow] = nums[fast] // 将新元素复制到慢指针位置
		}
	}

	// 慢指针索引+1即为唯一元素的个数
	return slow + 1
}

func main() {
	// 测试用例
	testCases := [][]int{
		{1, 1, 2},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		{5},          // 单个元素
		{2, 2, 2, 2}, // 所有元素重复
	}

	for _, nums := range testCases {
		// 复制原始数组用于输出对比（避免原数组被修改影响显示）
		original := make([]int, len(nums))
		copy(original, nums)

		k := removeDuplicates(nums)
		// 输出结果：原始数组、新长度、处理后的前k个元素
		fmt.Printf("输入: %v → 新长度: %d, 前%d个元素: %v\n", original, k, k, nums[:k])
	}
}
