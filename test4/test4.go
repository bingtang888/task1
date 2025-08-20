/**
  编写一个函数来查找字符串数组中的最长公共前缀。
  如果不存在公共前缀，返回空字符串 ""。

  示例 1：
  输入：strs = ["flower","flow","flight"]
  输出："fl"

  示例 2：
  输入：strs = ["dog","racecar","car"]
  输出：""
*/

package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 以第一个字符串为基准，遍历每个字符位置
	for i := 0; i < len(strs[0]); i++ {
		// 取第一个字符串当前位置的字符作为参照
		char := strs[0][i]

		//检查其他所有字符串的对应位置
		for j := 1; j < len(strs); j++ {
			//若其他字符串长度不足， 或字符不匹配，返回当前前缀
			if i >= len(strs[j]) || strs[j][i] != char {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func main() {
	// 测试用例
	testCases := [][]string{
		{"flower", "flow", "flight"},    // 输出 "fl"
		{"dog", "racecar", "car"},       // 输出 ""
		{"apple", "app", "application"}, // 输出 "app"
		{"a"},                           // 输出 "a"
		{"", "b"},                       // 输出 ""
	}

	for _, strs := range testCases {
		fmt.Printf("输入: %v, 最长公共前缀: %q\n", strs, longestCommonPrefix(strs))
	}
}
