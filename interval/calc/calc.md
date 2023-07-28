# Calculation Package Documentation
The calc package provides various mathematical and statistical functions for calculating metrics based on a collection of data points. It includes functions for calculating average, maximum, minimum, standard deviation, count, sum, variance, median, mode, quartiles, percentile, skewness, kurtosis, and covariance.

## Functions
### Avg
```golang
func Avg(data []*global.DataPoint) (float64, error)
```
Calculates the average (mean) value of a collection of data points.
### Max
```go
func Max(data []*global.DataPoint) (float64, error)
```
Finds the maximum value from a collection of data points.

### Min
``` go
func Min(data []*global.DataPoint) (float64, error)
```
Finds the minimum value from a collection of data points.

### StandardDeviation
```go
func StandardDeviation(data []*global.DataPoint) (float64, error)
```
Calculates the standard deviation of a collection of data points.

### Count
```go
func Count(data []*global.DataPoint) int
```
Counts the number of data points in the given collection.

### Sum
```go
func Sum(data []*global.DataPoint) (float64, error)
```
Calculates the sum of values in a collection of data points.

### Variance
```go
func Variance(data []*global.DataPoint) (float64, error)
```
Calculates the variance of a collection of data points.

### Median
```go
func Median(data []*global.DataPoint) (float64, error)
```
Finds the median value from a collection of data points.

### Mode
```go
func Mode(data []*global.DataPoint) ([]float64, error)
```
Finds the mode(s) in a collection of data points (value(s) that appear most frequently).

### Quartiles
```go
func Quartiles(data []*global.DataPoint) (Q1, Q2, Q3 float64, err error)
```
Calculates the quartiles (Q1, Q2, Q3) for a collection of data points.

### Percentile
```go
func Percentile(data []*global.DataPoint, percentile float64) (float64, error)
```
Calculates the specified percentile value from a collection of data points.

### Skewness
```go
func Skewness(data []*global.DataPoint) float64
```
Calculates the skewness of a collection of data points.

### Kurtosis
``` go
func Kurtosis(data []*global.DataPoint) float64
```
Calculates the kurtosis of a collection of data points.

### Covariance
``` go
func Covariance(dataX, dataY []*global.DataPoint) (float64, error)
```
Calculates the covariance between two sets of data points.

### DataPoint Struct
``` go
type DataPoint struct {
	Timestamp time.Time
	EventTime time.Time // New field to store event time
	Value     float64
}
```
Note: The DataPoint struct used in the functions above should be defined in the global package.
### Example Usage
```go
// Example data points
dataPoints := []*global.DataPoint{
{Value: 10},
{Value: 20},
{Value: 30},
// Add more data points as needed
}

// Calculate the average
avgValue, err := calc.Avg(dataPoints)
if err != nil {
// Handle the error
} else {
// Use the average value
fmt.Println("Average:", avgValue)
}

// Calculate the maximum
maxValue, err := calc.Max(dataPoints)
if err != nil {
// Handle the error
} else {
// Use the maximum value
fmt.Println("Maximum:", maxValue)
}

// Calculate the minimum
minValue, err := calc.Min(dataPoints)
if err != nil {
// Handle the error
} else {
// Use the minimum value
fmt.Println("Minimum:", minValue)
}
// Calculate other metrics in a similar manner
```
Please note that you need to import the required packages, including calc and global, to use the provided functions. Also, make sure to create the necessary data points in your application and pass them to the relevant calculation functions.