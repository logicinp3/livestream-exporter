package utils

import (
    "sync"

    "github.com/prometheus/client_golang/prometheus"
)

// Define LiveCollector Structure
type LiveCollector struct {
    mu              sync.Mutex
    framerateMetric     *prometheus.GaugeVec
    bitrateMetric       *prometheus.GaugeVec
}

// Init a Collector，for register
func NewLiveCollector() *LiveCollector {
    // Metrics Label
    //
    // @supplier: "hw" "tc"
    // @project: "g04" "g13" "g16" "g20" "g23" "g31"
    //
    return &LiveCollector{
        framerateMetric: prometheus.NewGaugeVec(
            prometheus.GaugeOpts{
                Name: "livestreaming_framerate",
                Help: "CloudPlatform framerate metrics",
            },
            []string{
                "supplier", 
                "project",
            },
        ),
        bitrateMetric: prometheus.NewGaugeVec(
            prometheus.GaugeOpts{
                Name: "livestreaming_bitrate",
                Help: "CloudPlatform bitrate metrics",
            },
            []string{
                "supplier",
                "project",
            },
        ),
    }
}

// Describe 方法用于描述指标
func (l *LiveCollector) Describe(ch chan<- *prometheus.Desc) {
    l.framerateMetric.Describe(ch)
    l.bitrateMetric.Describe(ch)
}

// Collect 方法用于收集数据并更新指标值
func (l *LiveCollector) Collect(ch chan<- prometheus.Metric) {
    l.mu.Lock()
    defer l.mu.Unlock()

    // 获取华为云 live 数据, 结构体数据
    // hwData := 
    // HaiweiAPI()

    // 获取腾讯云 live 数据
    // TencentAPI()

    // 模拟获取数据（可以根据实际场景替换为真实数据来源）
    data := []struct {
        supplier string
        project  string
        framerate float64
        bitrate   float64
    }{
        {"hw", "g31", 30.0, 10.11},
        {"tc", "g13", 40.0, 20.11},
        {"tc", "g20", 28.5, 12.34},
    }

    // 动态设置每个 supplier 和 project 标签的指标数据
    for _, entry := range data {
        l.framerateMetric.WithLabelValues(entry.supplier, entry.project).Set(entry.framerate)
        l.bitrateMetric.WithLabelValues(entry.supplier, entry.project).Set(entry.bitrate)
    }

    // 将指标发送到通道中
    l.framerateMetric.Collect(ch)
    l.bitrateMetric.Collect(ch)
}
