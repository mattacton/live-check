package main

import (
  "log"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("ALIVE"))
}

func main() {
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":80", nil))
}
