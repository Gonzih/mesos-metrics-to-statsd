package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/quipo/statsd"
)

var mesosMaster string
var statsdHost string
var statsdPrefix string

func init() {
	flag.StringVar(&mesosMaster, "mesos-master", "localhost:5050", "mesos master host:port to use to query metrics")
	flag.StringVar(&statsdHost, "statsd-host", "localhost:8125", "statsd host:port that metrics will be reported to")
	flag.StringVar(&statsdPrefix, "statsd-prefix", "mesos.metrics.", "statsd prefix")
	flag.Parse()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	mesosURL := fmt.Sprintf("http://%s/metrics/snapshot", mesosMaster)
	fmt.Printf("statsd is %s with %s prefix\n", statsdHost, statsdPrefix)
	statsdClient := statsd.NewStatsdClient(statsdHost, statsdPrefix)
	err := statsdClient.CreateSocket()
	checkErr(err)

	for {
		fmt.Printf("http get %s\n", mesosURL)
		resp, err := http.Get(mesosURL)
		checkErr(err)

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		checkErr(err)

		var metrics map[string]float64
		err = json.Unmarshal(body, &metrics)
		checkErr(err)

		for k, v := range metrics {
			key := strings.Replace(k, "/", ".", -1)
			fmt.Printf("%s:%f\n", key, v)
			statsdClient.FGauge(key, v)
		}

		time.Sleep(time.Second * 10)
	}
}
