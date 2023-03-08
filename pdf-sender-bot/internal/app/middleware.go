package app

import (
	"time"

	"github.com/zoglam/pdf-sender-bot/pkg/metrics"
	"gopkg.in/telebot.v3"
)

func (t *App) MiddleWareMetrics(endpoint string) telebot.MiddlewareFunc {

	metrics.TGMetrics.Counter(endpoint, 500)
	metrics.TGMetrics.Counter(endpoint, 200)
	rt := metrics.TGMetrics.ResposeTime(endpoint)

	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			startTime := time.Now()
			err := next(c)
			rt.UpdateDuration(startTime)
			if err != nil {
				metrics.TGMetrics.Counter(endpoint, 500).Inc()
			} else {
				metrics.TGMetrics.Counter(endpoint, 200).Inc()
			}
			return err
		}
	}

}
