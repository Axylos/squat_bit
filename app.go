package main

import(
  "fmt"
  "net/http"
  "net/http/httputil"
  "log"
  "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    dump, err := httputil.DumpRequest(r, true)
    if err != nil {
      http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
      return
    }

    f, err := os.OpenFile("/var/logs.txt", os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
      panic(err)
    }

    defer f.Close()

    if _, err = f.Write(dump); err != nil {
      panic(err)
    }

    f.WriteString("\n")

    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("404 - Not Found"))
  }

func main() {
  fmt.Println("hey there")
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
