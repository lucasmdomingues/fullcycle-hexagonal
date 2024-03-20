package web

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/lucasmdomingues/hexagonal/adapters/web/handler"
	"github.com/lucasmdomingues/hexagonal/application"
)

type Server struct {
	Service application.ProductServiceInterface
}

func NewServer(service application.ProductServiceInterface) *Server {
	return &Server{
		Service: service,
	}
}

func (w *Server) Serve() {
	router := mux.NewRouter()
	middleware := negroni.New(negroni.NewLogger())

	handler.NewProductHandlers(router, middleware, w.Service)
	http.Handle("/", router)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log:", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
