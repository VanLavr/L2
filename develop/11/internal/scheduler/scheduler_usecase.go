package scheduler

import (
	"time"

	"github.com/VanLavr/L2/develop/11/internal/models"
)

type Scheduler interface {
	CreateEvent(models.Event) (int, error)
	UpdateEvent(models.Event) (int, error)
	DeleteEvent(int) (int, error)
	FetchForDay(time.Time) []models.Event
	FetchForWeek(time.Time) []models.Event
	FetchForMonth(time.Time) []models.Event
}

type SchedulerRepository interface {
	CreateEvent(models.Event) (int, error)
	UpdateEvent(models.Event) (int, error)
	DeleteEvent(int) (int, error)
	Fetch(time.Time, time.Time) []models.Event
}

type Scheduler11 struct {
	repo SchedulerRepository
}

func New(repo SchedulerRepository) Scheduler {
	return &Scheduler11{repo: repo}
}

func (s *Scheduler11) CreateEvent(e models.Event) (int, error) {
	id, err := s.repo.CreateEvent(e)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Scheduler11) UpdateEvent(e models.Event) (int, error) {
	id, err := s.repo.UpdateEvent(e)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Scheduler11) DeleteEvent(id int) (int, error) {
	id, err := s.repo.DeleteEvent(id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Scheduler11) FetchForDay(day time.Time) []models.Event {
	period := day.AddDate(0, 0, 1)
	return s.repo.Fetch(day, period)
}

func (s *Scheduler11) FetchForWeek(day time.Time) []models.Event {
	period := day.AddDate(0, 0, 7)
	return s.repo.Fetch(day, period)
}

func (s *Scheduler11) FetchForMonth(day time.Time) []models.Event {
	period := day.AddDate(0, 1, 0)
	return s.repo.Fetch(day, period)
}
