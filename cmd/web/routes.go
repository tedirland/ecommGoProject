package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *applicaiton) routes() http.Handler {
	mux := chi.NewRouter()
	return mux
}
