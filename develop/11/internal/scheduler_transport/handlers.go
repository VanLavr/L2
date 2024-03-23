package schedulertransport

import (
	"net/http"

	"github.com/VanLavr/L2/develop/11/internal/models"
	"github.com/VanLavr/L2/develop/11/internal/scheduler"
)

type eventHandler struct {
	scheduler.Scheduler
}

func newEventHandler(usecase scheduler.Scheduler) *eventHandler {
	return &eventHandler{Scheduler: usecase}
}

func (e *eventHandler) HandleCreateEvent(w http.ResponseWriter, r *http.Request) {
	event, err := e.unmarshalBody(r)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.Write(content)
		return
	}

	id, err := e.CreateEvent(event)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.Write(content)
		return
	}

	conent, err := e.marshallPostResponse(id)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.Write(content)
		return
	}

	w.Write(conent)
}

func (e *eventHandler) HandleUpdateEvent(w http.ResponseWriter, r *http.Request) {
	event, err := e.unmarshalBody(r)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.Write(content)
		return
	}

	id, err := e.UpdateEvent(event)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.Write(content)
		return
	}

	conent, err := e.marshallPostResponse(id)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.Write(content)
		return
	}

	w.Write(conent)
}

func (e *eventHandler) HandleDeleteEvent(w http.ResponseWriter, r *http.Request) {
	event, err := e.unmarshalBody(r)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.Write(content)
		return
	}

	id, err := e.DeleteEvent(event.ID)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.Write(content)
		return
	}

	conent, err := e.marshallPostResponse(id)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.Write(content)
		return
	}

	w.Write(conent)
}

func (e *eventHandler) HandleGetForDay(w http.ResponseWriter, r *http.Request) {
	events := e.FetchForDay()
	content, err := e.marshallGetResponse(events)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.Write(content)
		return
	}

	w.Write(content)
}

func (e *eventHandler) HandleGetForWeek(w http.ResponseWriter, r *http.Request) {
	events := e.FetchForWeek()
	content, err := e.marshallGetResponse(events)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.Write(content)
		return
	}

	w.Write(content)
}

func (e *eventHandler) HandleGetForMonth(w http.ResponseWriter, r *http.Request) {
	events := e.FetchForMonth()
	content, err := e.marshallGetResponse(events)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.Write(content)
		return
	}

	w.Write(content)
}

type errorResponse struct {
	Content string `json:"error"`
}

type response struct {
	Content any `json:"result"`
}

func (e *eventHandler) unmarshalBody(r *http.Request) (models.Event, error) {
	panic("implement me")
}

func (e *eventHandler) marshallGetResponse(response []models.Event) ([]byte, error) {
	panic("implement me")
}

func (e *eventHandler) marshallPostResponse(id int) ([]byte, error) {
	panic("implement me")
}

func (e *eventHandler) marshallErrorResponse(err error) []byte {
	panic("implement me")
}
