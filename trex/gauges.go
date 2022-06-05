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
	TotalHashRate = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "total_hashrate",
		Help:      "hashrate of the worker",
	}, []string{"worker"})
	RejectedCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "rejected_count",
		Help:      "rejected share count",
	}, []string{"worker"})
	SolvedCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "solved_count",
		Help:      "solved block count",
	}, []string{"worker"})
	GPUHashRate = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "gpu_hashrate",
		Help:      "hashrate of the gpu",
	}, []string{"worker", "device_id", "name"})
	GPUPower = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "gpu_power",
		Help:      "gpu power",
	}, []string{"worker", "device_id", "name"})
	GPUTemperature = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "gpu_temperature",
		Help:      "gpu temperature",
	}, []string{"worker", "device_id", "name"})
	GPUMemoryTemperature = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "gpu_memory_temperature",
		Help:      "gpu memory temperature",
	}, []string{"worker", "device_id", "name"})
	GPUFanSpeed = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "gpu_fan_speed",
		Help:      "gpu fan speed",
	}, []string{"worker", "device_id", "name"})
	GPUAcceptedCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "gpu_accepted_count",
		Help:      "gpu accepted count",
	}, []string{"worker", "device_id", "name"})
	GPUInvalidCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "gpu_invalid_count",
		Help:      "gpu invalid count",
	}, []string{"worker", "device_id", "name"})
	GPURejectedCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "gpu_rejected_count",
		Help:      "gpu rejected count",
	}, []string{"worker", "device_id", "name"})
	GPULHRLockCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "gpu_lhr_lock_count",
		Help:      "gpu lhr lock count",
	}, []string{"worker", "device_id", "name"})
	GPULHRTune = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: nameSpace,
		Name:      "gpu_lhr_tune",
		Help:      "gpu lhr tune",
	}, []string{"worker", "device_id", "name"})
)
