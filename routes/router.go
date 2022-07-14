package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type Router struct {
	Users [][2]string
	Port  int
	Mux   *chi.Mux
}

func (route *Router) HandleBancho(paths []string) {
	for _, path := range paths {
		log.Println("[bancho]", path)
		route.Mux.Get(path, route.Bancho)
	}
}

func (route *Router) Serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", route.Port),
		Handler:           route.Mux,
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	log.Printf("Running on :%d", route.Port)
	return srv.ListenAndServe()
}
