package handlers

import "net/http"

func ListCalendars(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, map[string]any{
		"calendars": []map[string]any{
			{
				"id":          "cal-123",
				"title":       "Team Sync",
				"description": "Pick a date for the next team sync",
				"dates":       []string{"2025-01-05", "2025-01-06", "2025-01-08"},
			},
		},
	})
}

func CreateCalendar(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Dates       []string `json:"dates"`
	}

	if err := ParseJSON(r, &req); err != nil {
		RespondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.Title == "" || len(req.Dates) == 0 {
		RespondError(w, http.StatusBadRequest, "title and dates are required")
		return
	}

	RespondJSON(w, http.StatusCreated, map[string]any{
		"id":          "cal-new",
		"title":       req.Title,
		"description": req.Description,
		"dates":       req.Dates,
		"message":     "calendar created (mocked)",
	})
}

func GetCalendar(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	RespondJSON(w, http.StatusOK, map[string]any{
		"id":          id,
		"title":       "Team Sync",
		"description": "Pick a date for the next team sync",
		"dates":       []string{"2025-01-05", "2025-01-06", "2025-01-08"},
		"votes": []map[string]any{
			{"user": "alice", "available": []string{"2025-01-05", "2025-01-08"}},
			{"user": "bob", "available": []string{"2025-01-06"}},
		},
	})
}

func VoteCalendar(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var req struct {
		User      string   `json:"user"`
		Available []string `json:"available"`
	}

	if err := ParseJSON(r, &req); err != nil {
		RespondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.User == "" || len(req.Available) == 0 {
		RespondError(w, http.StatusBadRequest, "user and available dates are required")
		return
	}

	RespondJSON(w, http.StatusCreated, map[string]any{
		"id":        id,
		"user":      req.User,
		"available": req.Available,
		"message":   "vote recorded (mocked)",
	})
}
