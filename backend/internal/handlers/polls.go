package handlers

import "net/http"

func ListPolls(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, map[string]any{
		"polls": []map[string]any{
			{
				"id":       "poll-42",
				"question": "Where should we meet?",
				"options":  []string{"Cafe", "Office", "Park"},
			},
		},
	})
}

func CreatePoll(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Question string   `json:"question"`
		Options  []string `json:"options"`
	}

	if err := ParseJSON(r, &req); err != nil {
		RespondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.Question == "" || len(req.Options) < 2 {
		RespondError(w, http.StatusBadRequest, "question and at least 2 options are required")
		return
	}

	RespondJSON(w, http.StatusCreated, map[string]any{
		"id":       "poll-new",
		"question": req.Question,
		"options":  req.Options,
		"message":  "poll created (mocked)",
	})
}

func GetPoll(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	RespondJSON(w, http.StatusOK, map[string]any{
		"id":       id,
		"question": "Where should we meet?",
		"options":  []string{"Cafe", "Office", "Park"},
		"votes": []map[string]any{
			{"user": "alice", "choice": "Cafe"},
			{"user": "bob", "choice": "Office"},
		},
	})
}

func VotePoll(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var req struct {
		User   string `json:"user"`
		Choice string `json:"choice"`
	}

	if err := ParseJSON(r, &req); err != nil {
		RespondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.User == "" || req.Choice == "" {
		RespondError(w, http.StatusBadRequest, "user and choice are required")
		return
	}

	RespondJSON(w, http.StatusCreated, map[string]any{
		"id":      id,
		"user":    req.User,
		"choice":  req.Choice,
		"message": "vote recorded (mocked)",
	})
}
