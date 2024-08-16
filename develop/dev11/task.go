package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      int       `json:"user_id"`
	Date        time.Time `json:"date"`
}

type Calendar struct {
	events map[int]Event
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}

func handleCreateEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	err := parseEvent(r, &event)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
		return
	}
	calendar.CreateEvent(event)
	writeJSON(w, http.StatusOK, map[string]string{"result": "event created"})
}

func handleUpdateEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	err := parseEvent(r, &event)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, `{"error": "invalid event id"}`, http.StatusBadRequest)
		return
	}

	err = calendar.UpdateEvent(id, event)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusServiceUnavailable)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"result": "event updated"})
}

func handleDeleteEvent(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, `{"error": "invalid event id"}`, http.StatusBadRequest)
		return
	}

	err = calendar.DeleteEvent(id)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusServiceUnavailable)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"result": "event deleted"})
}

func handleGetEventsForDay(w http.ResponseWriter, r *http.Request) {
	date, err := parseDate(r.FormValue("date"))
	if err != nil {
		http.Error(w, `{"error": "invalid date"}`, http.StatusBadRequest)
		return
	}
	events := calendar.GetEventsForDay(date)
	writeJSON(w, http.StatusOK, events)
}

func handleGetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	date, err := parseDate(r.FormValue("date"))
	if err != nil {
		http.Error(w, `{"error": "invalid date"}`, http.StatusBadRequest)
		return
	}
	events := calendar.GetEventsForWeek(date)
	writeJSON(w, http.StatusOK, events)
}

func handleGetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	date, err := parseDate(r.FormValue("date"))
	if err != nil {
		http.Error(w, `{"error": "invalid date"}`, http.StatusBadRequest)
		return
	}
	events := calendar.GetEventsForMonth(date)
	writeJSON(w, http.StatusOK, events)
}

func parseEvent(r *http.Request, event *Event) error {
	event.Title = r.FormValue("title")
	event.Description = r.FormValue("description")
	userID, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		return fmt.Errorf("invalid user_id")
	}
	event.UserID = userID

	date, err := parseDate(r.FormValue("date"))
	if err != nil {
		return fmt.Errorf("invalid date")
	}
	event.Date = date

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		return fmt.Errorf("invalid event id")
	}
	event.ID = id

	return nil
}

func parseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func NewCalendar() *Calendar {
	return &Calendar{events: make(map[int]Event)}
}

func (c *Calendar) CreateEvent(event Event) {
	c.events[event.ID] = event
}

func (c *Calendar) UpdateEvent(id int, event Event) error {
	if _, exists := c.events[id]; !exists {
		return fmt.Errorf("event with id %d does not exist", id)
	}
	c.events[id] = event
	return nil
}

func (c *Calendar) DeleteEvent(id int) error {
	if _, exists := c.events[id]; !exists {
		return fmt.Errorf("event with id %d does not exist", id)
	}
	delete(c.events, id)
	return nil
}

func (c *Calendar) GetEventsForDay(date time.Time) []Event {
	return c.getEventsByDateRange(date, date.AddDate(0, 0, 1))
}

func (c *Calendar) GetEventsForWeek(date time.Time) []Event {
	startOfWeek := date.Truncate(time.Hour * 24 * 7)
	endOfWeek := startOfWeek.AddDate(0, 0, 7)
	return c.getEventsByDateRange(startOfWeek, endOfWeek)
}

func (c *Calendar) GetEventsForMonth(date time.Time) []Event {
	startOfMonth := date.Truncate(time.Hour * 24 * 30)
	endOfMonth := startOfMonth.AddDate(0, 1, 0)
	return c.getEventsByDateRange(startOfMonth, endOfMonth)
}

func (c *Calendar) getEventsByDateRange(start, end time.Time) []Event {
	var events []Event
	for _, event := range c.events {
		if event.Date.After(start) && event.Date.Before(end) {
			events = append(events, event)
		}
	}
	return events
}

var calendar = NewCalendar()

func main() {
	http.HandleFunc("/create_event", handleCreateEvent)
	http.HandleFunc("/update_event", handleUpdateEvent)
	http.HandleFunc("/delete_event", handleDeleteEvent)
	http.HandleFunc("/events_for_day", handleGetEventsForDay)
	http.HandleFunc("/events_for_week", handleGetEventsForWeek)
	http.HandleFunc("/events_for_month", handleGetEventsForMonth)

	handler := loggingMiddleware(http.DefaultServeMux)

	log.Println("Запуск сервера на:8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("Невозможно запустить сервер: %v", err)
	}
}
