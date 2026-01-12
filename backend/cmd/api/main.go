package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"runtime"
	"strings"
	"syscall"
	"time"

	"meeting-planner/backend/internal/db"
	"meeting-planner/backend/internal/handlers"
	"meeting-planner/backend/internal/middleware"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	ctx := context.Background()
	if err := db.Init(ctx); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/health", handlers.Healthcheck)
	mux.HandleFunc("POST /api/echo/{id}", handlers.Echo)

	mux.HandleFunc("GET /api/calendars", handlers.ListCalendars)
	mux.HandleFunc("POST /api/calendars", handlers.CreateCalendar)
	mux.HandleFunc("GET /api/calendars/{id}", handlers.GetCalendar)
	mux.HandleFunc("POST /api/calendars/{id}/vote", handlers.VoteCalendar)

	// TODO: Implement polls
	// mux.HandleFunc("GET /api/polls", handlers.ListPolls)
	// mux.HandleFunc("POST /api/polls", handlers.CreatePoll)
	// mux.HandleFunc("GET /api/polls/{id}", handlers.GetPoll)
	// mux.HandleFunc("POST /api/polls/{id}/vote", handlers.VotePoll)

	// TODO: Implement weather endpoints
	// mux.HandleFunc("GET /api/weather", handlers.GetWeather)
	// mux.HandleFunc("GET /api/weather/{location}", handlers.GetWeatherByLocation)

	staticDir := "public"
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodHead {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		if strings.HasPrefix(r.URL.Path, "/api/") {
			http.NotFound(w, r)
			return
		}
		requested := path.Clean(r.URL.Path)
		if requested == "/" {
			requested = "/index.html"
		}
		filePath := path.Join(staticDir, requested)
		if _, err := os.Stat(filePath); err == nil {
			http.ServeFile(w, r, filePath)
			return
		}
		indexPath := path.Join(staticDir, "index.html")
		if _, err := os.Stat(indexPath); err == nil {
			http.ServeFile(w, r, indexPath)
			return
		}
		http.NotFound(w, r)
	}))

	handler := middleware.Recovery(middleware.Logging(middleware.CORS(mux)))

	addr := ":8080"
	port := os.Getenv("PORT")
	if port != "" {
		addr = ":" + port
	}

	srv := &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	go func() {
		startupInfo := map[string]any{
			"message": "Server is starting",
			"addr":    srv.Addr,
			"time":    time.Now().Format(time.RFC3339),
			"mode":    "development",
			"sdk":     runtime.Version(),
		}

		log.Printf("%s", handlers.ToJSONPretty(startupInfo))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
