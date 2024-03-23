package schedulertransport

import "github.com/VanLavr/L2/develop/11/internal/pkg/middlewares"

func (s *Server) RegisterRoutes() {
	logger := middlewares.NewLogger()

	s.mux.HandleFunc("POST /create_event", logger.LogScheduleAction(s.eh.HandleCreateEvent))
	s.mux.HandleFunc("POST /update_event", logger.LogScheduleAction(s.eh.HandleUpdateEvent))
	s.mux.HandleFunc("POST /delete_event", logger.LogScheduleAction(s.eh.HandleDeleteEvent))
	s.mux.HandleFunc("GET /events_for_day/{day}", logger.LogScheduleAction(s.eh.HandleGetForDay))
	s.mux.HandleFunc("GET /events_for_week/{dateBeginning}", logger.LogScheduleAction(s.eh.HandleGetForWeek))
	s.mux.HandleFunc("GET /events_for_month/{dateBeginning}", logger.LogScheduleAction(s.eh.HandleGetForMonth))
}
