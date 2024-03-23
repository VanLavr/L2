package schedulertransport

import (
	"net/http"

	"github.com/VanLavr/L2/develop/11/internal/scheduler"
)

type Server struct {
	addr string
	mux  *http.ServeMux
	eh   *eventHandler
}

func New(addr string, usecase scheduler.Scheduler) *Server {
	handler := newEventHandler(usecase)

	return &Server{addr: addr, eh: handler, mux: http.NewServeMux()}
}

func (s Server) ListenAndServe() error {
	return http.ListenAndServe(s.addr, s.mux)
}
