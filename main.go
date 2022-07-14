package main

import (
	"github.com/go-chi/chi"
	"github.com/salad-server/proxy/routes"
	"github.com/salad-server/proxy/util"
)

func main() {
	cfg := util.LoadCfg()
	router := routes.Router{
		Users: cfg.Accounts,
		Port:  cfg.Port,
		Mux:   chi.NewRouter(),
	}

	router.HandleBancho(cfg.Proxy)
	router.Serve()
}
