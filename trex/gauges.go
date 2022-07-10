package trex

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var nameSpace = "trex"

var (
	uptimeGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "uptime",
		Help:      "miner uptime",
	}, []string{"worker"})
	driverGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "driver",
		Help:      "driver version",
	}, []string{"worker"})
	pingGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "ping",
		Help:      "ping between machine and stratum server",
	}, []string{"worker"})
	gpuTotalGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "gpu_total",
		Help:      "number of gpus",
	}, []string{"worker"})
	totalHashRate = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "total_hashrate",
		Help:      "hashrate of the worker",
	}, []string{"worker"})
	rejectedCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "rejected_count",
		Help:      "rejected share count",
	}, []string{"worker"})
	solvedCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "solved_count",
		Help:      "solved block count",
	}, []string{"worker"})
	gpuHashRate = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "gpu_hashrate",
		Help:      "hashrate of the gpu",
	}, []string{"worker", "device_id", "name"})
	gpuPower = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "gpu_power",
		Help:      "gpu power",
	}, []string{"worker", "device_id", "name"})
	gpuEfficiency = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "gpu_efficiency",
		Help:      "gpu efficiency",
	}, []string{"worker", "device_id", "name"})
	gpuTemperature = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "gpu_temperature",
		Help:      "gpu temperature",
	}, []string{"worker", "device_id", "name"})
	gpuMemoryTemperature = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "gpu_memory_temperature",
		Help:      "gpu memory temperature",
	}, []string{"worker", "device_id", "name"})
	gpuFanSpeed = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "gpu_fan_speed",
		Help:      "gpu fan speed",
	}, []string{"worker", "device_id", "name"})
	gpuAcceptedCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "gpu_accepted_count",
		Help:      "gpu accepted count",
	}, []string{"worker", "device_id", "name"})
	gpuInvalidCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "gpu_invalid_count",
		Help:      "gpu invalid count",
	}, []string{"worker", "device_id", "name"})
	gpuRejectedCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "gpu_rejected_count",
		Help:      "gpu rejected count",
	}, []string{"worker", "device_id", "name"})
	gpuLHRLockCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "gpu_lhr_lock_count",
		Help:      "gpu lhr lock count",
	}, []string{"worker", "device_id", "name"})
	gpuLHRTune = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "gpu_lhr_tune",
		Help:      "gpu lhr tune",
	}, []string{"worker", "device_id", "name"})
)
