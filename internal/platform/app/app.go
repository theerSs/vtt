package app

import (
	"context"
	"fmt"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/theerSs/vtt/internal/platform/database"
	"github.com/theerSs/vtt/internal/platform/env"
	"github.com/theerSs/vtt/internal/platform/router"
	"github.com/theerSs/vtt/internal/platform/server"
	"github.com/theerSs/vtt/internal/rooms"
)

func Run(ctx context.Context) error {
	pool, err := database.Connect(ctx)
	if err != nil {
		return fmt.Errorf("database connect failed: %w", err)
	}
	defer pool.Close()

	handler := buildModules(pool)

	addr := fmt.Sprintf(":%s", env.APIPort.GetValue())
	srv := server.New(addr, handler)

	log.Printf("Server started on %s", addr)

	if err := server.Serve(ctx, srv); err != nil {
		return fmt.Errorf("server error: %w", err)
	}

	log.Println("Server stopped")
	return nil
}

func buildModules(pool *pgxpool.Pool) *chi.Mux {
	roomsModule := rooms.NewModule(pool)

	return router.InitRouter(router.Deps{
		Modules: map[string]router.AppModule{
			"rooms": roomsModule,
		},
	})
}
