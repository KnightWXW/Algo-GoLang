package main

import "fmt"

/*
	斐波那契数（通常用 F(n) 表示）形成的序列称为 斐波那契数列 。
	该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。
	F(0) = 0，F(1) = 1
	F(n) = F(n - 1) + F(n - 2)，其中 n > 1
	给定n ，请计算 F(n) 。
*/

func main() {
	fmt.Print("暴力递归：")
	fmt.Println(fibonacci_A(7))
	fmt.Print("记忆化搜索：")
	hashMap := map[int]int{}
	fmt.Println(fibonacci_B(7, hashMap))
	fmt.Print("动态规划：")
	fmt.Println(fibonacci_C(7))
	fmt.Print("动态规划(空间优化)：")
	fmt.Println(fibonacci_D(7))
}

// 暴力递归：
func fibonacci_A(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	return fibonacci_A(n-1) + fibonacci_A(n-2)
}

// 记忆化搜索：
func fibonacci_B(n int, hashMap map[int]int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	hashMap[0], hashMap[1] = 1, 1
	if v, ok := hashMap[n]; ok {
		return v
	} else {
		return fibonacci_B(n-1, hashMap) + fibonacci_B(n-2, hashMap)
	}
}

// 动态规划：
func fibonacci_C(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 动态规划(空间优化)：
func fibonacci_D(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	a, b := 1, 1
	for i := 2; i <= n; i++ {
		c := a + b
		a, b = b, c
	}
	return b
}
