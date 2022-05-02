package trex

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func NewCollector(trexApiAddress, trexWorkerName string) *Collector {
	return &Collector{
		TrexAPIAddress: trexApiAddress,
		TrexWorkerName: trexWorkerName,
	}
}

func (c *Collector) StartCollect(collectInterval time.Duration) {
	go func() {
		for {
			c.collectMetrics()
			time.Sleep(collectInterval)
		}
	}()
}

func (c *Collector) collectMetrics() {
	summary, err := c.getSummary()
	if err != nil {
		writeError(err)
		return
	}

	pingGauge.WithLabelValues(c.TrexWorkerName).Set(float64(summary.ActivePool.Ping))
	gpuTotalGauge.WithLabelValues(c.TrexWorkerName).Set(float64(summary.GPUTotal))
	TotalHashRate.WithLabelValues(c.TrexWorkerName).Set(float64(summary.TotalHashRate))
	RejectedCount.WithLabelValues(c.TrexWorkerName).Set(float64(summary.RejectedCount))
	SolvedCount.WithLabelValues(c.TrexWorkerName).Set(float64(summary.SolvedCount))

	for _, gpuSummary := range summary.GPUs {
		GPUHashRate.WithLabelValues(c.TrexWorkerName, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.HashRate))
		GPUPower.WithLabelValues(c.TrexWorkerName, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.Power))
		GPUTemperature.WithLabelValues(c.TrexWorkerName, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.Temperature))
		GPUMemoryTemperature.WithLabelValues(c.TrexWorkerName, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.MemoryTemperature))
		GPUFanSpeed.WithLabelValues(c.TrexWorkerName, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.FanSpeed))
		GPUAcceptedCount.WithLabelValues(c.TrexWorkerName, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.AcceptedCount))
		GPUInvalidCount.WithLabelValues(c.TrexWorkerName, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.InvalidCount))
		GPURejectedCount.WithLabelValues(c.TrexWorkerName, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.RejectedCount))
		GPULHRLockCount.WithLabelValues(c.TrexWorkerName, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.LHRLockCount))
		GPULHRTune.WithLabelValues(c.TrexWorkerName, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.LHRTune))
	}
}

func writeError(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "[trex_exporter ERROR] %s\n", err.Error())
}
