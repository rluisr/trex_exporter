trex_exporter
=============

[![release](https://github.com/rluisr/trex_exporter/actions/workflows/release.yml/badge.svg)](https://github.com/rluisr/trex_exporter/actions/workflows/release.yml)
[![lint](https://github.com/rluisr/trex_exporter/actions/workflows/lint.yml/badge.svg?branch=master)](https://github.com/rluisr/trex_exporter/actions/workflows/lint.yml)

Get metrics from T-Rex API.

Environment
-----------

Not supported these configs as a flag only environment variables.

Name                        | Default | Require | Description
----------------------------|---------|---------| ----------
LISTEN_PORT                 | 49152   | no      |

Prometheus Configuration
------------------------

```yaml
- job_name: trex
  metrics_path: /metrics
  static_configs:
    - targets:
        - "http://<mining machine ip>?worker=<worker name>"
        - "http://<mining machine ip>?worker=<worker name>"
        - "http://<mining machine ip>?worker=<worker name>"
  relabel_configs:
    - source_labels: [__address__]
      target_label: __param_target
    - source_labels: [__param_target]
      target_label: instance
    - target_label: __address__
      replacement: <your exporter addr>
```

Metrics
-------

```
# HELP trex_gpu_accepted_count gpu accepted count
# TYPE trex_gpu_accepted_count gauge
trex_gpu_accepted_count{device_id="0",name="RTX 2080 Ti",worker="daimon"} 0

# HELP trex_gpu_fan_speed gpu fan speed
# TYPE trex_gpu_fan_speed gauge
trex_gpu_fan_speed{device_id="0",name="RTX 2080 Ti",worker="daimon"} 100

# HELP trex_gpu_hashrate hashrate of the gpu
# TYPE trex_gpu_hashrate gauge
trex_gpu_hashrate{device_id="0",name="RTX 2080 Ti",worker="daimon"} 6.1286126e+07

# HELP trex_gpu_invalid_count gpu invalid count
# TYPE trex_gpu_invalid_count gauge
trex_gpu_invalid_count{device_id="0",name="RTX 2080 Ti",worker="daimon"} 0

# HELP trex_gpu_lhr_lock_count gpu lhr lock count
# TYPE trex_gpu_lhr_lock_count gauge
trex_gpu_lhr_lock_count{device_id="0",name="RTX 2080 Ti",worker="daimon"} 0

# HELP trex_gpu_lhr_tune gpu lhr tune
# TYPE trex_gpu_lhr_tune gauge
trex_gpu_lhr_tune{device_id="0",name="RTX 2080 Ti",worker="daimon"} 0

# HELP trex_gpu_memory_temperature gpu memory temperature
# TYPE trex_gpu_memory_temperature gauge
trex_gpu_memory_temperature{device_id="0",name="RTX 2080 Ti",worker="daimon"} 90

# HELP trex_gpu_power gpu power
# TYPE trex_gpu_power gauge
trex_gpu_power{device_id="0",name="RTX 2080 Ti",worker="daimon"} 0

# HELP trex_gpu_rejected_count gpu rejected count
# TYPE trex_gpu_rejected_count gauge
trex_gpu_rejected_count{device_id="0",name="RTX 2080 Ti",worker="daimon"} 0

# HELP trex_gpu_temperature gpu temperature
# TYPE trex_gpu_temperature gauge
trex_gpu_temperature{device_id="0",name="RTX 2080 Ti",worker="daimon"} 73

# HELP trex_gpu_total number of gpus
# TYPE trex_gpu_total gauge
trex_gpu_total{worker="daimon"} 5

# HELP trex_ping ping between machine and stratum server
# TYPE trex_ping gauge
trex_ping{worker="daimon"} 36

# HELP trex_rejected_count rejected share count
# TYPE trex_rejected_count gauge
trex_rejected_count{worker="daimon"} 0

# HELP trex_solved_count solved block count
# TYPE trex_solved_count gauge
trex_solved_count{worker="daimon"} 0

# HELP trex_total_hashrate hashrate of the worker
# TYPE trex_total_hashrate gauge
trex_total_hashrate{worker="daimon"} 3.40924666e+08
```

Grafana
-------

[https://grafana.com/grafana/dashboards/16218](https://grafana.com/grafana/dashboards/16218)
