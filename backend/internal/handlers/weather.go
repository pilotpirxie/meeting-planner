package handlers

import "net/http"

func GetWeather(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, map[string]any{
		"defaultLocation": "san-francisco",
		"forecast": []map[string]any{
			{"day": "today", "condition": "sunny", "highC": 21, "lowC": 14},
			{"day": "tomorrow", "condition": "partly-cloudy", "highC": 19, "lowC": 12},
		},
	})
}

func GetWeatherByLocation(w http.ResponseWriter, r *http.Request) {
	location := r.PathValue("location")
	RespondJSON(w, http.StatusOK, map[string]any{
		"location": location,
		"forecast": []map[string]any{
			{"day": "today", "condition": "cloudy", "highC": 18, "lowC": 11},
			{"day": "tomorrow", "condition": "rain", "highC": 16, "lowC": 10},
		},
	})
}
