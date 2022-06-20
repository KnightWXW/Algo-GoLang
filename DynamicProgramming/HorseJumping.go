package main

import (
	"fmt"
)

/*
	在中国象棋的棋盘中，把整个棋盘放入第一象限,
	棋盘的最左下角是(0.0)位置，整个棋盘就是横坐标上9条线、纵坐标上10条线的区域。
	给你三个参数x, y. k。
	马走日，返回“马”从(O,0)位置出发，必须走 k 步最后落在(x.y)上的方法数有多少种。
*/

func main() {
	fmt.Print("暴力递归：")
	fmt.Println(horseJumping_A(5, 5, 10)) // 1102475
	fmt.Print("动态规划：")
	fmt.Println(horseJumping_B(5, 5, 10)) // 1102475
}

// 暴力递归：
func horseJumping_A(x, y, k int) int {
	if x < 0 || y < 0 || x > 8 || y > 9 || k <= 0 {
		return 0
	}
	return dfsHorseJumping_A(0, 0, x, y, k)
}

func dfsHorseJumping_A(i, j, x, y, k int) int {
	if i < 0 || j < 0 || i > 8 || j > 9 {
		return 0
	}
	if k == 0 {
		if i == x && j == y {
			return 1
		} else {
			return 0
		}
	}
	cnt := 0
	cnt = dfsHorseJumping_A(i+1, j+2, x, y, k-1) +
		dfsHorseJumping_A(i+2, j+1, x, y, k-1) +
		dfsHorseJumping_A(i+2, j-1, x, y, k-1) +
		dfsHorseJumping_A(i+1, j-2, x, y, k-1) +
		dfsHorseJumping_A(i-1, j-2, x, y, k-1) +
		dfsHorseJumping_A(i-2, j-1, x, y, k-1) +
		dfsHorseJumping_A(i-2, j+1, x, y, k-1) +
		dfsHorseJumping_A(i-1, j+2, x, y, k-1)
	return cnt
}

// 动态规划：
func horseJumping_B(x, y, k int) int {
	if x < 0 || y < 0 || x > 8 || y > 9 || k <= 0 {
		return 0
	}
	var arr [9][10][]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 10; j++ {
			arr[i][j] = make([]int, k+1)
		}
	}
	return dfsHorseJumping_B(x, y, k, arr)
}

func dfsHorseJumping_B(x, y, k int, arr [9][10][]int) int {
	arr[x][y][0] = 1
	for t := 1; t <= k; t++ {
		for i := 0; i < 9; i++ {
			for j := 0; j < 10; j++ {
				cnt := 0
				cnt += pickNumFromArr(i+1, j+2, t-1, arr)
				cnt += pickNumFromArr(i+2, j+1, t-1, arr)
				cnt += pickNumFromArr(i+2, j-1, t-1, arr)
				cnt += pickNumFromArr(i+1, j-2, t-1, arr)
				cnt += pickNumFromArr(i-1, j-2, t-1, arr)
				cnt += pickNumFromArr(i-2, j-1, t-1, arr)
				cnt += pickNumFromArr(i-2, j+1, t-1, arr)
				cnt += pickNumFromArr(i-1, j+2, t-1, arr)
				arr[i][j][t] = cnt
			}
		}
	}
	return arr[0][0][k]
}

func pickNumFromArr(i, j, t int, arr [9][10][]int) int {
	if i < 0 || j < 0 || i > 8 || j > 9 {
		return 0
	} else {
		return arr[i][j][t]
	}
}
