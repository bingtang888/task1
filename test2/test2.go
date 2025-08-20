/**
  给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
  回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
  例如，121 是回文，而 123 不是。
*/

package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversedHalf := 0

	for x > reversedHalf {
		reversedHalf = reversedHalf*10 + x%10
		x /= 10
	}

	return x == reversedHalf || x == reversedHalf/10
}

func main() {
	// 测试用例
	testCases := []int{121, -121, 120, 0, 131, 12321}
	for _, num := range testCases {
		result := isPalindrome(num)
		fmt.Printf("%d 是回文数吗？ %v\n", num, result)
	}
}
