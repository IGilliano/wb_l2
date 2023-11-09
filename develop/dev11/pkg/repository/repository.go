package repository

import (
	"dev11/pkg/models"
	"errors"
	"sync"
	"time"
)

type Repository struct {
	cache map[int][]models.Event
	mutex sync.RWMutex
}

func NewRepository() *Repository {
	return &Repository{cache: make(map[int][]models.Event), mutex: sync.RWMutex{}}
}

func (r *Repository) CreateInCalendar(event models.Event) int {
	event.EventId = len(r.cache[event.UserId]) + 1
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.cache[event.UserId] = append(r.cache[event.UserId], event)
	return event.EventId
}

func (r *Repository) UpdateInCalendar(event models.Event) error {
	for _, v := range r.cache[event.UserId] {
		if v.EventId == event.EventId {
			r.mutex.Lock()
			v = event
			r.mutex.Unlock()
			return nil
		}
	}
	return errors.New("event doesnt exist")
}

func (r *Repository) DeleteFromCalendar(userId, eventId int) error {
	for i, v := range r.cache[userId] {
		if v.EventId == eventId {
			r.mutex.Lock()
			r.cache[userId] = append(r.cache[userId][:i], r.cache[userId][i+1:]...)
			r.mutex.Unlock()
			return nil
		}
	}
	return errors.New("event doesnt exist")
}

func (r *Repository) GetFromCalendar(period, userId int, date time.Time) ([]models.Event, error) {
	if _, ok := r.cache[userId]; !ok {
		return nil, errors.New("user doesnt exist")
	}
	var result []models.Event
	switch period {
	case 1:
		r.mutex.Lock()
		defer r.mutex.Unlock()
		for i, v := range r.cache[userId] {
			if v.Date.Year() == date.Year() && v.Date.Month() == date.Month() && v.Date.Day() == date.Day() {
				result = append(result, r.cache[userId][i])
			}
		}
	case 7:
		r.mutex.Lock()
		defer r.mutex.Unlock()
		y, w := date.ISOWeek()
		for i, v := range r.cache[userId] {
			y1, w1 := v.Date.ISOWeek()
			if y == y1 && w == w1 {
				result = append(result, r.cache[userId][i])
			}
		}

	case 31:
		r.mutex.Lock()
		defer r.mutex.Unlock()
		for i, v := range r.cache[userId] {
			if v.Date.Year() == date.Year() && v.Date.Month() == date.Month() {
				result = append(result, r.cache[userId][i])
			}
		}
	}
	return result, nil

}
