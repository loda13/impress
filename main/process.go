package main

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
)

/**
 * @Author: tang
 * @mail: yuetang2
 * @Date: 2022/8/23 19:02
 * @Desc:
 */

func main() {
	info, _ := process.Pids() //获取当前所有进程的pid
	fmt.Println(info)

	for _, pid := range info {
		var p process.Process
		p = process.Process{Pid: pid}
		fmt.Println(p.CPUPercent())
		fmt.Println(p.MemoryPercent())
		fmt.Println("--------------------------------")
	}

}
