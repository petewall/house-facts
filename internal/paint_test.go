package internal_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	io_prometheus_client "github.com/prometheus/client_model/go"

	"github.com/petewall/house-facts/internal"
)

var _ = Describe("PaintColor", func() {
	Describe("CreateMetric", func() {
		It("creates a Prometheus metric", func() {
			paintColor := &internal.PaintColor{
				Room:     "office",
				Floor:    "main",
				Location: "ceiling",
				Color:    0,
			}

			mfs, err := metricRegisterer.Gather()
			Expect(err).ToNot(HaveOccurred())
			Expect(mfs).To(HaveLen(0))

			paintColor.CreateMetric()

			mfs, err = metricRegisterer.Gather()
			Expect(err).ToNot(HaveOccurred())
			Expect(mfs).To(HaveLen(1))
			Expect(*mfs[0].Name).To(Equal("house_fact_paint_color_rgb"))
			Expect(*mfs[0].Help).To(Equal("The decimal version of an RGB paint color"))
			Expect(*mfs[0].Type).To(Equal(io_prometheus_client.MetricType_GAUGE))

			Expect(mfs[0].Metric).To(HaveLen(1))
			metric := mfs[0].Metric[0]
			Expect(metric.Label).To(HaveLen(3))
			Expect(*metric.Label[0].Name).To(Equal("floor"))
			Expect(*metric.Label[0].Value).To(Equal("main"))
			Expect(*metric.Label[1].Name).To(Equal("location"))
			Expect(*metric.Label[1].Value).To(Equal("ceiling"))
			Expect(*metric.Label[2].Name).To(Equal("room"))
			Expect(*metric.Label[2].Value).To(Equal("office"))
			Expect(*metric.Gauge.Value).To(BeEquivalentTo(0))
		})
	})
})
