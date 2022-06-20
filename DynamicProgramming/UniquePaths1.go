package main

import (
	"fmt"
	"strconv"
)

/*
	一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
	机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
	问总共有多少条不同的路径？

	输入：m = 3, n = 7
	输出：28

	输入：m = 3, n = 2
	输出：3
	解释：  从左上角开始，总共有 3 条路径可以到达右下角。
			1. 向右 -> 向下 -> 向下
			2. 向下 -> 向下 -> 向右
			3. 向下 -> 向右 -> 向下

	输入：m = 7, n = 3
	输出：28

	输入：m = 3, n = 3
	输出：6
*/

func main() {
	m, n := 7, 5
	fmt.Print("暴力递归：")
	fmt.Println(uniquePaths1_A(m, n)) // 210
	fmt.Print("记忆化搜索：")
	fmt.Println(uniquePaths1_B(m, n)) // 210
	fmt.Print("动态规划：")
	fmt.Println(uniquePaths1_C(m, n)) // 210
}

// 暴力递归：
func uniquePaths1_A(m, n int) int {
	return dfsUniquePaths1_A(0, 0, m, n)
}

func dfsUniquePaths1_A(i, j, m, n int) int {
	if i < 0 || i > m || j < 0 || j > n {
		return 0
	}
	if i == m-1 && j == n-1 {
		return 1
	}
	return dfsUniquePaths1_A(i+1, j, m, n) + dfsUniquePaths1_A(i, j+1, m, n)
}

// 记忆化搜索：
func uniquePaths1_B(m, n int) int {
	hashMap := map[string]int{}
	return dfsUniquePaths1_B(0, 0, m, n, hashMap)
}

func dfsUniquePaths1_B(i, j, m, n int, hashMap map[string]int) int {
	v := strconv.Itoa(i) + strconv.Itoa(j)
	if i < 0 || i > m || j < 0 || j > n {
		hashMap[v] = 0
		return 0
	}
	if i == m-1 && j == n-1 {
		hashMap[v] = 1
		return 1
	}
	v1, v2 := strconv.Itoa(i+1)+strconv.Itoa(j), strconv.Itoa(i)+strconv.Itoa(j+1)
	p1, p2 := 0, 0
	if hashMap[v1] > 0 {
		p1 = hashMap[v1]
	} else {
		p1 = dfsUniquePaths1_B(i+1, j, m, n, hashMap)
	}
	if hashMap[v2] > 0 {
		p2 = hashMap[v2]
	} else {
		p2 = dfsUniquePaths1_B(i, j+1, m, n, hashMap)
	}
	hashMap[v] = p1 + p2
	return hashMap[v]
}

// 动态规划：
func uniquePaths1_C(m, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// 初始化：
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
