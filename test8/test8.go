package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// 创建哈希表，键为元素值，值为元素下标
	numMap := make(map[int]int)

	// 遍历数组
	for i, num := range nums {
		// 计算需要找到的互补值
		complement := target - num

		// 检查互补值是否在哈希表中
		if j, exists := numMap[complement]; exists {
			// 存在则返回两个下标（注意顺序：先找到的在前）
			return []int{j, i}
		}

		// 不存在则将当前元素存入哈希表
		numMap[num] = i
	}

	// 题目保证有解，此处仅为语法兼容
	return nil
}

func main() {
	// 测试用例
	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{2, 7, 11, 15}, 9}, // 输出 [0,1]
		{[]int{3, 2, 4}, 6},      // 输出 [1,2]
		{[]int{3, 3}, 6},         // 输出 [0,1]
		{[]int{5, 7, -2, 3}, 1},  // 输出 [2,3]（-2 + 3 = 1）
	}

	for _, tc := range testCases {
		result := twoSum(tc.nums, tc.target)
		fmt.Printf("nums = %v, target = %d → 结果下标: %v\n", tc.nums, tc.target, result)
	}
}
