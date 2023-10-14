package web_server

import (
	"github.com/almeidacavalcante/ports-and-adapters/adapters/web_server/handler"
	"github.com/almeidacavalcante/ports-and-adapters/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	ProductService application.ProductServiceInterface
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Serve() {

	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(r, n, s.ProductService)
	http.Handle("/", r)

	server := &http.Server{
		Addr:              ":9000",
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
