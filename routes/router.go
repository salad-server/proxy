package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/salad-server/proxy/util"
)

type Router struct {
	Users [][2]string
	Port  int
	Mux   *chi.Mux
}

func (router *Router) Index(path, fn string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		util.ParseTemplate(w, fn, len(router.Users))
	}
}

func (router *Router) HandleBancho(paths []string) {
	for _, path := range paths {
		log.Println("[bancho]", path)
		router.Mux.Get(path, router.Bancho)
	}
}

func (router *Router) HandleIndex(paths map[string]string) {
	for path, fn := range paths {
		log.Println(path, fn)
		router.Mux.Get(path, router.Index(path, fn))
	}

	router.Mux.Handle("/static/*", http.StripPrefix(
		"/static/",
		http.FileServer(http.Dir("static")),
	))
}

func (router *Router) Serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", router.Port),
		Handler:           router.Mux,
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	log.Printf("Running on :%d", router.Port)
	return srv.ListenAndServe()
}
