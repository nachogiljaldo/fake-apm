package main

import (
    "os"
    "fmt"
    "log"
    "net/http"
    "strconv"
)

func handlerRoot(w http.ResponseWriter, r *http.Request) {
}

func handleIntakeV2(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusAccepted)
}

func handleIntakeV1(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusAccepted)
}

func main() {
    fmt.Printf("Parameter is %s", os.Args[1])
    port, err := strconv.Atoi(os.Args[1])
    if err != nil {
      fmt.Printf("Port being set to %d", 8443)
      port = 8443
    }
    http.HandleFunc("/", handlerRoot)
    http.HandleFunc("/healthcheck", handlerRoot)
    http.HandleFunc("/intake/v2/events", handleIntakeV2)
    http.HandleFunc("/intake/v2/rum/events", handleIntakeV2)
    http.HandleFunc("/v1/transactions", handleIntakeV1)
    http.HandleFunc("/v1/rum/transactions", handleIntakeV1)
    http.HandleFunc("/v1/client-side/transactions", handleIntakeV1)
    http.HandleFunc("/v1/errors", handleIntakeV1)
    http.HandleFunc("/v1/rum/errors", handleIntakeV1)
    http.HandleFunc("/v1/client-side/errors", handleIntakeV1)
    http.HandleFunc("/v1/rum/sourcemaps", handleIntakeV1)
    http.HandleFunc("/v1/client-side/sourcemaps", handleIntakeV1)
    log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%d", port), "/app/config/certs/node.crt", "/app/config/certs/node.key", nil))
}