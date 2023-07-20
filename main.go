package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
)

type Server struct{}
type Todo struct {
    UserId      int
    Id          int
    Title       string
    Completed   bool
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(&Todo{UserId: 1, Id: 1, Title: "delectus aut autem", Completed: false})
}

func main() {
    port := 8080
    fmt.Printf("Starting HTTP server on port %d", port)
    s := &Server{}
    http.Handle("/", s)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil))
}
