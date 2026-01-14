package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handler) Healthcheck(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, map[string]any{"status": "ok"})
}

func (h *Handler) Echo(w http.ResponseWriter, r *http.Request) {
	var payloadBody struct {
		Message string `json:"message" validate:"required,min=1"`
	}

	var payloadQuery struct {
		Name string `query:"name" validate:"required"`
		Age  int    `query:"age" validate:"gte=18,lte=99"`
	}

	var payloadParams struct {
		ID string `param:"id" validate:"required"`
	}

	var payloadHeaders struct {
		AuthToken string `header:"Authorization" validate:"required"`
	}

	err := ParseRequest(r, RequestOptions{
		Body:    &payloadBody,
		Params:  &payloadParams,
		Query:   &payloadQuery,
		Headers: &payloadHeaders,
	})

	if err != nil {
		fmt.Printf("Error parsing request: %v\n", err)
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, fmt.Sprintf("Body: %s, Params: %s, Query: %s, Headers: %s", ToJSON(payloadBody), ToJSON(payloadParams), ToJSON(payloadQuery), ToJSON(payloadHeaders)))
}
