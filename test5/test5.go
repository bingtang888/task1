package main

/**
    给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
    将大整数加 1，并返回结果的数字数组。

示例 1：
输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
加 1 后得到 123 + 1 = 124。
因此，结果应该是 [1,2,4]。

示例 2：
输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。
加 1 后得到 4321 + 1 = 4322。
因此，结果应该是 [4,3,2,2]。

示例 3：
输入：digits = [9]
输出：[1,0]
解释：输入数组表示数字 9。
加 1 得到了 9 + 1 = 10。
因此，结果应该是 [1,0]。
*/

import (
	"fmt"
)

func plusOne(digits []int) []int {
	// 函数逻辑不变
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] < 10 {
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}

func main() {
	testCases := [][]int{
		{1, 2, 3},
		{4, 3, 2, 1},
		{9},
		{9, 9, 9},
		{1, 9, 9},
	}

	for _, original := range testCases {
		// 复制原始切片，避免修改原数据
		digitsCopy := make([]int, len(original))
		copy(digitsCopy, original)

		result := plusOne(digitsCopy)
		fmt.Printf("输入: %v, 加1后: %v\n", original, result)
	}
}
