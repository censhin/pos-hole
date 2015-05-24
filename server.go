package main

import (
    "net/http"

    "github.com/censhin/pos-hole/handlers"
)

func initServer() {
    handlers.InitHandlers()
    http.ListenAndServe(":8080", nil)
}

func main() {
    initServer()
}

