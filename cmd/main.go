package main

import (
	"fmt"
	"log"
	"net/http"

	authhttp "github.com/theerSs/vtt/internal/auth/http"
	"github.com/theerSs/vtt/internal/platform/env"
	"github.com/theerSs/vtt/internal/platform/router"
)

func main() {
	err := env.Load()
	if err != nil {
		log.Fatalf("env load failed: %v", err)
	}
	
	authHandler := authhttp.NewHandler()

	r := router.InitRouter(router.Deps{
		APIs: map[string]router.ApiHandler{
			"auth": authHandler,
		},
	})

	addr := fmt.Sprintf(":%s", env.APIPort.GetValue())
	log.Printf("Server started on %s", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}