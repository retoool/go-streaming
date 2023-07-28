package calc

import (
	"errors"
	"go-streaming/interval/global"
	"math"
	"sort"
)

// Avg 均值计算函数，根据给定的数据点计算均值
func Avg(data []*global.DataPoint) (float64, error) {
	if len(data) == 0 {
		return 0, errors.New("数据点为空")
	}

	var sum float64
	for _, dp := range data {
		sum += dp.Value
	}
	return sum / float64(len(data)), nil
}

// Max 最大值计算函数，根据给定的数据点计算最大值
func Max(data []*global.DataPoint) (float64, error) {
	if len(data) == 0 {
		return 0, errors.New("数据点为空")
	}

	max := data[0].Value
	for _, dp := range data {
		if dp.Value > max {
			max = dp.Value
		}
	}
	return max, nil
}

// Min 最小值计算函数，根据给定的数据点计算最小值
func Min(data []*global.DataPoint) (float64, error) {
	if len(data) == 0 {
		return 0, errors.New("数据点为空")
	}

	min := data[0].Value
	for _, dp := range data {
		if dp.Value < min {
			min = dp.Value
		}
	}
	return min, nil
}

// StandardDeviation 标准差计算函数，根据给定的数据点计算标准差
func StandardDeviation(data []*global.DataPoint) (float64, error) {
	if len(data) == 0 {
		return 0, errors.New("数据点为空")
	}

	var sum float64
	for _, dp := range data {
		sum += dp.Value
	}
	average := sum / float64(len(data))

	var variance float64
	for _, dp := range data {
		variance += math.Pow(dp.Value-average, 2)
	}
	variance /= float64(len(data))
	return math.Sqrt(variance), nil
}

// Count 计数计算函数，计算数据点的数量
func Count(data []*global.DataPoint) int {
	return len(data)
}

// Sum 求和计算函数，根据给定的数据点计算数值总和
func Sum(data []*global.DataPoint) (float64, error) {
	if len(data) == 0 {
		return 0, errors.New("数据点为空")
	}

	var sum float64
	for _, dp := range data {
		sum += dp.Value
	}
	return sum, nil
}

// Variance 方差计算函数，根据给定的数据点计算方差
func Variance(data []*global.DataPoint) (float64, error) {
	if len(data) == 0 {
		return 0, errors.New("数据点为空")
	}

	var sum float64
	for _, dp := range data {
		sum += dp.Value
	}
	average := sum / float64(len(data))

	var variance float64
	for _, dp := range data {
		variance += math.Pow(dp.Value-average, 2)
	}
	variance /= float64(len(data))
	return variance, nil
}

// Median 中位数计算函数，根据给定的数据点计算中位数
func Median(data []*global.DataPoint) (float64, error) {
	if len(data) == 0 {
		return 0, errors.New("数据点为空")
	}

	sortedData := make([]float64, len(data))
	for i, dp := range data {
		sortedData[i] = dp.Value
	}

	sort.Float64s(sortedData)
	n := len(sortedData)

	if n%2 == 0 {
		// 偶数个数据点，取中间两个数的平均值
		return (sortedData[n/2-1] + sortedData[n/2]) / 2, nil
	}

	// 奇数个数据点，返回中间的数值
	return sortedData[n/2], nil
}

// Mode 众数计算函数，根据给定的数据点计算众数
func Mode(data []*global.DataPoint) ([]float64, error) {
	if len(data) == 0 {
		return nil, errors.New("数据点为空")
	}

	freqMap := make(map[float64]int)
	for _, dp := range data {
		freqMap[dp.Value]++
	}

	var maxFreq int
	var modes []float64
	for value, freq := range freqMap {
		if freq > maxFreq {
			maxFreq = freq
			modes = []float64{value}
		} else if freq == maxFreq {
			modes = append(modes, value)
		}
	}

	if maxFreq == 1 {
		return nil, errors.New("没有众数，所有值都是唯一的")
	}

	return modes, nil
}

// Quartiles 四分位数计算函数，根据给定的数据点计算四分位数
func Quartiles(data []*global.DataPoint) (Q1, Q2, Q3 float64, err error) {
	if len(data) == 0 {
		err = errors.New("数据点为空")
		return
	}

	sortedData := make([]float64, len(data))
	for i, dp := range data {
		sortedData[i] = dp.Value
	}

	sort.Float64s(sortedData)
	n := len(sortedData)

	Q2, _ = Median(data)

	midIndex := n / 2
	if n%2 == 0 {
		Q1, _ = Median(data[:midIndex])
		Q3, _ = Median(data[midIndex:])
	} else {
		Q1, _ = Median(data[:midIndex+1])
		Q3, _ = Median(data[midIndex:])
	}

	return
}

// Percentile 百分位数计算函数，根据给定的数据点和百分位数计算相应的百分位数值
func Percentile(data []*global.DataPoint, percentile float64) (float64, error) {
	if len(data) == 0 {
		return 0, errors.New("数据点为空")
	}

	if percentile < 0 || percentile > 100 {
		return 0, errors.New("百分位数必须介于0到100之间")
	}

	sortedData := make([]float64, len(data))
	for i, dp := range data {
		sortedData[i] = dp.Value
	}

	sort.Float64s(sortedData)
	n := len(sortedData)

	rank := (percentile / 100) * float64(n-1)
	lowerIndex := int(rank)
	upperIndex := lowerIndex + 1

	if lowerIndex >= n-1 {
		return sortedData[n-1], nil
	}

	weight := rank - float64(lowerIndex)

	return sortedData[lowerIndex]*(1-weight) + sortedData[upperIndex]*weight, nil
}

// Skewness 偏度计算函数，根据给定的数据点计算偏度
func Skewness(data []*global.DataPoint) float64 {
	if len(data) == 0 {
		return 0
	}

	var sum float64
	var sumSqr float64
	for _, dp := range data {
		sum += dp.Value
		sumSqr += dp.Value * dp.Value
	}

	n := float64(len(data))
	mean := sum / n
	variance := (sumSqr/n - mean*mean) * (n / (n - 1))
	stdDev := math.Sqrt(variance)

	var sumCub float64
	for _, dp := range data {
		sumCub += math.Pow(dp.Value-mean, 3)
	}

	return sumCub / (n * stdDev * stdDev * stdDev)
}

// Kurtosis 峰度计算函数，根据给定的数据点计算峰度
func Kurtosis(data []*global.DataPoint) float64 {
	if len(data) == 0 {
		return 0
	}

	var sum float64
	var sumSqr float64
	for _, dp := range data {
		sum += dp.Value
		sumSqr += dp.Value * dp.Value
	}

	n := float64(len(data))
	mean := sum / n
	variance := (sumSqr/n - mean*mean) * (n / (n - 1))

	var sumFourth float64
	for _, dp := range data {
		sumFourth += math.Pow(dp.Value-mean, 4)
	}

	return sumFourth / (n * variance * variance)
}

// Covariance 协方差计算函数，根据给定的两组数据点计算协方差
func Covariance(dataX, dataY []*global.DataPoint) (float64, error) {
	if len(dataX) != len(dataY) || len(dataX) == 0 {
		return 0, errors.New("无效的数据点")
	}

	var sumXY float64
	var sumX float64
	var sumY float64

	for i := 0; i < len(dataX); i++ {
		sumXY += dataX[i].Value * dataY[i].Value
		sumX += dataX[i].Value
		sumY += dataY[i].Value
	}

	n := float64(len(dataX))
	return (sumXY - sumX*sumY/n) / (n - 1), nil
}
