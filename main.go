package main

import (
	"log"
	"math/rand"
	"net/http"
	"regexp"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var registerp1 = prometheus.NewRegistry()
var registerp2 = prometheus.NewRegistry()
var register = prometheus.NewRegistry()

var (
	// func NewDesc(fqName, help string,
	//		variableLabels []string, constLabels Labels) *Desc
	p1Desc = prometheus.NewDesc(
		"func_1",
		"is successfully running",
		[]string{"p1_string"},
		nil,
	)
	p2Desc = prometheus.NewDesc(
		"func_2",
		"is successfully running",
		[]string{"p2_string"},
		nil,
	)
	invalidChars = regexp.MustCompile("[^a-zA-Z0-9:_]")
)

type p1Collector struct {
	// declaring an empty structñ
}

type p2Collector struct {
	// declaring an empty struct
}

func (p1 p1Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- p1Desc
}

func (p1 p1Collector) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(p1Desc, prometheus.GaugeValue, 1, "p1_string")
	c := rand.Int()
	name := invalidChars.ReplaceAllLiteralString("p1_metric", "_")
	desc := prometheus.NewDesc(name+"_total", "p1 metric ", []string{"name"}, nil)
	ch <- prometheus.MustNewConstMetric(desc, prometheus.CounterValue, float64(c), "priyam")
}

func (p2 p2Collector) Describe(ch2 chan<- *prometheus.Desc) {
	ch2 <- p2Desc
}

func (p2 p2Collector) Collect(ch2 chan<- prometheus.Metric) {
	ch2 <- prometheus.MustNewConstMetric(p2Desc, prometheus.GaugeValue, 1, "p2_string")
	c := rand.Int()
	name := invalidChars.ReplaceAllLiteralString("p2_metric", "_")
	desc := prometheus.NewDesc(name+"_total", "p2 metric ", []string{"name"}, nil)
	ch2 <- prometheus.MustNewConstMetric(desc, prometheus.CounterValue, float64(c), "priyam_2")
}

func main() {

	p1 := p1Collector{}
	p2 := p2Collector{}

	registerp1.MustRegister(p1)
	registerp2.MustRegister(p2)
	register.MustRegister(p1)
	register.MustRegister(p2)

	http.Handle("/p1/metrics", promhttp.HandlerFor(registerp1, promhttp.HandlerOpts{}))
	http.Handle("/p2/metrics", promhttp.HandlerFor(registerp2, promhttp.HandlerOpts{}))
	http.Handle("/metrics", promhttp.HandlerFor(register, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
