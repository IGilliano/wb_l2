package handler

import (
	"dev11/pkg/models"
	"dev11/pkg/service"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"
)

/*
POST /create_event
POST /update_event
POST /delete_event
GET /events_for_day
GET /events_for_week
GET /events_for_month
*/

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/create_event", h.createEvent)
	mux.HandleFunc("/update_event", h.updateEvent)
	mux.HandleFunc("/delete_event", h.deleteEvent)
	mux.HandleFunc("/events_for_day", h.getEventsForDay)
	mux.HandleFunc("/events_for_week", h.getEventsForWeek)
	mux.HandleFunc("/events_for_month", h.getEventsForMonth)

	return mux
}

func (h *Handler) createEvent(w http.ResponseWriter, r *http.Request) {

	event, err := parseEvent(r)
	if err != nil {
		errorResponse(w, err.Error(), 400)
		return
	}

	if err = validateEvent(event); err != nil {
		errorResponse(w, err.Error(), 400)
		return
	}
	var result []models.Event
	event.EventId = h.service.Create(event)
	result = append(result, event)
	responseWrite(w, "Event created successfully", result)

}

func (h *Handler) updateEvent(w http.ResponseWriter, r *http.Request) {

	event, err := parseEvent(r)
	if err != nil {
		errorResponse(w, err.Error(), 400)
		return
	}

	if err = validateEvent(event); err != nil {
		errorResponse(w, err.Error(), 400)
		return
	}

	if err = h.service.Update(event); err != nil {
		errorResponse(w, err.Error(), 400)
		return
	}

	var result []models.Event
	result = append(result, event)
	responseWrite(w, "Event updated successfully", result)
}

func (h *Handler) deleteEvent(w http.ResponseWriter, r *http.Request) {
	event, err := parseEvent(r)
	if err != nil {
		errorResponse(w, err.Error(), 400)
		return
	}

	if event.EventId == 0 || event.UserId == 0 {
		errorResponse(w, "user_id and event_id cannot be empty", 400)
		return
	}

	if err = h.service.Delete(event.UserId, event.EventId); err != nil {
		errorResponse(w, err.Error(), 400)
		return
	}

	responseWrite(w, "Event deleted successfully", nil)

}

func (h *Handler) getEventsForDay(w http.ResponseWriter, r *http.Request) {
	var result []models.Event
	event, err := parseEvent(r)
	if err != nil {
		errorResponse(w, err.Error(), 400)
		return
	}

	if result, err = h.service.Get(1, event.UserId, event.Date); err != nil {
		errorResponse(w, err.Error(), 400)
		return
	}
	responseWrite(w, "Events: ", result)
}

func (h *Handler) getEventsForWeek(w http.ResponseWriter, r *http.Request) {
	event, err := parseEvent(r)
	if err != nil {
		errorResponse(w, err.Error(), 400)
	} else if result, err := h.service.Get(7, event.UserId, event.Date); err != nil {
		errorResponse(w, err.Error(), 400)
	} else {
		responseWrite(w, "Events: ", result)
	}
}

func (h *Handler) getEventsForMonth(w http.ResponseWriter, r *http.Request) {
	var result []models.Event
	event, err := parseEvent(r)
	if err != nil {
		errorResponse(w, err.Error(), 400)
		return
	}

	if result, err = h.service.Get(31, event.UserId, event.Date); err != nil {
		errorResponse(w, err.Error(), 400)
		return
	}

	responseWrite(w, "Events: ", result)
}

func parseEvent(r *http.Request) (models.Event, error) {
	var event models.Event
	var err error
	uIdStr := r.FormValue("user_id")
	if uIdStr != "" {
		event.UserId, err = strconv.Atoi(uIdStr)
		if err != nil {
			return event, err
		}
	}
	event.Date, err = time.Parse(time.DateOnly, r.FormValue("date"))
	if err != nil {
		return event, err
	}
	event.Description = r.FormValue("description")

	eIdStr := r.FormValue("event_id")
	if eIdStr != "" {
		event.EventId, err = strconv.Atoi(eIdStr)
		if err != nil {
			return event, err
		}
	}

	return event, nil
}

func validateEvent(event models.Event) error {
	if event.UserId < 0 {
		return errors.New("incorrect ID")
	}

	if event.Date.IsZero() {
		return errors.New("incorrect date")
	}

	return nil
}

func errorResponse(w http.ResponseWriter, error string, sc int) {
	errorResponse := struct {
		Error string `json:"error"`
	}{Error: error}

	js, err := json.Marshal(errorResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(sc)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func responseWrite(w http.ResponseWriter, result string, events []models.Event) {
	resultResponse := struct {
		Result string         `json:"result"`
		Events []models.Event `json:"events"`
	}{Result: result, Events: events}

	js, err := json.Marshal(resultResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
