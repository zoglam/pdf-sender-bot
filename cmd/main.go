package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/zoglam/pdf-sender-bot/internal/app"
	"github.com/zoglam/pdf-sender-bot/internal/service"
	"golang.org/x/sys/unix"
)

func main() {

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, unix.SIGKILL, unix.SIGTERM, unix.SIGQUIT, unix.SIGINT)
		sig := <-ch
		time.Sleep(1 * time.Second)
		log.Fatalf("Bot interrupted with signal: %d = %v", sig, sig)
	}()

	pdfService := service.NewPDFService()
	app := app.NewApp(pdfService)

	r := mux.NewRouter()
	r.HandleFunc("/get", app.GetPDF).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Starting server on 0.0.0.0:8000")
	log.Fatal(srv.ListenAndServe())
}
