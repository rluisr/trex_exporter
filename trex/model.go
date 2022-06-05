package trex

type Collector struct {
	TrexAPIAddress string
	TrexWorkerName string
}

type summary struct {
	ActivePool    activePool `json:"active_pool"`
	Driver        string     `json:"driver"`
	TotalHashRate int        `json:"hashrate"`
	AcceptedCount int        `json:"accepted_count"`
	RejectedCount int        `json:"rejected_count"`
	SolvedCount   int        `json:"solved_count"`
	GPUTotal      int        `json:"gpu_total"`
	Uptime        int        `json:"uptime"`
	GPUs          []gpuSummary
}

type activePool struct {
	Ping int `json:"ping"`
}

type gpuSummary struct {
	DeviceId          int    `json:"device_id"`
	Name              string `json:"name"`
	HashRate          int
	Power             int
	Temperature       int     `json:"temperature"`
	MemoryTemperature int     `json:"memory_temperature"`
	FanSpeed          int     `json:"fan_speed"`
	AcceptedCount     int     `json:"accepted_count"`
	InvalidCount      int     `json:"invalid_count"`
	RejectedCount     int     `json:"rejected_count"`
	LHRLockCount      int     `json:"lhr_lock_count"`
	LHRTune           float64 `json:"lhr_tune"`
}
