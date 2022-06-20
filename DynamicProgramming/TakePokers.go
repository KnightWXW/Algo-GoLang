package main

import "fmt"

/*
	给定一个 整型数组arr，代表数值不同的纸牌排成一条线
	玩家A 和 玩家B 依次拿走 每张纸牌, 规定玩家A先拿，玩家B后拿,
	但是每个玩家每次只能拿走 最左 或 最右 的纸牌,
	玩家A 和 玩家B 都 聪明绝顶,
	请返回最后获胜者的分数。
*/

func main() {
	arr := []int{5, 7, 4, 5, 8, 1, 6, 0, 3, 4, 6, 1, 7}
	fmt.Print("暴力递归：")
	fmt.Println(takePokers_A(arr))
	fmt.Print("记忆化搜索：")
	fmt.Println(takePokers_B(arr))
	fmt.Print("动态规划：")
	fmt.Println(takePokers_C(arr))
}

// 暴力递归：
func takePokers_A(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	a := dfsTakePokersFirstHand_A(0, len(arr)-1, arr)
	b := dfsTakePokersBackHand_A(0, len(arr)-1, arr)
	return max(a, b)
}

// 先手 取得的 最好成绩：
func dfsTakePokersFirstHand_A(left, right int, arr []int) int {
	if left >= right {
		return arr[left]
	}
	// 自己是先手时，比较 拿走 arr[left] 与 arr[right] 的 分数大小，取较大值：
	p1 := arr[left] + dfsTakePokersBackHand_A(left+1, right, arr)
	p2 := arr[right] + dfsTakePokersBackHand_A(left, right-1, arr)
	return max(p1, p2)
}

func dfsTakePokersBackHand_A(left, right int, arr []int) int {
	if left >= right {
		return 0
	}
	// 自己是后手时，比较 对方 取走 arr[left] 与 arr[right] 的 分数大小，取较小值：
	// 对方 先手 取走 arr[left] 等于 自己 先手 dfsTakePokersFirstHand(left+1, right, arr)
	// 对方 先手 取走 arr[right] 等于 自己 先手 dfsTakePokersFirstHand(left, right-1, arr)
	p1 := dfsTakePokersFirstHand_A(left+1, right, arr)
	p2 := dfsTakePokersFirstHand_A(left, right-1, arr)
	return min(p1, p2)
}

// 记忆化搜索：
func takePokers_B(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	n := len(arr)

	arrFirstHand := make([][]int, n)
	arrBackHand := make([][]int, n)
	for i := range arrFirstHand {
		arrFirstHand[i] = make([]int, n)
	}
	for i := range arrBackHand {
		arrBackHand[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			arrFirstHand[i][j] = -1
			arrBackHand[i][j] = -1
		}
	}
	a := dfsTakePokersFirstHand_B(0, len(arr)-1, arr, arrFirstHand, arrBackHand)
	b := dfsTakePokersBackHand_B(0, len(arr)-1, arr, arrFirstHand, arrBackHand)
	return max(a, b)
}

// 先手 取得的 最好成绩：
func dfsTakePokersFirstHand_B(left, right int, arr []int, arrFirstHand, arrBackHand [][]int) int {
	if arrFirstHand[left][right] != -1 {
		return arrFirstHand[left][right]
	}
	point := 0
	if left >= right {
		point = arr[left]
	} else {
		// 自己是 先手时，比较 取走 arr[left] 与 arr[right] 的 分数大小，取较大值：
		p1 := arr[left] + dfsTakePokersBackHand_B(left+1, right, arr, arrFirstHand, arrBackHand)
		p2 := arr[right] + dfsTakePokersBackHand_B(left, right-1, arr, arrFirstHand, arrBackHand)
		point = max(p1, p2)
	}
	arrBackHand[left][right] = point
	return point
}

func dfsTakePokersBackHand_B(left, right int, arr []int, arrFirstHand, arrBackHand [][]int) int {
	if arrFirstHand[left][right] != -1 {
		return arrFirstHand[left][right]
	}
	point := 0
	if left >= right {
		point = 0
	} else {
		// 自己是后手时，比较 对方 取走 arr[left] 与 arr[right] 的 分数大小，取较小值：
		// 对方 先手 取走 arr[left] 等于 自己 先手 dfsTakePokersFirstHand(left+1, right, arr)
		// 对方 先手 取走 arr[right] 等于 自己 先手 dfsTakePokersFirstHand(left, right-1, arr)
		p1 := dfsTakePokersFirstHand_B(left+1, right, arr, arrFirstHand, arrBackHand)
		p2 := dfsTakePokersFirstHand_B(left, right-1, arr, arrFirstHand, arrBackHand)
		point = min(p1, p2)
	}
	arrBackHand[left][right] = point
	return point
}

// 动态规划：
func takePokers_C(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	n := len(arr)

	arrFirstHand := make([][]int, n)
	arrBackHand := make([][]int, n)
	for i := range arrFirstHand {
		arrFirstHand[i] = make([]int, n)
	}
	for i := range arrBackHand {
		arrBackHand[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		arrFirstHand[i][i] = arr[i]
	}

	// 每条对角线都从 [0,col] 开始：
	for col := 1; col < n; col++ {
		i, j := 0, col
		for j < n {
			arrFirstHand[i][j] = max(arr[i]+arrBackHand[i+1][j], arr[j]+arrBackHand[i][j-1])
			arrBackHand[i][j] = min(arrFirstHand[i+1][j], arrFirstHand[i][j-1])
			i++
			j++
		}
	}
	return max(arrFirstHand[0][n-1], arrBackHand[0][n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
