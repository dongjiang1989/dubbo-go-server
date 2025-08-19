package metrics

import (
	"dubbo.apache.org/dubbo-go/v3"

	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"dubbo.apache.org/dubbo-go/v3/metrics"
)

func NewMetricsInstanceOption(port int) dubbo.InstanceOption {
	return dubbo.WithMetrics(
		metrics.WithEnabled(),                   // default false
		metrics.WithPrometheus(),                // set prometheus metric, default prometheus
		metrics.WithPrometheusExporterEnabled(), // enable prometheus exporter default false
		metrics.WithPort(port),                  // prometheus http exporter listen at 9099,default 9090
		metrics.WithPath("/metrics"),            // prometheus http exporter url path, default /metrics
		metrics.WithMetadataEnabled(),           // enable metadata center metrics, default true
		metrics.WithRegistryEnabled(),           // enable registry metrics, default true
		metrics.WithConfigCenterEnabled(),       // enable config center metrics, default true

		metrics.WithAggregationEnabled(), // enable rpc metrics aggregations，Most of the time there is no need to turn it on, default false
		metrics.WithAggregationTimeWindowSeconds(30),
		metrics.WithAggregationBucketNum(10), // agg bucket num
	)
}
