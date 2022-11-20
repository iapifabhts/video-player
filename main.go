package main

import (
	"github.com/gorilla/mux"
	h "github.com/iapifabhts/video-player/handlers"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/files/{_}/{_}", h.Middleware(h.Get)).Methods(http.MethodGet)
	r.HandleFunc("/files", h.Middleware(h.Upload)).Methods(http.MethodPost)
	r.HandleFunc("/items", h.Middleware(h.GetAll)).Methods(http.MethodGet)
	r.HandleFunc("/items", h.Middleware(h.UploadItem)).Methods(http.MethodPost, http.MethodOptions)
	r.Use(mux.CORSMethodMiddleware(r))
	
	log.Fatal(http.ListenAndServe(":80", r))
}
