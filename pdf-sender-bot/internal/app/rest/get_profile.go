package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/zoglam/pdf-sender-bot/pkg/logg"
	"github.com/zoglam/pdf-sender-bot/pkg/metrics"
)

func (a *Rest) GetProfile(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	ctx = context.WithValue(ctx, "x-request-id", r.Header.Get("X-Request-Id"))
	ctx = context.WithValue(ctx, "x-telegram-id", r.Header.Get("X-Telegram-Id"))

	logg.Info(ctx).Msgf("Пришел запрос на get %+v", r.Body)

	id, err := strconv.ParseInt(r.Header.Get("X-Telegram-Id"), 10, 64)
	if err != nil {
		logg.Error(ctx).Msgf("Error ParseInt id: %v", err)
		http.Error(w, "Error ParseInt id", http.StatusBadRequest)
		metrics.RestMetrics.Counter(r.URL.Path, http.StatusBadRequest).Inc()
		return
	}

	user, err := a.userService.GetUserProfile(id)
	if err != nil {
		logg.Error(ctx).Msgf("Error GetUserProfile: %v", err)
		http.Error(w, "Error GetUserProfile", http.StatusBadRequest)
		metrics.RestMetrics.Counter(r.URL.Path, http.StatusBadRequest).Inc()
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		logg.Error(ctx).Msgf("Error Marshal: %v", err)
		http.Error(w, "Error Marshal", http.StatusBadRequest)
		metrics.RestMetrics.Counter(r.URL.Path, http.StatusBadRequest).Inc()
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	w.WriteHeader(http.StatusOK)

	metrics.RestMetrics.Counter(r.URL.Path, http.StatusOK).Inc()
}
