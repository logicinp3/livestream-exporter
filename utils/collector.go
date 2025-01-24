package utils

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"livestream-exporter/config"

	"github.com/prometheus/client_golang/prometheus"
)

// Define LiveCollector Structure
type LiveCollector struct {
	mu              sync.Mutex
	framerateMetric *prometheus.GaugeVec
	bitrateMetric   *prometheus.GaugeVec
}

// Init a Collector，for register
func NewLiveCollector() *LiveCollector {
	// Metrics Label
	//
	// @provider: "hw" "tc"
	// @project: "project1" "project2"
	// @streamUrl: "xx-push.live.com/app/stream"
	//
	return &LiveCollector{
		framerateMetric: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "livestream_framerate",
				Help: "CloudPlatform framerate metrics",
			},
			[]string{
				"provider",
				"project",
				"streamUrl",
			},
		),
		bitrateMetric: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "livestream_bitrate",
				Help: "CloudPlatform bitrate metrics",
			},
			[]string{
				"provider",
				"project",
				"streamUrl",
			},
		),
	}
}

// Ddescribe metrics
func (l *LiveCollector) Describe(ch chan<- *prometheus.Desc) {
	l.framerateMetric.Describe(ch)
	l.bitrateMetric.Describe(ch)
}

// collect all metrics data
func (l *LiveCollector) Collect(ch chan<- prometheus.Metric) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// loads config file
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Error loading config: %s", err)
	}
	// goroutine for watches config file
	go config.WatchConfig()
	// get config
	hwProviders := config.AppConfig.Haiwei
	// tcProviders := config.AppConfig.Tencent

	// metric data struct
	type MetricsData struct {
		// provider  string
		project   string
		framerate float64
		bitrate   float64
		streamUrl string
	}
	hwDataSlice := make([]MetricsData, 0)
	tcDataSlice := make([]MetricsData, 0)

	// loop haiwei provider
	for project, provider := range hwProviders {
		hwCloud := NewHaiweicloud(provider.AK, provider.SK, provider.ProjectID)

		// loop push stream list
		for _, stream := range provider.PushStreamList {
			streamSlice := strings.Split(stream, "/")
			// init data struct
			var data MetricsData
			data.project = project
			data.streamUrl = stream
			// getStreamFrameRate
			frameRes, err := hwCloud.GetStreamFrameRate(streamSlice[0], streamSlice[1], streamSlice[2])
			if err != nil {
				fmt.Printf("Failed to get hw stream framerate data: %v", err)
				data.framerate = float64(0)
			} else {
				framerateInfoList := *frameRes.FramerateInfoList
				dataList := func() []int64 {
					if len(*&framerateInfoList) > 0 {
						return *(framerateInfoList)[0].DataList
					}
					return nil
				}()
				firstElement := func() int64 {
					if dataList != nil {
						return dataList[0]
					}
					return 0
				}()
				data.framerate = float64(firstElement)
			}
			// getStreamBitRate
			bitRes, err := hwCloud.GetStreamBitRate(streamSlice[0], streamSlice[1], streamSlice[2])
			if err != nil {
				fmt.Printf("Failed to get hw stream bitrate data: %v", err)
				data.bitrate = float64(0)
			} else {
				bitrateInfoList := *bitRes.BitrateInfoList
				dataList := func() []int64 {
					if len(*&bitrateInfoList) > 0 {
						return *(bitrateInfoList)[0].DataList
					}
					return nil
				}()
				firstElement := func() int64 {
					if dataList != nil {
						return dataList[0]
					}
					return 0
				}()
				data.bitrate = float64(firstElement)
			}
			hwDataSlice = append(hwDataSlice, data)
		}
	}
	fmt.Println(hwDataSlice)

	// loop tecent provider
	tcDataSlice = append(tcDataSlice, MetricsData{"g33", 30.0, 1024.11, "tc-push.live.com/app01/h-5"})

	// write metrics data
	for _, v := range hwDataSlice {
		l.framerateMetric.WithLabelValues("hw", v.project, v.streamUrl).Set(v.framerate)
		l.bitrateMetric.WithLabelValues("hw", v.project, v.streamUrl).Set(v.bitrate)
	}
	for _, v := range tcDataSlice {
		l.framerateMetric.WithLabelValues("tc", v.project, v.streamUrl).Set(v.framerate)
		l.bitrateMetric.WithLabelValues("tc", v.project, v.streamUrl).Set(v.bitrate)
	}

	// send metrics to channel
	l.framerateMetric.Collect(ch)
	l.bitrateMetric.Collect(ch)
}
