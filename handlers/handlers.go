package handlers

import (
    "encoding/json"
    "log"
    "net/http"
)

type HealthCheck struct {
    Msg string
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    h := &HealthCheck{"ok"}
    body := make(map[string]interface{})

    body["msg"] = h.Msg

    b, err := json.Marshal(body)
    if err != nil {
        log.Panic(err)
    }

    w.Write(b)
}

func InitHandlers() {
    http.HandleFunc("/health", healthHandler)
}

