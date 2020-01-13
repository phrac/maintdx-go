package assets

import (
	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{id}", GetAsset)
	router.Get("/", GetAllAssets)
	return router
}
