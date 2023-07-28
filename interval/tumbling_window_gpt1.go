package main

import (
	"go-streaming/interval/global"
	"sync"
	"time"
)

type TumblingWindow struct {
	Tag           string
	Metric        string
	WindowSize    time.Duration
	Data          []*global.DataPoint
	MaxDataPoints int // Maximum number of data points to keep in the window
	mu            sync.Mutex
}

func NewTumblingWindow(windowSize time.Duration, tag string, metric string, maxDataPoints int) *TumblingWindow {
	return &TumblingWindow{
		Tag:           tag,
		Metric:        metric,
		WindowSize:    windowSize,
		Data:          make([]*global.DataPoint, 0),
		MaxDataPoints: maxDataPoints,
	}
}

func (tw *TumblingWindow) AddDataPoint(dp global.DataPoint) {
	tw.mu.Lock()
	defer tw.mu.Unlock()

	tw.Data = append(tw.Data, &dp)
	if len(tw.Data) > tw.MaxDataPoints {
		tw.Data = tw.Data[len(tw.Data)-tw.MaxDataPoints:]
	}
}

func (tw *TumblingWindow) ClearData() {
	tw.mu.Lock()
	defer tw.mu.Unlock()

	tw.Data = make([]*global.DataPoint, 0) // Clear the data slice
}
