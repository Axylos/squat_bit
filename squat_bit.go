package main

import(
  "fmt"
  "net/http"
  "log"
  "github.com/axylos/squat_bit/db"
)

var conn *db.SquatConn

func handler(w http.ResponseWriter, r *http.Request) {
    _, err := conn.SaveReq(r)
    if err != nil {
      log.Fatal(err)
    }

    if err != nil {
      http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
      return
    }

    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("404 - Not Found"))
  }

func main() {
  conn = db.GetConn()
  fmt.Println("hey there")
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
