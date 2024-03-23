package scheduler

import "github.com/VanLavr/L2/develop/11/internal/models"

type Scheduler interface {
	CreateEvent(models.Event) (int, error)
	UpdateEvent(models.Event) (int, error)
	DeleteEvent(int) (int, error)
	FetchForDay() []models.Event
	FetchForWeek() []models.Event
	FetchForMonth() []models.Event
}

type SchedulerRepository interface{}
