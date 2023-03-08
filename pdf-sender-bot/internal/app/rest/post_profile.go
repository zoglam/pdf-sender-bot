package rest

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/zoglam/pdf-sender-bot/internal/dto"
	"github.com/zoglam/pdf-sender-bot/pkg/logg"
	"github.com/zoglam/pdf-sender-bot/pkg/metrics"
)

func (a *Rest) PostProfile(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	ctx = context.WithValue(ctx, "x-request-id", r.Header.Get("X-Request-Id"))
	ctx = context.WithValue(ctx, "x-telegram-id", r.Header.Get("X-Telegram-Id"))

	logg.Info(ctx).Msgf("Пришел запрос на post")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logg.Error(ctx).Msgf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		metrics.RestMetrics.Counter(r.URL.Path, http.StatusBadRequest).Inc()
		return
	}

	var user *dto.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		logg.Error(ctx).Msgf("Error Unmarshal body: %v", err)
		http.Error(w, "can't Unmarshal body", http.StatusBadRequest)
		metrics.RestMetrics.Counter(r.URL.Path, http.StatusBadRequest).Inc()
		return
	}

	id, err := strconv.ParseInt(r.Header.Get("X-Telegram-Id"), 10, 64)
	if err != nil {
		logg.Error(ctx).Msgf("Error ParseInt id: %v", err)
		http.Error(w, "Error ParseInt id", http.StatusBadRequest)
		metrics.RestMetrics.Counter(r.URL.Path, http.StatusBadRequest).Inc()
		return
	}

	user.UserID = id
	err = a.userService.SaveUserProfile(id, user)
	if err != nil {
		logg.Error(ctx).Msgf("Error SaveUserProfile: %v", err)
		http.Error(w, "Error SaveUserProfile", http.StatusBadRequest)
		metrics.RestMetrics.Counter(r.URL.Path, http.StatusBadRequest).Inc()
		return
	}

	logg.Info(ctx).Msgf("Пришел запрос на post %+v", user)

	w.Header().Set("Content-Type", "plain/text")
	w.WriteHeader(http.StatusOK)

	metrics.RestMetrics.Counter(r.URL.Path, http.StatusOK).Inc()
}
