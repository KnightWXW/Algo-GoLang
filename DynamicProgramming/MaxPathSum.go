package main

import (
	"fmt"
	"strconv"
)

/*
	有一个 m × n 的矩阵，现要从左上角走到右下角，并且方向只能是向下或者向右，
	现规定一条路径的权值为走此路径所经过的值的和。给定一个矩阵，请找出权值最大的一条路径。
	输入：
		2 5 6 4
		3 9 4 3
		7 9 1 7
	输出： 33.  所找到的路径为2->5->9->9->1->7,最大路径和为33。
*/

func main() {
	arr := [][]int{{2, 5, 6, 4}, {3, 9, 4, 3}, {7, 9, 1, 7}}
	fmt.Print("暴力递归：")
	fmt.Println(maxPathSum_A(arr)) // 33
	fmt.Print("记忆化搜索：")
	fmt.Println(maxPathSum_B(arr)) // 33
	fmt.Print("动态规划：")
	fmt.Println(maxPathSum_C(arr)) // 33
	fmt.Print("动态规划(空间优化)：")
	fmt.Println(maxPathSum_D(arr)) // 33
}

// 暴力递归：
func maxPathSum_A(arr [][]int) int {
	return dfsMaxPathSum_A(0, 0, arr)
}

func dfsMaxPathSum_A(i, j int, arr [][]int) int {
	m, n := len(arr), len(arr[0])
	if i < 0 || i > m || j < 0 || j > n {
		return 0
	}
	if i == m-1 && j == n-1 {
		return arr[i][j]
	}
	if i == m-1 {
		return dfsMaxPathSum_A(i, j+1, arr) + arr[i][j]
	}
	if j == n-1 {
		return dfsMaxPathSum_A(i+1, j, arr) + arr[i][j]
	}
	return max(dfsMaxPathSum_A(i+1, j, arr), dfsMaxPathSum_A(i, j+1, arr)) + arr[i][j]
}

// 记忆化搜索：
func maxPathSum_B(arr [][]int) int {
	hashMap := map[string]int{}
	return dfsMaxPathSum_B(0, 0, arr, hashMap)
}

func dfsMaxPathSum_B(i, j int, arr [][]int, hashMap map[string]int) int {
	m, n := len(arr), len(arr[0])
	v := strconv.Itoa(i) + strconv.Itoa(j)
	if i < 0 || i > m || j < 0 || j > n {
		return 0
	}
	if i == m-1 && j == n-1 {
		hashMap[v] = arr[i][j]
		return hashMap[v]
	}
	if i == m-1 {
		if hashMap[v] > 0 {
			return hashMap[v]
		} else {
			return dfsMaxPathSum_B(i, j+1, arr, hashMap) + arr[i][j]
		}
	}
	if j == n-1 {
		if hashMap[v] > 0 {
			return hashMap[v]
		} else {
			return dfsMaxPathSum_B(i+1, j, arr, hashMap) + arr[i][j]
		}
	}
	hashMap[v] = max(dfsMaxPathSum_B(i+1, j, arr, hashMap), dfsMaxPathSum_B(i, j+1, arr, hashMap)) + arr[i][j]
	return hashMap[v]
}

// 动态规划：
// Time: O(m * n)
// Space: O(m * n)
func maxPathSum_C(arr [][]int) int {
	m, n := len(arr), len(arr[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + arr[i-1][j-1]
		}
	}
	return dp[m][n]
}

// 动态规划(空间优化)：
// Time: O(m * n)
// Space: O(n)
func maxPathSum_D(arr [][]int) int {
	m, n := len(arr), len(arr[0])
	dp := make([]int, n+1)

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[j] = max(dp[j-1], dp[j]) + arr[i-1][j-1]
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
