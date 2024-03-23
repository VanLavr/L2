package schedulertransport

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/VanLavr/L2/develop/11/internal/models"
	"github.com/VanLavr/L2/develop/11/internal/pkg/errors"
	"github.com/VanLavr/L2/develop/11/internal/pkg/middlewares"
	"github.com/VanLavr/L2/develop/11/internal/scheduler"
)

type eventHandler struct {
	scheduler scheduler.Scheduler
	validator *middlewares.EventValidator
}

func newEventHandler(usecase scheduler.Scheduler, validator *middlewares.EventValidator) *eventHandler {
	return &eventHandler{scheduler: usecase, validator: validator}
}

func (e *eventHandler) HandleCreateEvent(w http.ResponseWriter, r *http.Request) {
	event, err := e.unmarshalBody(r)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(content)
		return
	}

	modelEvent, err := e.validateEvent(event)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(content)
		return
	}

	id, err := e.scheduler.CreateEvent(modelEvent)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(content)
		return
	}

	conent, err := e.marshallPostResponse(id)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(content)
		return
	}

	w.Write(conent)
}

func (e *eventHandler) HandleUpdateEvent(w http.ResponseWriter, r *http.Request) {
	event, err := e.unmarshalBody(r)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(content)
		return
	}

	modelEvent, err := e.validateEvent(event)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(content)
		return
	}

	id, err := e.scheduler.UpdateEvent(modelEvent)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(content)
		return
	}

	conent, err := e.marshallPostResponse(id)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(content)
		return
	}

	w.Write(conent)
}

func (e *eventHandler) HandleDeleteEvent(w http.ResponseWriter, r *http.Request) {
	event, err := e.unmarshalBody(r)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(content)
		return
	}

	id, err := e.scheduler.DeleteEvent(event.ID)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(content)
		return
	}

	conent, err := e.marshallPostResponse(id)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(content)
		return
	}

	w.Write(conent)
}

func (e *eventHandler) HandleGetForDay(w http.ResponseWriter, r *http.Request) {
	day := r.PathValue("day")

	date, err := e.validator.ValidateDate(day)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(content)
		return
	}

	events := e.scheduler.FetchForDay(date)
	content, err := e.marshallGetResponse(events)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(content)
		return
	}

	w.Write(content)
}

func (e *eventHandler) HandleGetForWeek(w http.ResponseWriter, r *http.Request) {
	day := r.PathValue("dateBeginning")

	date, err := e.validator.ValidateDate(day)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(content)
		return
	}

	events := e.scheduler.FetchForWeek(date)
	content, err := e.marshallGetResponse(events)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(content)
		return
	}

	w.Write(content)
}

func (e *eventHandler) HandleGetForMonth(w http.ResponseWriter, r *http.Request) {
	day := r.PathValue("dateBeginning")

	date, err := e.validator.ValidateDate(day)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(content)
		return
	}

	events := e.scheduler.FetchForMonth(date)
	content, err := e.marshallGetResponse(events)
	if err != nil {
		content := e.marshallErrorResponse(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(content)
		return
	}

	w.Write(content)
}

type dto struct {
	ID   int    `json:"event_id"`
	Date string `json:"event_date"`
	Name string `json:"event_name"`
}

type errorResponse struct {
	Content string `json:"error"`
}

type response struct {
	Content any `json:"result"`
}

func (e *eventHandler) unmarshalBody(r *http.Request) (dto, error) {
	var event dto
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		return dto{}, errors.ErrInvalidData
	}

	return event, nil
}

func (e *eventHandler) marshallGetResponse(resp []models.Event) ([]byte, error) {
	r, err := json.Marshal(response{Content: resp})
	if err != nil {
		return nil, errors.ErrInternal
	}

	return r, nil
}

func (e *eventHandler) marshallPostResponse(id int) ([]byte, error) {
	resp, err := json.Marshal(response{Content: id})
	if err != nil {
		return nil, errors.ErrInternal
	}

	return resp, nil
}

func (e *eventHandler) marshallErrorResponse(err error) []byte {
	resp, err := json.Marshal(errorResponse{Content: err.Error()})
	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func (e *eventHandler) validateEvent(event dto) (models.Event, error) {
	err := e.validator.ValidateEventName(event.Name)
	if err != nil {
		return models.Event{}, err
	}

	date, err := e.validator.ValidateDate(event.Date)
	if err != nil {
		return models.Event{}, err
	}

	model := models.Event{ID: event.ID, Name: event.Name, Date: date}
	return model, nil
}
