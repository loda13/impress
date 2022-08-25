package service

import (
	"github.com/shirou/gopsutil/process"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"time"
)

// FetchCPU 获取CPU信息
func FetchCPU() ([]float64, error) {
	return cpu.Percent(time.Second, false)
}

// FetchMEM 获取内存信息
func FetchMEM() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}

// FetchPROCESSCPU 获取進程CPU信息
func FetchPROCESSCPU(pid int) (float64, error) {
	p := process.Process{Pid: int32(pid)}
	return p.CPUPercent()
}

// FetchPROCESSMEM 获取進程MEM信息
func FetchPROCESSMEM(pid int) (float32, error) {
	p := process.Process{Pid: int32(pid)}
	return p.MemoryPercent()
}
