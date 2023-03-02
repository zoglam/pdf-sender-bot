package app

import (
	"time"

	"github.com/zoglam/pdf-sender-bot/pkg/metrics"
	"gopkg.in/telebot.v3"
)

func (t *App) MiddleWareMetrics(endpoint string) telebot.MiddlewareFunc {

	if endpoint == "\atext" {
		endpoint = "any_text"
	}

	metrics.TGMetrics.Counter(endpoint, "failed")
	metrics.TGMetrics.Counter(endpoint, "success")
	rt := metrics.TGMetrics.ResposeTime(endpoint)

	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			startTime := time.Now()
			err := next(c)
			rt.UpdateDuration(startTime)
			if err != nil {
				metrics.TGMetrics.Counter(endpoint, "failed").Inc()
			} else {
				metrics.TGMetrics.Counter(endpoint, "success").Inc()
			}
			return err
		}
	}

}
