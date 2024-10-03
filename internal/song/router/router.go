package router

import (
	"github.com/CracherX/Song-service/internal/song/endpoints"
	"github.com/gorilla/mux"
)

// Setup устанавливает главный роутер.
func Setup() *mux.Router {
	r := mux.NewRouter()
	return r
}

// Songs устанавливает саброутер для работы с библиотекой песен для основного роутера.
func Songs(mr *mux.Router, ep *endpoints.SongsEndpoint) *mux.Router {
	r := mr.PathPrefix("/songs").Subrouter()
	r.HandleFunc("", ep.Library).Methods("GET")
	r.HandleFunc("/lyrics/{id}", ep.Text).Methods("GET")
	r.HandleFunc("/{id}", ep.Delete).Methods("DELETE")
	r.HandleFunc("/{id}", ep.UpdateSong).Methods("PATCH")
	r.HandleFunc("/add", ep.AddSong).Methods("POST")
	return r
}
