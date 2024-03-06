package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	pollingInterval = 10 * time.Second // 轮询间隔时间
	apiURL          = "http://deviceshifu-plate-reader.deviceshifu.svc.cluster.local/get_measurement" // 酶标仪接口URL
)

func main() {
	for {
		// 轮询获取数据
		data, err := fetchMeasurementData()
		if err != nil {
			log.Printf("Error fetching measurement data: %v", err)
		} else {
			// 处理数据
			average := calculateAverage(data)
			fmt.Printf("Average measurement: %.2f\n", average)
		}

		time.Sleep(pollingInterval)
	}
}

// 获取酶标仪数据
func fetchMeasurementData() ([]float64, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析数据
	data := parseMeasurementData(body)

	return data, nil
}

// 解析酶标仪数据
func parseMeasurementData(body []byte) []float64 {
	lines := strings.Split(string(body), "\n")
	var data []float64

	for _, line := range lines {
		// 忽略空行
		if line == "" {
			continue
		}

		values := strings.Split(line, " ")

		for _, value := range values {
			// 解析浮点数
			if value == ""{
				continue
			}
			num, err := strconv.ParseFloat(value, 64)
			if err != nil {
				log.Printf("Error parsing measurement value: %v", err)
				continue
			}

			data = append(data, num)
		}
	}

	return data
}

// 计算平均值
func calculateAverage(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}

	sum := 0.0
	for _, value := range data {
		sum += value
	}

	return sum / float64(len(data))
}