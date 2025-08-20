package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	// 边界情况：空数组或只有一个区间，直接返回
	if len(intervals) <= 1 {
		return intervals
	}

	// 按区间的起始值排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 初始化结果集，加入第一个区间
	result := [][]int{intervals[0]}

	// 遍历剩余区间，逐步合并
	for i := 1; i < len(intervals); i++ {
		// 取结果集最后一个区间
		last := result[len(result)-1]
		// 当前区间
		current := intervals[i]

		// 若当前区间的起始值 <= 最后区间的结束值，说明重叠，合并
		if current[0] <= last[1] {
			// 新的结束值取两者的最大值
			newEnd := last[1]
			if current[1] > newEnd {
				newEnd = current[1]
			}
			// 更新结果集最后一个区间为合并后的区间
			result[len(result)-1] = []int{last[0], newEnd}
		} else {
			// 不重叠，直接加入结果集
			result = append(result, current)
		}
	}

	return result
}

func main() {
	// 测试用例
	testCases := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, // 输出 [[1,6],[8,10],[15,18]]
		{{1, 4}, {4, 5}},                    // 输出 [[1,5]]
		{{1, 5}, {2, 3}},                    // 输出 [[1,5]]
		{{5, 10}, {1, 3}, {2, 4}},           // 排序后 [[1,3],[2,4],[5,10]] → 合并后 [[1,4],[5,10]]
	}

	for _, intervals := range testCases {
		merged := merge(intervals)
		fmt.Printf("输入: %v → 合并后: %v\n", intervals, merged)
	}
}
