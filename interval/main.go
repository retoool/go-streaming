package main

import (
	"fmt"
	"go-streaming/interval/calc"
	"go-streaming/interval/global"
	"math/rand"
	"time"
)

func main() {
	tumblingWindow := NewTumblingWindow(5*time.Second, "DTHYJK:Q1:NSFC:W001", "WNAC_WdSpd", 100) // 设置Tumbling Window的大小为5秒

	// 模拟数据源每隔1秒发送一个数据
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			// 模拟从数据源获取数据，这里简单地使用随机数生成一个数据点
			dataPoint := global.DataPoint{
				Timestamp: time.Now(),
				Value:     rand.Float64() * 100, // 在0到100之间生成一个随机数
			}
			fmt.Println("_____", dataPoint)
			// 将数据点添加到Tumbling Window
			tumblingWindow.AddDataPoint(dataPoint)
		}
	}()

	// 定期处理Tumbling Window的数据
	tickerWindow := time.NewTicker(5 * time.Second) // 设置处理Tumbling Window的时间间隔
	defer tickerWindow.Stop()

	for range tickerWindow.C {
		// 处理当前窗口的数据
		//tumblingWindow.ProcessWindow()
		//max := tumblingWindow.CalculateMax()
		result, err := calc.Avg(tumblingWindow.Data)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Window Mode:", result)
	}
}
