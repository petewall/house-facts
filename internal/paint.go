package internal

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const PaintColorMetricName = "house_fact_paint_color_rgb"
const PaintColorMetricDescription = "The decimal version of an RGB paint color"

type PaintColor struct {
	Room     string
	Floor    string
	Location string
	Color    int
	metric   prometheus.Gauge
}

func (pc *PaintColor) CreateMetric() {
	pc.metric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: PaintColorMetricName,
		Help: PaintColorMetricDescription,
		ConstLabels: prometheus.Labels{
			"room":     pc.Room,
			"floor":    pc.Floor,
			"location": pc.Location,
		},
	})
	pc.metric.Set(float64(pc.Color))
}
