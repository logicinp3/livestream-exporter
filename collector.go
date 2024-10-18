package main

import (
    "sync"

    "github.com/prometheus/client_golang/prometheus"

    "live-supplier-exporter/utils"
)

// Define LiveCollector Structure
type LiveCollector struct {
    mu              sync.Mutex
    frameMetric     *prometheus.GaugeVec
    bitMetric       *prometheus.GaugeVec
}

// Init a Collector，for register
func NewLiveCollector() *LiveCollector {
    // Metrics Label
    //
    // @supplier: "hw" "tx"
    // @project: "g04" "g13" "g16" "g18" "g20" "g23"
    //
    return &LiveCollector{
        frameMetric: prometheus.NewGaugeVec(
            prometheus.GaugeOpts{
                Name: "live_streaming_quality_metric_framerate",
                Help: "framerate data",
            },
            []string{"supplier", "project"},
        ),
        bitMetric: prometheus.NewGaugeVec(
            prometheus.GaugeOpts{
                Name: "live_streaming_quality_metric_bitrate",
                Help: "bitrate data",
            },
            []string{"supplier", "project"},
        ),
    }
}

// Describe 方法用于描述指标
func (c *LiveCollector) Describe(ch chan<- *prometheus.Desc) {
    c.frameMetric.Describe(ch)
    c.bitMetric.Describe(ch)
}

// Collect 方法用于收集数据并更新指标值
func (c *LiveCollector) Collect(ch chan<- prometheus.Metric) {
    c.mu.Lock()
    defer c.mu.Unlock()

    // 获取华为云 live 数据
    utils.HwAPI

    // 获取腾讯云 live 数据
    utils.TcAPI

    // 解析数据
    result := map[string]interface{}{
        "project": "g04",
        "value":   10.12,
    }
    project := result["project"].(string)
    value := result["value"].(float64)

    // 写入指标
    c.frameMetric.WithLabelValues(project).Set(value)
    c.bitMetric.WithLabelValues(project).Set(value)

    // 将指标发送到通道中
    c.frameMetric.Collect(ch)
    c.bitMetric.Collect(ch)
}
