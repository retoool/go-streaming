package global

import "time"

type DataPoint struct {
	Timestamp time.Time
	EventTime time.Time // New field to store event time
	Value     float64
}
