package main

import (
	"fmt"
	"strconv"
)

/*
	一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
	机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。
	现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
	网格中的障碍物和空位置分别用 1 和 0 来表示。

	示例 1：
	输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
	输出：2
	解释：3x3 网格的正中间有一个障碍物。
			从左上角到右下角一共有 2 条不同的路径：
			1. 向右 -> 向右 -> 向下 -> 向下
			2. 向下 -> 向下 -> 向右 -> 向右

	示例 2：
	输入：obstacleGrid = [[0,1],[0,0]]
	输出：1。
*/

func main() {
	arr := [][]int{{0, 0, 0, 1}, {0, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}}
	fmt.Print("暴力递归：")
	fmt.Println(uniquePaths2_A(arr)) // 10
	fmt.Print("记忆化搜索：")
	fmt.Println(uniquePaths2_B(arr)) // 10
	fmt.Print("动态规划：")
	fmt.Println(uniquePaths2_C(arr)) // 10
	fmt.Print("动态规划(空间优化)：")
	fmt.Println(uniquePaths2_D(arr)) // 10
}

// 暴力递归：
func uniquePaths2_A(arr [][]int) int {
	return dfsUniquePaths2_A(0, 0, arr)
}

func dfsUniquePaths2_A(i, j int, arr [][]int) int {
	m, n := len(arr), len(arr[0])
	if i < 0 || i > m || j < 0 || j > n {
		return 0
	}
	if arr[i][j] == 1 {
		return 0
	}
	if i == m-1 && j == n-1 {
		return 1
	}
	if i == m-1 {
		return dfsUniquePaths2_A(i, j+1, arr)
	}
	if j == n-1 {
		return dfsUniquePaths2_A(i+1, j, arr)
	}
	return dfsUniquePaths2_A(i+1, j, arr) + dfsUniquePaths2_A(i, j+1, arr)
}

// 记忆化搜索：
func uniquePaths2_B(arr [][]int) int {
	hashMap := map[string]int{}
	return dfsUniquePaths2_B(0, 0, arr, hashMap)
}

func dfsUniquePaths2_B(i, j int, arr [][]int, hashMap map[string]int) int {
	m, n := len(arr), len(arr[0])
	v := strconv.Itoa(i) + strconv.Itoa(j)
	if i < 0 || i > m || j < 0 || j > n {
		return 0
	}
	if arr[i][j] == 1 {
		hashMap[v] = 0
		return 0
	}
	if i == m-1 && j == n-1 {
		hashMap[v] = 1
		return 1
	}
	if i == m-1 {
		if hashMap[v] > 0 {
			return hashMap[v]
		} else {
			return dfsUniquePaths2_B(i, j+1, arr, hashMap)
		}
	}
	if j == n-1 {
		if hashMap[v] > 0 {
			return hashMap[v]
		} else {
			return dfsUniquePaths2_B(i+1, j, arr, hashMap)
		}
	}
	hashMap[v] = dfsUniquePaths2_B(i+1, j, arr, hashMap) + dfsUniquePaths2_B(i, j+1, arr, hashMap)
	return hashMap[v]
}

// 动态规划：
// Time: O(m * n)
// Space: O(m * n)
func uniquePaths2_C(arr [][]int) int {
	m, n := len(arr), len(arr[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m && arr[i][0] != 1; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n && arr[0][j] != 1; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if arr[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// 动态规划(空间优化)：
// Time: O(m * n)
// Space: O(n)
func uniquePaths2_D(arr [][]int) int {
	m, n := len(arr), len(arr[0])
	dp := make([]int, n)

	for i := 0; i < m; i++ {
		if arr[i][0] == 0 {
			dp[0] = 1
		}
		for j := 1; j < n; j++ {
			if arr[i][j] == 0 {
				dp[j] = dp[j-1] + dp[j]
			} else {
				dp[j] = 0
			}
		}
	}
	return dp[n-1]
}
