package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/riesinger/preseeder/api/routes"
)

type HTTPServer struct {
	addr string
	mux  *mux.Router
}

func NewHTTPServer(addr string, renderer routes.Renderer) *HTTPServer {
	srv := &HTTPServer{
		addr: addr,
		mux:  mux.NewRouter(),
	}
	srv.mux.StrictSlash(false)
	srv.mux.HandleFunc(routes.GetPreseedHandler(renderer))

	return srv
}

func (s *HTTPServer) Start() {
	http.ListenAndServe(s.addr, s.mux)
}
