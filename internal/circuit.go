package internal

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const CircuitFactMetricName = "house_fact_electric_appliance_circuit_breaker"
const CircuitFactMetricDescription = "The circuit breaker number of an electrical appliance"

type CircuitFact struct {
	Room     string `json:"room"`
	Floor    string `json:"floor"`
	Location string `json:"location"`
	Type     string `json:"type"`
	Breaker  int    `json:"breaker"`
	metric   prometheus.Gauge
}

func (cf *CircuitFact) CreateMetric() {
	cf.metric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: CircuitFactMetricName,
		Help: CircuitFactMetricDescription,
		ConstLabels: prometheus.Labels{
			"room":     cf.Room,
			"floor":    cf.Floor,
			"location": cf.Location,
			"type":     cf.Type,
		},
	})
	cf.metric.Set(float64(cf.Breaker))
}
