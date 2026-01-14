package middleware

import (
	"log"
	"net/http"
	"runtime/debug"
)

func Recovery(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if panicError := recover(); panicError != nil {
				log.Printf("Panic: %v\n%s", panicError, debug.Stack())
				http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		nextHandler.ServeHTTP(writer, request)
	})
}
