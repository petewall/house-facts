package internal_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/prometheus/client_golang/prometheus"
)

func TestInternal(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Internal Suite")
}

var metricRegisterer *prometheus.Registry

var _ = BeforeEach(func() {
	metricRegisterer = prometheus.NewRegistry()
	prometheus.DefaultRegisterer = metricRegisterer
})
