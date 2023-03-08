package rest

import (
	"net/http"
	"time"

	"github.com/zoglam/pdf-sender-bot/pkg/metrics"
)

func (a *Rest) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metrics.RestMetrics.Counter(r.URL.Path, 500)
		metrics.RestMetrics.Counter(r.URL.Path, 400)
		metrics.RestMetrics.Counter(r.URL.Path, 200)
		rt := metrics.RestMetrics.ResposeTime(r.URL.Path)
		startTime := time.Now()
		next.ServeHTTP(w, r)
		rt.UpdateDuration(startTime)
	})
}
