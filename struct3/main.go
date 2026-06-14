package main

import (
    "log"
    "net/http"
)

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc(
        "GET /hello",
        func(w http.ResponseWriter, r *http.Request) {
           w.Header().Set("Content-Type", "application/json")
           w.WriteHeader(http.StatusOK)
           w.Write([]byte(`{"message": "hello world!"}`))
    })

    s := http.Server{
        Addr: ":8090",
        Handler: mux,
    }
    
    log.Fatal(s.ListenAndServe())
}
