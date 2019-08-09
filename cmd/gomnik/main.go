package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/ncaunt/gomnik"
	"github.com/oklog/run"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	serial   = kingpin.Flag("serial", "Serial number of the inverter.").OverrideDefaultFromEnvar("OMNIK_SERIAL").Short('s').Required().Int()
	addr     = kingpin.Flag("address", "Address on which the inverter listens (example: 10.0.0.1:8899).").OverrideDefaultFromEnvar("OMNIK_ADDR").Short('a').Required().String()
	interval = kingpin.Flag("interval", "Number of seconds between queries of the inverter.").OverrideDefaultFromEnvar("OMNIK_INTERVAL").Short('i').Default("10").Int()
	metrics  = kingpin.Flag("metrics", "Endpoint on which to serve Prometheus metrics.").OverrideDefaultFromEnvar("OMNIK_METRICS").Short('m').Default("0.0.0.0:9100").String()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()

	reg := prometheus.NewRegistry()
	reg.MustRegister(
		version.NewCollector("gomnik"),
		prometheus.NewGoCollector(),
		prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}),
	)

	m := gomnik.NewMetrics(reg)

	var g run.Group
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	httpBindAddr := *metrics
	l, err := net.Listen("tcp", httpBindAddr)
	if err != nil {
		log.Fatal(errors.Wrap(err, "listen metrics address"))
	}

	g.Add(func() error {
		log.Fatal(errors.Wrap(http.Serve(l, mux), "serve metrics"))
		return nil
	}, func(error) {
	})

	ticker := time.NewTicker(time.Duration(*interval) * time.Second)
	g.Add(func() error {
		return loop(ticker, m)
	}, func(err error) {
		fmt.Println(err)
		ticker.Stop()
	})

	g.Run()
}
