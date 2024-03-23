package schedulertransport

import (
	"net/http"

	"github.com/VanLavr/L2/develop/11/internal/pkg/middlewares"
	"github.com/VanLavr/L2/develop/11/internal/scheduler"
)

type Server struct {
	addr string
	mux  *http.ServeMux
	eh   *eventHandler
}

func New(addr string, usecase scheduler.Scheduler, validator *middlewares.EventValidator) *Server {
	handler := newEventHandler(usecase, validator)

	return &Server{addr: addr, eh: handler, mux: http.NewServeMux()}
}

func (s Server) ListenAndServe() error {
	return http.ListenAndServe(s.addr, s.mux)
}
