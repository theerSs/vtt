package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/theerSs/vtt/internal/platform/env"
	"github.com/theerSs/vtt/internal/platform/router"
	"github.com/theerSs/vtt/internal/rooms"
)

func main() {
	err := env.Load()
	if err != nil {
		log.Fatalf("env load failed: %v", err)
	}

	roomsModule := rooms.NewModule()

	r := router.InitRouter(router.Deps{
		Modules: map[string]router.AppModule{
			"rooms": roomsModule,
		},
	})

	addr := fmt.Sprintf(":%s", env.APIPort.GetValue())
	log.Printf("Server started on %s", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
