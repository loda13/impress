package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"net/http"
)

/**
 * @Author: tang
 * @mail: yuetang2
 * @Date: 2022/8/9 16:21
 * @Desc:
 */
const DefaultPerCPU bool = false

func metrics() {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)

}
func fetchCPU() ([]cpu.TimesStat, error) {
	return cpu.Times(DefaultPerCPU)
}

func main() {
	metrics()
	fetchCPU()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
