package main

//
//type DataPoint struct {
//	Timestamp time.Time
//	Value     float64
//}
//
//type TumblingWindow struct {
//	Tag        string
//	Metric     string
//	WindowSize time.Duration
//	Data       []DataPoint
//	mu         sync.Mutex
//}
//
//func NewTumblingWindow(windowSize time.Duration, tag string, metric string) *TumblingWindow {
//	return &TumblingWindow{
//		Tag:        tag,
//		Metric:     metric,
//		WindowSize: windowSize,
//		Data:       make([]DataPoint, 0),
//	}
//}
//
//func (tw *TumblingWindow) AddDataPoint(dp DataPoint) {
//	tw.mu.Lock()
//	defer tw.mu.Unlock()
//
//	tw.Data = append(tw.Data, dp)
//}
//
//func (tw *TumblingWindow) ProcessWindow() {
//	tw.mu.Lock()
//	defer tw.mu.Unlock()
//
//	// Process the data in the current window
//	if len(tw.Data) > 0 {
//		var sum float64
//		for _, dp := range tw.Data {
//			sum += dp.Value
//		}
//		average := sum / float64(len(tw.Data))
//		fmt.Println("Window Average:", average)
//
//		// Clear the data in the window for the next window
//
//	}
//}
//
//func (tw *TumblingWindow) CalculateAverage() float64 {
//	tw.mu.Lock()
//	defer tw.mu.Unlock()
//
//	if len(tw.Data) == 0 {
//		return 0
//	}
//
//	var sum float64
//	for _, dp := range tw.Data {
//		sum += dp.Value
//	}
//	average := sum / float64(len(tw.Data))
//	tw.Data = make([]DataPoint, 0)
//	return average
//}
//
//func (tw *TumblingWindow) CalculateMax() float64 {
//	tw.mu.Lock()
//	defer tw.mu.Unlock()
//
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
//	tw.Data = make([]DataPoint, 0)
//	return max
//}
//
//func (tw *TumblingWindow) CalculateMin() float64 {
//	tw.mu.Lock()
//	defer tw.mu.Unlock()
//
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
//	tw.Data = make([]DataPoint, 0)
//	return min
//}
//
//func (tw *TumblingWindow) CalculateStandardDeviation() float64 {
//	tw.mu.Lock()
//	defer tw.mu.Unlock()
//
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
//	tw.Data = make([]DataPoint, 0)
//	return math.Sqrt(variance)
//}
