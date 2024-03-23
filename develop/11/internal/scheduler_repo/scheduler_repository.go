package schedulerrepo

import (
	"sync"
	"time"

	"github.com/VanLavr/L2/develop/11/internal/models"
	"github.com/VanLavr/L2/develop/11/internal/pkg/errors"
	"github.com/VanLavr/L2/develop/11/internal/scheduler"
)

type repository struct {
	events    map[int]models.Event
	currentID int
	sync.RWMutex
}

func New() scheduler.SchedulerRepository {
	return &repository{events: make(map[int]models.Event), currentID: 0}
}

func (r *repository) CreateEvent(event models.Event) (int, error) {
	r.Lock()
	defer r.Unlock()

	r.currentID++

	event.ID = r.currentID
	r.events[r.currentID] = event

	return r.currentID, nil
}

func (r *repository) UpdateEvent(event models.Event) (int, error) {
	r.Lock()
	defer r.Unlock()

	_, found := r.events[event.ID]
	if !found {
		return 0, errors.ErrNotFound
	}

	r.events[event.ID] = event

	return event.ID, nil
}

func (r *repository) DeleteEvent(id int) (int, error) {
	r.Lock()
	defer r.Unlock()

	_, found := r.events[id]
	if !found {
		return 0, errors.ErrNotFound
	}

	delete(r.events, id)

	return id, nil
}

func (r *repository) Fetch(day time.Time, period time.Time) []models.Event {
	r.RLock()
	defer r.RUnlock()

	var result []models.Event

	for _, event := range r.events {
		if event.Date.UnixNano() >= day.UnixNano() && event.Date.UnixNano() <= period.UnixNano() {
			result = append(result, event)
		}
	}

	return result
}
