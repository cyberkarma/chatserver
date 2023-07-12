package server

import "github.com/go-chi/chi/v5"

type RouterBuilder struct {
}

func (b *RouterBuilder) Build() *chi.Mux {
	r := chi.NewRouter()

	return r
}
