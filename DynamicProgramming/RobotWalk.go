package main

import "fmt"

/*
	假设有排成一行的 N 个位置，记为 1~N ,N>=2
	开始时机器人在其中的 M 位置上(M一定是 1~N 中的一个)
	如果机器人来到 1 位置，那么下一步只能往右来到 2 位置;
    如果机器人来到 N 位置，那么下一步只能往左来到 N-1 位置;
	如果机器人来到中间位置，那么下一步可以往左走或者往右走;
	规定机器人必须走K步，最终能来到 P 位置(P 也是 1-N 中的一个)的方法有多少种
	给定四个参数N、M、K、P，返回方法数.
*/

func main() {
	fmt.Print("暴力递归：")
	fmt.Println(RobotWalk_A(5, 2, 4, 4))
	fmt.Print("记忆化搜索：")
	fmt.Println(RobotWalk_B(5, 2, 4, 4))
	fmt.Print("动态规划：")
	fmt.Println(RobotWalk_C(5, 2, 4, 4))
}

// 暴力递归：
func RobotWalk_A(n, cur, cnt, end int) int {
	if n <= 1 || cur <= 0 || cur > n || cnt < 0 || end <= 0 || end > n {
		return -1
	}
	return RobotWalk_A_Process(n, cur, cnt, end)
}

func RobotWalk_A_Process(n, cur, cnt, end int) int {
	if cnt == 0 {
		if cur == end {
			return 1
		} else {
			return 0
		}
	} else if cur == 1 {
		return RobotWalk_A_Process(n, cur+1, cnt-1, end)
	} else if cur == n {
		return RobotWalk_A_Process(n, cur-1, cnt-1, end)
	} else {
		return RobotWalk_A_Process(n, cur+1, cnt-1, end) + RobotWalk_A_Process(n, cur-1, cnt-1, end)
	}
}

// 记忆化搜索：
func RobotWalk_B(n, cur, cnt, end int) int {
	if n <= 1 || cur <= 0 || cur > n || cnt < 0 || end <= 0 || end > n {
		return -1
	}
	dp := make([][]int, cnt+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= cnt; i++ {
		for j := 0; j <= n; j++ {
			dp[i][j] = -1
		}
	}
	return RobotWalk_B_Process(n, cur, cnt, end, dp)
}

func RobotWalk_B_Process(n, cur, cnt, end int, dp [][]int) int {
	if dp[cnt][cur] != -1 {
		return dp[cnt][cur]
	}
	ans := 0
	if cnt == 0 {
		if cur == end {
			ans = 1
		} else {
			ans = 0
		}
	} else if cur == 1 {
		ans = RobotWalk_B_Process(n, cur+1, cnt-1, end, dp)
	} else if cur == n {
		ans = RobotWalk_B_Process(n, cur-1, cnt-1, end, dp)
	} else {
		ans = RobotWalk_B_Process(n, cur+1, cnt-1, end, dp) + RobotWalk_B_Process(n, cur-1, cnt-1, end, dp)
	}
	dp[cnt][cur] = ans
	return ans
}

// 动态规划：
func RobotWalk_C(n, cur, cnt, end int) int {
	if n <= 1 || cur <= 0 || cur > n || cnt < 0 || end <= 0 || end > n {
		return -1
	}
	dp := make([][]int, cnt+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	dp[0][end] = 1
	// 状态转移方程：
	for i := 1; i <= cnt; i++ {
		dp[i][1] = dp[i-1][2]
		for j := 2; j < n; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j+1]
		}
		dp[i][n] = dp[i-1][n-1]
	}
	return dp[cnt][cur]
}
