package main

/**
 * @Author: tang
 * @mail: yuetang2
 * @Date: 2022/5/28 16:48
 * @Desc: 求所有n的和，其中n属于10000以内的3的倍数的正整数
 */

import (
	"fmt"
)

func Sum() {
	n := 0
	for i := 3; i < 10000; i = i + 3 {
		n = n + i
		//fmt.Printf("main for n  : %d \n", n)
	}
	fmt.Printf("0~10000范围内3的倍数的和为: %d \n", n)
}

func main() {
	Sum()
}
