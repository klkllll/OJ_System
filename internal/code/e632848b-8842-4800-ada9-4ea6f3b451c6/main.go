package main

import (
	"fmt"
)

// 两数之和
func main() {
    var a, b int
	fmt.Println("请输入两个整数：")
	fmt.Scanln(&a, &b)
	sum := a + b
	fmt.Printf("两数之和为：%d", sum)
}
