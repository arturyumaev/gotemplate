package openapi

import (
	"encoding/json"
	"net/http"
	"time"
)

type statusResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Service   struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"service"`
	Failures map[string]string `json:"failures,omitempty"`
}

const (
	StatusOK                 = "OK"
	StatusPartiallyAvailable = "Partially Available"
	StatusUnavailable        = "Unavailable"
)

func Status(serviceName, serviceVersion string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		status := &statusResponse{
			Status:    StatusOK,
			Timestamp: time.Now(),
		}
		status.Service.Name = serviceName
		status.Service.Version = serviceVersion
		status.Failures = make(map[string]string)

		response, _ := json.Marshal(&status)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
