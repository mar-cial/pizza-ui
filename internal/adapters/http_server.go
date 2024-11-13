package adapters

import (
	"net/http"

	"github.com/mar-cial/pizza-ui/internal/ports"
)

type httpServer struct {
	server *http.Server
}

func (h *httpServer) Start() error {
	return h.server.ListenAndServe()
}

func NewHTTPServer(port string, handler http.Handler) ports.Server {
	return &httpServer{
		server: &http.Server{
			Addr:    port,
			Handler: handler,
		},
	}
}
