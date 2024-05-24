package transport

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/instinctG/exchanges/internal/service"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type RequestBody struct {
	Amount    int   `json:"amount"`
	Banknotes []int `json:"banknotes"`
}

type Handler struct {
	Router *mux.Router
	Server *http.Server
}

func NewHandler(host, port string) *Handler {
	h := &Handler{}

	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/exchange", h.PostExchange).Methods("POST")

	h.Server = &http.Server{
		Addr:    host + port,
		Handler: h.Router,
	}

	return h
}

func (h *Handler) PostExchange(w http.ResponseWriter, r *http.Request) {
	var result [][]int
	var reqBody RequestBody

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to decode", http.StatusBadRequest)
		return
	}

	service.Exchange(reqBody.Amount, reqBody.Banknotes, []int{}, 0, &result)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	logrus.Info("exchange completed")
}

func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			logrus.Info(err.Error())
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)
	logrus.Info("shut down gracefully")
	return nil
}
