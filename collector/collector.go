package collector

import (
    "encoding/json"
    "net/http"
    "sync"

    "github.com/prometheus/client_golang/prometheus"
)

// 定义用于存储指标的结构体
type DataCollector struct {
    mu          sync.Mutex
    gaugeMetric *prometheus.GaugeVec
}

// 创建新的数据收集器
func NewDataCollector() *DataCollector {
    return &DataCollector{
        gaugeMetric: prometheus.NewGaugeVec(
            prometheus.GaugeOpts{
                Name: "third_party_data_metric",
                Help: "Gauge metric for data from third party API",
            },
            []string{"status"},
        ),
    }
}

// Describe 方法用于描述指标
func (c *DataCollector) Describe(ch chan<- *prometheus.Desc) {
    c.gaugeMetric.Describe(ch)
}

// Collect 方法用于收集数据并更新指标值
func (c *DataCollector) Collect(ch chan<- prometheus.Metric) {
    c.mu.Lock()
    defer c.mu.Unlock()

    // 从第三方接口获取数据
    resp, err := http.Get("https://api.example.com/data")
    if err != nil {
        log.Printf("Error fetching data: %v", err)
        return
    }
    defer resp.Body.Close()

    // 解析 JSON 数据
    var result map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        log.Printf("Error decoding JSON: %v", err)
        return
    }

    // 假设我们从 JSON 中提取出一个状态值并更新 Gauge 指标
    status := result["status"].(string)
    value := result["value"].(float64)

    c.gaugeMetric.WithLabelValues(status).Set(value)

    // 将指标发送到通道中
    c.gaugeMetric.Collect(ch)
}