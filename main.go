package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/showwin/speedtest-go/speedtest"
)

var results = make(Results, 7)

func main() {

	if _, err := os.Stat("results.json"); err != nil {
		results[0].Ping = 0.0
		results[0].Download = 0.0
		results[0].Upload = 0.0
		results[0].MeasureTime = ""
		file, _ := json.MarshalIndent(results, "", " ")
		_ = ioutil.WriteFile("results.json", file, 0777)
	}

	jsonFile, _ := os.Open("results.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &results)

	ping, download, upload, measureTime := testInternetSpeed()

	if results[0].MeasureTime == "" {
		results[0].Ping = ping
		results[0].Download = download
		results[0].Upload = upload
		results[0].MeasureTime = measureTime
	} else if results[1].MeasureTime == "" {
		results[1].Ping = ping
		results[1].Download = download
		results[1].Upload = upload
		results[1].MeasureTime = measureTime
	} else if results[2].MeasureTime == "" {
		results[2].Ping = ping
		results[2].Download = download
		results[2].Upload = upload
		results[2].MeasureTime = measureTime
	} else if results[3].MeasureTime == "" {
		results[3].Ping = ping
		results[3].Download = download
		results[3].Upload = upload
		results[3].MeasureTime = measureTime
	} else if results[4].MeasureTime == "" {
		results[4].Ping = ping
		results[4].Download = download
		results[4].Upload = upload
		results[4].MeasureTime = measureTime
	} else if results[5].MeasureTime == "" {
		results[5].Ping = ping
		results[5].Download = download
		results[5].Upload = upload
		results[5].MeasureTime = measureTime
	} else if results[6].MeasureTime == "" {
		results[6].Ping = ping
		results[6].Download = download
		results[6].Upload = upload
		results[6].MeasureTime = measureTime
	} else {
		temp1 := results[1]
		temp2 := results[2]
		temp3 := results[3]
		temp4 := results[4]
		temp5 := results[5]
		temp6 := results[6]

		results[5] = temp6
		results[4] = temp5
		results[3] = temp4
		results[2] = temp3
		results[1] = temp2
		results[0] = temp1

		results[6].Ping = ping
		results[6].Download = download
		results[6].Upload = upload
		results[6].MeasureTime = measureTime
	}

	file, _ := json.MarshalIndent(results, "", " ")
	_ = ioutil.WriteFile("results.json", file, 0777)

	page := components.NewPage()
	page.AddCharts(
		lineShowLabel(),
	)
	f, err := os.Create("index.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))

}

func lineShowLabel() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Measure the Internet Speed of Sh*tcom every 30 minutes",
			Subtitle: "A litte dirty and buggy tool, that measures the speed of your internet every 30 minutes",
			Link:     "https://github.com/arno4000/go-speedtest",
		}),
	)
	values := []string{

		results[0].MeasureTime,
		results[1].MeasureTime,
		results[2].MeasureTime,
		results[3].MeasureTime,
		results[4].MeasureTime,
		results[5].MeasureTime,
		results[6].MeasureTime,
	}
	downloadValues := make([]opts.LineData, 0)
	uploadValues := make([]opts.LineData, 0)
	pingValues := make([]opts.LineData, 0)

	for i := 0; i <= 6; i++ {
		downloadValues = append(downloadValues, opts.LineData{Value: results[i].Download})
		uploadValues = append(uploadValues, opts.LineData{Value: results[i].Upload})
		pingValues = append(pingValues, opts.LineData{Value: results[i].Ping})
	}

	line.SetXAxis(values).
		AddSeries("Download", downloadValues).
		AddSeries("Upload", uploadValues).
		AddSeries("Ping", pingValues).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
		)
	return line
}

func testInternetSpeed() (ping float64, download float64, upload float64, measureTime string) {
	user, err := speedtest.FetchUserInfo()
	if err != nil {
		log.Fatalln(err)
	}
	serverList, err := speedtest.FetchServerList(user)
	if err != nil {
		log.Fatalln(err)
	}
	targets, err := serverList.FindServer([]int{})
	if err != nil {
		log.Fatalln(err)
	}
	for _, s := range targets {
		s.PingTest()
		s.DownloadTest(true)
		s.UploadTest(true)
		ping = float64(s.Latency.Milliseconds())
		download = math.Round(s.DLSpeed * 100 / 100)
		upload = math.Round(s.ULSpeed * 100 / 100)
	}
	measureTime = time.Now().Format("15:04:05")
	return ping, download, upload, measureTime
}

type Results []struct {
	MeasureTime string  `json:"MeasureTime"`
	Ping        float64 `json:"Ping"`
	Download    float64 `json:"Download"`
	Upload      float64 `json:"Upload"`
}
