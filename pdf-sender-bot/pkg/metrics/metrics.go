package metrics

import (
	"fmt"

	"github.com/VictoriaMetrics/metrics"
)

type MetricsStruct struct {
	system string // telegram(tg), rest
}

var (
	TGMetrics   = &MetricsStruct{system: "tg"}
	RestMetrics = &MetricsStruct{system: "rest"}
)

func (m *MetricsStruct) ResposeTime(endpoint string) *metrics.Histogram {
	return metrics.GetOrCreateHistogram(fmt.Sprintf(`pdf_requests_duration_seconds{path="%s",type="%s"}`, endpoint, m.system))
}
func (m *MetricsStruct) Counter(endpoint string, code int) *metrics.Counter {
	return metrics.GetOrCreateCounter(fmt.Sprintf(`pdf_requests_count{path="%s",code="%d",type="%s"}`, endpoint, code, m.system))
}
