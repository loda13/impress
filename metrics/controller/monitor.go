package controller

import (
	"github.com/prometheus/client_golang/prometheus"
	"metrics/model/service"
	"strconv"
	"sync"
)

// Metrics metrics指标
type Metrics struct {
	metrics map[string]*prometheus.Desc
	mutex   sync.Mutex
}

type SystemInfo struct {
	CPU, MEM, PROCESSCPU, PROCESSMEM map[string]float64
}

const (
	CPU_METRICS        = "node_cpu_metrics"
	MEM_METRICS        = "node_mem_metrics"
	PROCESSCPU_METRICS = "process_cpu_metrics"
	PROCESSMEM_METRICS = "process_mem_metrics"

	DefaultPerNic    bool   = true
	DefaultMetricsNS string = "prom"
)

func Registry() prometheus.Gatherer {
	registry := prometheus.NewRegistry()
	registry.MustRegister(NewMetrics(DefaultMetricsNS))
	return registry
}

// 初始化metrics
func newMetric(namespace string, metricName string, docString string, labels []string) *prometheus.Desc {
	return prometheus.NewDesc(namespace+"_"+metricName, docString, labels, nil)
}

// NewMetrics 初始化metrics
func NewMetrics(namespace string) *Metrics {
	return &Metrics{
		metrics: map[string]*prometheus.Desc{
			CPU_METRICS:        newMetric(namespace, "cpu_usage", "The description of cpu_metric", []string{"name"}),
			MEM_METRICS:        newMetric(namespace, "mem_usage", "The description of mem_metric", []string{"name"}),
			PROCESSCPU_METRICS: newMetric(namespace, "process_cpu_usage", "The description of process_cpu_usage", []string{"pid", "user", "cmd", "cmd_line", "update_time"}),
			PROCESSMEM_METRICS: newMetric(namespace, "process_mem_usage", "The description of process_mem_usage", []string{"pid", "user", "cmd", "cmd_line", "update_time"}),
		},
	}
}

// Describe 注册metric结构体，然后将metric指标信息放入chan队列中
func (c *Metrics) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range c.metrics {
		ch <- m
	}
}

// Collect 采集metric指标数据，然后将监控指标数据存放到对应的chan队列中
func (c *Metrics) Collect(ch chan<- prometheus.Metric) {

	c.mutex.Lock() // 加锁
	defer c.mutex.Unlock()

	allMetrics, err := fetchAllMetrics()
	if err != nil {
		panic(err)
	}

	for k, v := range allMetrics.CPU {
		ch <- prometheus.MustNewConstMetric(c.metrics[CPU_METRICS], prometheus.CounterValue, v, k)
	}
	for k, v := range allMetrics.MEM {
		ch <- prometheus.MustNewConstMetric(c.metrics[MEM_METRICS], prometheus.CounterValue, v, k)
	}
	res, _ := service.FetchAllRunningProcess()
	var pid, uptime string
	for i := range res {
		processCPUInfo, _ := service.FetchPROCESSCPU(res[i].PID)
		processMEMInfo, _ := service.FetchPROCESSMEM(res[i].PID)
		pid = strconv.Itoa(res[i].PID)
		uptime = strconv.Itoa(res[i].UpdateTime)
		ch <- prometheus.MustNewConstMetric(c.metrics[PROCESSCPU_METRICS], prometheus.CounterValue, float64(processCPUInfo), pid, res[i].User, res[i].Cmd, res[i].CmdLine, uptime)
		ch <- prometheus.MustNewConstMetric(c.metrics[PROCESSMEM_METRICS], prometheus.CounterValue, float64(processMEMInfo), pid, res[i].User, res[i].Cmd, res[i].CmdLine, uptime)
	}

}

// 获取所有信息
func fetchAllMetrics() (SystemInfo, error) {

	cpuInfo, err := service.FetchCPU()
	if err != nil {
		return SystemInfo{}, err
	}

	memInfo, err := service.FetchMEM()
	if err != nil {
		return SystemInfo{}, err
	}

	//processCPUInfo, err := service.FetchPROCESSCPU(2173)
	//if err != nil {
	//	return SystemInfo{}, err
	//}
	//
	//processMEMInfo, err := service.FetchPROCESSMEM(2173)
	//if err != nil {
	//	return SystemInfo{}, err
	//}

	return SystemInfo{
		CPU: map[string]float64{"usage": cpuInfo[0]},
		MEM: map[string]float64{
			"total": float64(memInfo.Total / 1024 / 1024),
			"free":  float64(memInfo.Free / 1024 / 1024),
			"used":  float64(memInfo.Used / 1024 / 1024),
			"usage": memInfo.UsedPercent,
		},
	}, nil
	//PROCESSCPU: map[string]float64{
	//	res[i]: float64(processCPUInfo),
	//},
	//PROCESSMEM: map[string]float64{
	//	res[i]: float64(processMEMInfo),
	//},
}
