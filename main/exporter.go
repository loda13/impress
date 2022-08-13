package main

/**
 * @Author: tang
 * @mail: yuetang2
 * @Date: 2022/8/9 16:21
 * @Desc:
 */

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	cpupercent = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpupercent",
		Help: "Current percent of the CPU.",
	})
	mempercent = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "mempercent",
		Help: "Current percent of the MEM.",
	})
)

func GetCpuPercent() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	for {
		percent, _ := cpu.Percent(time.Second, false)
		//fmt.Printf("cpu percent:%v\n", percent)
		//return percent[len(percent)-1]
		cpupercent.Set(percent[len(percent)-1])
	}

}

func GetMemPercent() {
	memInfo, _ := mem.VirtualMemory()
	//return memInfo.UsedPercent
	mempercent.Set(memInfo.UsedPercent)
}

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(cpupercent)
	prometheus.MustRegister(mempercent)
}

func main() {
	go GetCpuPercent()
	go GetMemPercent()
	//cpupercent.Set(GetCpuPercent())
	//mempercent.Set(GetMemPercent())
	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
