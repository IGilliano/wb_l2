package service

import (
	"dev11/pkg/models"
	"dev11/pkg/repository"
	"time"
)

type Service struct {
	repo *repository.Repository
}

func NewService(r *repository.Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Create(event models.Event) int {
	return s.repo.CreateInCalendar(event)
}

func (s *Service) Update(ec models.Event) error {
	return s.repo.UpdateInCalendar(ec)
}

func (s *Service) Delete(userId, eventId int) error {
	return s.repo.DeleteFromCalendar(userId, eventId)
}

func (s *Service) Get(period, userId int, date time.Time) ([]models.Event, error) {
	return s.repo.GetFromCalendar(period, userId, date)
}
