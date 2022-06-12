package main

import "fmt"

/*
	在象棋的棋盘中，把整个棋盘放入第一象限,
	棋盘的最左下角是(0.0)位置，整个棋盘就是横坐标上9条线、纵坐标上10条线的区域。
	给你三个参数x, y. k。
	返回“马”从(O,0)位置出发，必须走k步最后落在(x.y)上的方法数有多少种。
*/

func main() {
	fmt.Print("暴力递归：")
	fmt.Println(HorseJumping_A())
	fmt.Print("记忆化搜索：")
	fmt.Println(RobotWalk_B(5, 2, 4, 4))
	fmt.Print("动态规划：")
	fmt.Println(RobotWalk_C(5, 2, 4, 4))
}

func HorseJumping_A(x, y, z int) int {

}
