package utils

import (
	"math"
	"strconv"
	"strings"
)

func StrToInt64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}
func IntToStr(i int) string {
	s := strconv.Itoa(i)
	return s
}
func Int64ToStr(i int64) string {
	s := strconv.FormatInt(i, 10)
	return s
}

func StrToFloat(str string) float64 {
	float, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(err)
	}
	return float
}

func FloatToStr(f float64, prec int) string {
	s := strconv.FormatFloat(f, 'f', prec, 64)
	return s
}

func Round(f float64, n int) float64 {
	scale := math.Pow(10, float64(n))
	return math.Round(f*scale) / scale
}

func FindNearestIndex(arr []float64, val float64) int {
	nearestIdx := 0
	nearestDiff := math.Abs(arr[0] - val)
	for i := 1; i < len(arr); i++ {
		diff := math.Abs(arr[i] - val)
		if diff < nearestDiff {
			nearestIdx = i
			nearestDiff = diff
		}
	}
	return nearestIdx
}

func GetFTCodeByFullCode(HashKey string) []string {
	hashKeySplits := strings.Split(HashKey, ":")
	project := hashKeySplits[0]
	farm := hashKeySplits[1]
	term := hashKeySplits[2]
	term_full := strings.Join([]string{project, farm, term}, ":")
	farm_full := strings.Join([]string{project, farm}, ":")
	return []string{term_full, farm_full}
}
