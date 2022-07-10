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

	registry.MustRegister(driverGauge, pingGauge, uptimeGauge, gpuTotalGauge, totalHashRate, rejectedCount, solvedCount, gpuHashRate, gpuPower, gpuEfficiency, gpuTemperature, gpuMemoryTemperature, gpuFanSpeed, gpuAcceptedCount, gpuInvalidCount, gpuRejectedCount, gpuLHRLockCount, gpuLHRTune)

	// TODO I don't know how to T-Rex returns AMD driver version.
	// If only numbers, We should change this code.
	driverFloat64, err := stringToFloat64(summary.Driver)
	if err == nil {
		driverGauge.WithLabelValues(worker).Set(driverFloat64)
	}

	pingGauge.WithLabelValues(worker).Set(float64(summary.ActivePool.Ping))
	uptimeGauge.WithLabelValues(worker).Set(float64(summary.Uptime))
	gpuTotalGauge.WithLabelValues(worker).Set(float64(summary.GPUTotal))
	totalHashRate.WithLabelValues(worker).Set(float64(summary.TotalHashRate))
	rejectedCount.WithLabelValues(worker).Set(float64(summary.RejectedCount))
	solvedCount.WithLabelValues(worker).Set(float64(summary.SolvedCount))

	for _, gpuSummary := range summary.GPUs {
		// Support older version
		gpuEfficiencyFloat64, err := stringToFloat64(gpuSummary.Efficiency)
		if err == nil {
			gpuEfficiency.WithLabelValues(worker, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(gpuEfficiencyFloat64)
		}

		gpuHashRate.WithLabelValues(worker, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.HashRate))
		gpuPower.WithLabelValues(worker, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.Power))
		gpuTemperature.WithLabelValues(worker, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.Temperature))
		gpuMemoryTemperature.WithLabelValues(worker, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.MemoryTemperature))
		gpuFanSpeed.WithLabelValues(worker, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.FanSpeed))
		gpuAcceptedCount.WithLabelValues(worker, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.AcceptedCount))
		gpuInvalidCount.WithLabelValues(worker, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.InvalidCount))
		gpuRejectedCount.WithLabelValues(worker, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.RejectedCount))
		gpuLHRLockCount.WithLabelValues(worker, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(float64(gpuSummary.LHRLockCount))
		gpuLHRTune.WithLabelValues(worker, strconv.Itoa(gpuSummary.DeviceId), gpuSummary.Name).Set(gpuSummary.LHRTune)
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
