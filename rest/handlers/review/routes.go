package review

import (
	middleware "kholabazar/rest/middlewares"
	"net/http"
)

func (h *Handler) ReviewRoutes(mux *http.ServeMux, mngr *middleware.Manager) {
	mux.Handle("GET /reviews",
		mngr.With(
			http.HandlerFunc(h.GetReview),
		),
	)
}
