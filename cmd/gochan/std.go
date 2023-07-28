package main

import (
	"fmt"
	ext "go-streaming/streams/extension"
	"go-streaming/streams/flow"
	"go-streaming/utils"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type message struct {
	Msg string
}

func (msg *message) String() string {
	return msg.Msg
}

func main() {

	source := ext.NewChanSource(tickerChan(time.Second * 1))
	//mapFlow := flow.NewMap(addUTC, 1)
	mapFlow := flow.NewMap(EnCode, 1)
	tumblingWindow := flow.NewTumblingWindow(time.Second * 6)

	sink := ext.NewStdoutSink()

	source.
		Via(mapFlow).
		Via(tumblingWindow).
		To(sink)

}

var addUTC = func(msg *message) *message {
	msg.Msg += "AAA"
	return msg
}

func tickerChan(repeat time.Duration) chan interface{} {
	ticker := time.NewTicker(repeat)
	oc := ticker.C
	nc := make(chan interface{})
	go func() {
		for range oc {
			//nc <- &message{strconv.FormatInt(time.Now().UTC().UnixNano(), 10)}
			timeum := time.Now().UnixMilli()
			timestr := strconv.Itoa(int(timeum))
			value := utils.FloatToStr(rand.Float64()*100, 2)
			Msg := "DTHYJK:NSFC:Q1:W001:WNAC_WdSpd_AVG_10m@float:" + value + ":" + timestr
			nc <- &message{Msg}
			fmt.Println("source :" + Msg)
		}
	}()
	return nc
}

type RedisMsg struct {
	Key       string  `json:"key"`
	Metric    string  `json:"metric"`
	DataType  string  `json:"data_type"`
	Value     float64 `json:"value"`
	TimeStamp int64   `json:"time_stamp"`
}

func EnCode(msg *message) *RedisMsg {
	rs := strings.Split(msg.Msg, "@")
	if len(rs) != 2 {
		return nil
	}
	point := rs[0]
	pointValue := rs[1]

	lastColonIndex := strings.LastIndex(point, ":")
	devCode := point[:lastColonIndex]
	metric := point[lastColonIndex+1:]
	pvs := strings.Split(pointValue, ":")
	dataType := pvs[0]
	value, err := strconv.ParseFloat(pvs[1], 6)
	if err != nil {
		panic(err)
	}
	timestamp, err := strconv.ParseInt(pvs[2], 10, 64)
	if err != nil {
		panic(err)
	}
	result := &RedisMsg{
		Key:       devCode,
		Metric:    metric,
		DataType:  dataType,
		Value:     value,
		TimeStamp: timestamp,
	}
	//i, err := json.Marshal(result)
	//if err != nil {
	//	panic(err)
	//}
	return result
}
