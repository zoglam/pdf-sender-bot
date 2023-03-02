package metrics

import (
	"fmt"

	"github.com/VictoriaMetrics/metrics"
)

type TGMetricsStruct struct{}

var (
	TGMetrics = &TGMetricsStruct{}
)

func (m *TGMetricsStruct) ResposeTime(endpoint string) *metrics.Histogram {
	return metrics.GetOrCreateHistogram(fmt.Sprintf(`bg_tg_requests_duration_seconds{path="%s"}`, endpoint))
}
func (m *TGMetricsStruct) Counter(endpoint string, code string) *metrics.Counter {
	return metrics.GetOrCreateCounter(fmt.Sprintf(`bg_tg_requests_count{path="%s",code="%s"}`, endpoint, code))
}
