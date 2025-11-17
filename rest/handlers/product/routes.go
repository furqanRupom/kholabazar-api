package product

import (
	"kholabazar/rest/middlewares"
	"net/http"
)

func (h *Handler) ProductRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("GET /products",
		manager.With(
			http.HandlerFunc(h.GetProducts),
		),
	)

	mux.Handle("POST /products",
		manager.With(
			http.HandlerFunc(h.CreateProduct),
			h.middlewares.AuthJWT,
		),
	)

	mux.Handle("GET /products/{id}",
		manager.With(
			http.HandlerFunc(h.GetProduct),
		),
	)
	mux.Handle("PUT /products/{id}",
		manager.With(
			http.HandlerFunc(h.UpdateProduct),
			h.middlewares.AuthJWT,
		),
	)
	mux.Handle("DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(h.DeleteProduct),
			h.middlewares.AuthJWT,
		),
	)

}
