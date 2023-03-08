package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/VictoriaMetrics/metrics"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/zoglam/pdf-sender-bot/internal/app"
	"github.com/zoglam/pdf-sender-bot/internal/app/rest"
	"github.com/zoglam/pdf-sender-bot/internal/repository"
	"github.com/zoglam/pdf-sender-bot/internal/service"
	"github.com/zoglam/pdf-sender-bot/pkg/logg"
	"github.com/zoglam/pdf-sender-bot/pkg/telegram"
	"golang.org/x/sys/unix"
)

func main() {
	ctx := context.Background()

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, unix.SIGKILL, unix.SIGTERM, unix.SIGQUIT, unix.SIGINT)
		sig := <-ch
		time.Sleep(1 * time.Second)
		log.Fatalf("Bot interrupted with signal: %d = %v", sig, sig)
	}()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		logg.Fatal().Msg("cannot read from a config. " + err.Error())
	}

	viper.SetDefault("server.port", "8000")

	serverHost := viper.Get("server.port").(string)
	telegramToken := viper.Get("telegram.token").(string)
	telegramWebAppURL := viper.Get("telegram.webappurl").(string)
	dsn := viper.Get("db.dsn").(string)

	db, err := repository.NewDB(ctx, dsn)
	if err != nil {
		logg.Fatal().Msgf("cannot connect db: %v", err)
	}
	err = db.Ping(ctx)
	if err != nil {
		logg.Fatal().Msgf("cannot ping db: %v", err)
	}
	defer db.Close()

	dao := repository.NewDAO()
	bot := telegram.InitBot(telegramToken)
	userService := service.NewUserService(dao)
	pdfService := service.NewPDFService()

	rest := rest.NewRest(
		pdfService,
		userService,
	)
	app := app.NewApp(
		pdfService,
		bot,
		userService,
	)

	bot.RegisterWebAppURL(telegramWebAppURL)
	bot.RegisterMiddleWareMetrics(app.MiddleWareMetrics)
	bot.RegisterHandler("/start", app.GetStart)
	bot.RegisterHandler("/pdf", app.GetPDF)

	go bot.Bot.Start()
	defer bot.Bot.Close()

	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.Use(rest.Middleware)
	api.HandleFunc("/profile", rest.PostProfile).Methods("POST")
	api.HandleFunc("/profile", rest.GetProfile).Methods("GET")

	private := r.PathPrefix("/private").Subrouter()
	private.Use(rest.Middleware)
	private.HandleFunc("/getpdf", rest.GetPDF).Methods("GET")

	r.HandleFunc("/metrics", func(w http.ResponseWriter, _ *http.Request) {
		metrics.WritePrometheus(w, true)
	}).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:" + serverHost,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Starting server on 0.0.0.0:" + serverHost)
	log.Fatal(srv.ListenAndServe())
}
