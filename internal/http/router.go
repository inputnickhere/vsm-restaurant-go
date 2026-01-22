package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/inputnickhere/vsm-restaurant-go/internal/ingredients"
	"github.com/inputnickhere/vsm-restaurant-go/internal/menu"
)

func NewRouter(pool *pgxpool.Pool, adminToken string, supplierToken string) http.Handler {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		if err := pool.Ping(r.Context()); err != nil {
			http.Error(w, "db not ready", http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	// menu wiring
	menuRepo := menu.NewRepo(pool)
	menuSvc := menu.NewService(menuRepo)
	menuH := menu.NewHandler(menuSvc)

	// ingredients wiring
	ingRepo := ingredients.NewRepo(pool)
	ingSvc := ingredients.NewService(ingRepo)
	ingH := ingredients.NewHandler(ingSvc)

	// public API
	r.Route("/api/public", func(r chi.Router) {
		menuH.RegisterPublic(r)
	})

	// admin API (protected)
	r.Route("/api/admin", func(r chi.Router) {
		r.Use(StaticBearerAuth(adminToken))
		menuH.RegisterAdmin(r)
		ingH.RegisterAdmin(r)
	})

	// supplier API (protected)
	r.Route("/api/supplier", func(r chi.Router) {
		r.Use(StaticBearerAuth(supplierToken))
		ingH.RegisterSupplier(r)
	})

	return r
}
