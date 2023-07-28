# 计算包说明文档
本计算包calc提供了一系列用于统计和计算数据的函数。这些函数包括均值、最大值、最小值、标准差、计数、总和、方差、中位数、众数、四分位数、百分位数、偏度、峰度和协方差的计算。以下是每个函数的详细说明：

## Avg - 均值计算函数
* 函数签名: Avg(data []*global.DataPoint) (float64, error)
* 功能: 根据给定的数据点计算均值。
* 参数:
* data: global.DataPoint类型的数据点切片。
* 返回值:
float64: 计算得到的均值。
error: 如果数据点为空，则返回错误信息。
Max - 最大值计算函数
函数签名: Max(data []*global.DataPoint) (float64, error)
功能: 根据给定的数据点计算最大值。
参数:
data: global.DataPoint类型的数据点切片。
返回值:
float64: 计算得到的最大值。
error: 如果数据点为空，则返回错误信息。
Min - 最小值计算函数
函数签名: Min(data []*global.DataPoint) (float64, error)
功能: 根据给定的数据点计算最小值。
参数:
data: global.DataPoint类型的数据点切片。
返回值:
float64: 计算得到的最小值。
error: 如果数据点为空，则返回错误信息。
StandardDeviation - 标准差计算函数
函数签名: StandardDeviation(data []*global.DataPoint) (float64, error)
功能: 根据给定的数据点计算标准差。
参数:
data: global.DataPoint类型的数据点切片。
返回值:
float64: 计算得到的标准差。
error: 如果数据点为空，则返回错误信息。
Count - 计数计算函数
函数签名: Count(data []*global.DataPoint) int
功能: 计算数据点的数量。
参数:
data: global.DataPoint类型的数据点切片。
返回值:
int: 数据点的数量。
Sum - 求和计算函数
函数签名: Sum(data []*global.DataPoint) (float64, error)
功能: 根据给定的数据点计算数值总和。
参数:
data: global.DataPoint类型的数据点切片。
返回值:
float64: 计算得到的数值总和。
error: 如果数据点为空，则返回错误信息。
Variance - 方差计算函数
函数签名: Variance(data []*global.DataPoint) (float64, error)
功能: 根据给定的数据点计算方差。
参数:
data: global.DataPoint类型的数据点切片。
返回值:
float64: 计算得到的方差。
error: 如果数据点为空，则返回错误信息。
Median - 中位数计算函数
函数签名: Median(data []*global.DataPoint) (float64, error)
功能: 根据给定的数据点计算中位数。
参数:
data: global.DataPoint类型的数据点切片。
返回值:
float64: 计算得到的中位数。
error: 如果数据点为空，则返回错误信息。
Mode - 众数计算函数
函数签名: Mode(data []*global.DataPoint) ([]float64, error)
功能: 根据给定的数据点计算众数。
参数:
data: global.DataPoint类型的数据点切片。
返回值:
[]float64: 计算得到的众数，可能包含多个值。
error: 如果数据点为空或所有值都是唯一的，则返回错误信息。
Quartiles - 四分位数计算函数
函数签名: Quartiles(data []*global.DataPoint) (Q1, Q2, Q3 float64, err error)
功能: 根据给定的数据点计算四分位数。
参数:
data: global.DataPoint类型的数据点切片。
返回值:
Q1: 计算得到的第一四分位数。
Q2: 计算得到的中位数（第二四分位数）。
Q3: 计算得到的第三四分位数。
err: 如果数据点为空，则返回错误信息。
Percentile - 百分位数计算函数
函数签名: Percentile(data []*global.DataPoint, percentile float64) (float64, error)
功能: 根据给定的数据点和百分位数计算相应的百分位数值。
参数:
data: global.DataPoint类型的数据点切片。
percentile: 百分位数，必须介于0到100之间。
返回值:
float64: 计算得到的百分位数值。
error: 如果数据点为空或百分位数无效，则返回错误信息。
Skewness - 偏度计算函数
函数签名: Skewness(data []*global.DataPoint) float64
功能: 根据给定的数据点计算偏度。
参数:
data: global.DataPoint类型的数据点切片。
返回值:
float64: 计算得到的偏度。
Kurtosis - 峰度计算函数
函数签名: Kurtosis(data []*global.DataPoint) float64
功能: 根据给定的数据点计算峰度。
参数:
data: global.DataPoint类型的数据点切片。
返回值:
float64: 计算得到的峰度。
Covariance - 协方差计算函数
函数签名: Covariance(dataX, dataY []*global.DataPoint) (float64, error)
功能: 根据给定的两组数据点计算协方差。
参数:
dataX: global.DataPoint类型的第一组数据点切片。
dataY: global.DataPoint类型的第二组数据点切片，与dataX长度相同。
返回值:
float64: 计算得到的协方差。
error: 如果数据点为空或两组数据点长度不同，则返回错误信息。