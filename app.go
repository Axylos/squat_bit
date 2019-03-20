package main

import(
  "fmt"
  "time"
  "net/http"
  "net/http/httputil"
  "log"
  "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL)
    fmt.Println(r.Host)
    fmt.Println(r.UserAgent())
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

    f.WriteString("Entry\n")
    now := time.Now().Format(time.RFC850)
    f.WriteString(now)
    f.WriteString("\n")
    if _, err = f.Write(dump); err != nil {
      panic(err)
    }

    f.WriteString("\n")
    f.WriteString("\n")

    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("500 - Not Found"))
  }

func main() {
  fmt.Println("hey there")
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":7000", nil))
}
