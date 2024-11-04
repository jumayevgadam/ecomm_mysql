package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

var r *chi.Mux

// RegisterRoutes is
func RegisterRoutes(handler *Handler) *chi.Mux {
	r = chi.NewRouter()

	r.Route("/products", func(r chi.Router) {
		r.Post("/", handler.CreateProduct)
		r.Get("/", handler.ListProducts)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.GetProduct)
			r.Patch("/", handler.UpdateProduct)
			r.Delete("/", handler.DeleteProduct)
		})
	})

	r.Route("/orders", func(r chi.Router) {
		r.Post("/", handler.CreateOrder)
		r.Get("/", handler.ListOrders)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.GetOrder)
			r.Delete("/", handler.DeleteOrder)
		})
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", handler.CreateUser)
		r.Get("/", handler.ListUsers)
		r.Patch("/", handler.UpdateUser)

		r.Route("/{id}", func(r chi.Router) {
			r.Delete("/", handler.DeleteUser)
		})

		r.Route("/login", func(r chi.Router) {
			r.Post("/", handler.LoginUser)
		})

		r.Route("/logout", func(r chi.Router) {
			r.Post("/", handler.LogOutUser)
		})

		r.Route("/tokens", func(r chi.Router) {
			r.Route("/renew", func(r chi.Router) {
				r.Post("/", handler.RenewAccessToken)
			})

			r.Route("/revoke/{id}", func(r chi.Router) {
				r.Post("/", handler.RevokeSession)
			})
		})

	})

	return r
}

func Start(addr string) error {
	return http.ListenAndServe(addr, r)
}
