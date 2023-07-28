package main

//
//import (
//	"math"
//	"sync"
//	"time"
//)
//
////对滚动窗口来说，Watermark 机制并不是必需的，因为滚动窗口的触发和清理是基于时间间隔而不是事件时间。
////滚动窗口是一种基于处理时间或事件时间的固定长度窗口，窗口的计算触发是由固定的时间间隔来控制的，而与数据的事件时间无关。
//
//type DataPoint struct {
//	Timestamp time.Time
//	EventTime time.Time // New field to store event time
//	Value     float64
//}
//
//type TumblingWindow struct {
//	Tag           string
//	Metric        string
//	WindowSize    time.Duration
//	Data          []*DataPoint
//	MaxDataPoints int // Maximum number of data points to keep in the window
//	mu            sync.Mutex
//	Watermark     time.Time // Watermark to track event time progress
//}
//
//func NewTumblingWindow(windowSize time.Duration, tag string, metric string, maxDataPoints int) *TumblingWindow {
//	return &TumblingWindow{
//		Tag:           tag,
//		Metric:        metric,
//		WindowSize:    windowSize,
//		Data:          make([]*DataPoint, 0),
//		MaxDataPoints: maxDataPoints,
//		Watermark:     time.Time{}, // Initialize watermark to zero time
//	}
//}
//
//// GenerateWatermark generates the watermark based on the current event time progress
//func (tw *TumblingWindow) GenerateWatermark() {
//	tw.mu.Lock()
//	defer tw.mu.Unlock()
//
//	// Find the minimum event time in the window
//	minEventTime := time.Now()
//	for _, dp := range tw.Data {
//		if dp.EventTime.Before(minEventTime) {
//			minEventTime = dp.EventTime
//		}
//	}
//
//	// Update the watermark to the minimum event time
//	tw.Watermark = minEventTime
//}
//func (tw *TumblingWindow) AddDataPoint(dp DataPoint) {
//	tw.mu.Lock()
//	defer tw.mu.Unlock()
//
//	// Set the event time of the data point to the current time
//	dp.EventTime = time.Now()
//
//	tw.Data = append(tw.Data, &dp)
//	if len(tw.Data) > tw.MaxDataPoints {
//		tw.Data = tw.Data[len(tw.Data)-tw.MaxDataPoints:]
//	}
//
//	// Generate and process the watermark
//	tw.GenerateWatermark()
//	tw.ProcessWatermark()
//}
//func (tw *TumblingWindow) ProcessWatermark() {
//	// Compare the current time with the watermark to trigger window computation
//	if time.Now().Sub(tw.Watermark) >= tw.WindowSize {
//		// Perform window computation here, e.g., CalculateMax(), CalculateMin(), etc.
//		// Clear the data after computation
//		tw.ClearData()
//
//		// Generate the new watermark
//		tw.GenerateWatermark()
//	}
//}
//func (tw *TumblingWindow) ClearData() {
//	tw.mu.Lock()
//	defer tw.mu.Unlock()
//
//	tw.Data = make([]*DataPoint, 0) // Clear the data slice
//}
//func (tw *TumblingWindow) CalculateMax() float64 {
//	if len(tw.Data) == 0 {
//		return 0
//	}
//
//	max := tw.Data[0].Value
//	for _, dp := range tw.Data {
//		if dp.Value > max {
//			max = dp.Value
//		}
//	}
//	return max
//}
//
//func (tw *TumblingWindow) CalculateMin() float64 {
//	if len(tw.Data) == 0 {
//		return 0
//	}
//
//	min := tw.Data[0].Value
//	for _, dp := range tw.Data {
//		if dp.Value < min {
//			min = dp.Value
//		}
//	}
//	return min
//}
//
//func (tw *TumblingWindow) CalculateStandardDeviation() float64 {
//	if len(tw.Data) == 0 {
//		return 0
//	}
//
//	var sum float64
//	for _, dp := range tw.Data {
//		sum += dp.Value
//	}
//	average := sum / float64(len(tw.Data))
//
//	var variance float64
//	for _, dp := range tw.Data {
//		variance += math.Pow(dp.Value-average, 2)
//	}
//	variance /= float64(len(tw.Data))
//	return math.Sqrt(variance)
//}
