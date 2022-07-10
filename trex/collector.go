package trex

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"os"
	"regexp"
	"strconv"
)

func Probe(ctx context.Context, target, worker string, registry *prometheus.Registry) bool {
	summary, err := getSummary(target)
	if err != nil {
		writeError(err)
		return false
	}

	registry.MustRegister(driverGauge, pingGauge, uptimeGauge, gpuTotalGauge, TotalHashRate, RejectedCount, SolvedCount, GPUHashRate, GPUPower, GPUTemperature, GPUMemoryTemperature, GPUFanSpeed, GPUAcceptedCount, GPUInvalidCount, GPURejectedCount, GPULHRLockCount, GPULHRTune)

	// TODO I don't know how to T-Rex returns AMD driver version.
	// If only numbers, We should change this code.
	driverFloat64, err := stringToFloat64(summary.Driver)
	if err == nil {
		driverGauge.WithLabelValues(worker).Set(driverFloat64)
	}

	pingGauge.WithLabelValues(worker).Set(float64(summary.ActivePool.Ping))
	uptimeGauge.WithLabelValues(worker).Set(float64(summary.Uptime))
	gpuTotalGauge.WithLabelValues(worker).Set(float64(summary.GPUTotal))
	TotalHashRate.WithLabelValues(worker).Set(float64(summary.TotalHashRate))
	RejectedCount.WithLabelValues(worker).Set(float64(summary.RejectedCount))
	SolvedCount.WithLabelValues(worker).Set(float64(summary.SolvedCount))

	for _, gpuSummary := range summary.GPUs {
		// Support older version
		gpuEfficiency, err := stringToFloat64(gpuSummary.Efficiency)
		if err == nil {
			GPUEfficiency.WithLabelValues(worker, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(gpuEfficiency)
		}

		GPUHashRate.WithLabelValues(worker, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.HashRate))
		GPUPower.WithLabelValues(worker, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.Power))
		GPUTemperature.WithLabelValues(worker, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.Temperature))
		GPUMemoryTemperature.WithLabelValues(worker, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.MemoryTemperature))
		GPUFanSpeed.WithLabelValues(worker, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.FanSpeed))
		GPUAcceptedCount.WithLabelValues(worker, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.AcceptedCount))
		GPUInvalidCount.WithLabelValues(worker, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.InvalidCount))
		GPURejectedCount.WithLabelValues(worker, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.RejectedCount))
		GPULHRLockCount.WithLabelValues(worker, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.LHRLockCount))
		GPULHRTune.WithLabelValues(worker, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(gpuSummary.LHRTune)
	}

	return true
}

func writeError(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "[trex_exporter ERROR] %s\n", err.Error())
}

func stringToFloat64(s string) (float64, error) {
	rex := regexp.MustCompile("kH/W")
	s = rex.ReplaceAllString(s, "")
	return strconv.ParseFloat(s, 64)
}
