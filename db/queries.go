package db


import (
  "fmt"
  "log"
  "net/http"
  "database/sql"
  "net/http/httputil"
  _ "github.com/lib/pq"
  "os"
)

func (conn *SquatConn) SaveReq(req *http.Request) (int64, error) {
  stmt, err := conn.Conn.Prepare(`
  INSERT INTO requests
  (url, host, remote_addr, protocol, method, req_uri, content_length, raw_request)
  VALUES
  ($1, $2, $3, $4, $5, $6, $7, $8)
  ` )

  if err != nil {
    log.Fatal(err)
    return 0, err
  }

  dump, err := httputil.DumpRequest(req, true)

  _, stmtErr := stmt.Exec(
    req.URL.String(),
    req.Host,
    req.Header["Real-Ip"],
    req.Proto,
    req.Method,
    req.RequestURI,
    req.ContentLength,
    dump,
  )

  if stmtErr != nil {
    log.Fatal(err)
    return 0, err
  }

  defer stmt.Close()

  return 0, nil
}

type SquatConn struct {
  Conn *sql.DB
}

func GetConn() *SquatConn {
  connStr := fmt.Sprintf(os.Getenv("DB_URL"))
  fmt.Println(connStr)
  fmt.Println("called")
  db, err := sql.Open("postgres", connStr)

  if err != nil {
    log.Fatal("first", err)
  }

  return &SquatConn{ db } 
}

func foo() {
  connStr := "user=wagnerizing dbname=squat_db sslmode=disable"
  db, err := sql.Open("postgres", connStr)

  if err != nil {
    log.Fatal("first", err)
  }

  //rows, err := db.Query("INSERT into requests  VALUES('bar')")

  rows, err := db.Query("SELECT url FROM requests")
  if err != nil {
    log.Fatal(err)
    fmt.Println("error: ", err)
  }

  defer rows.Close()

  for rows.Next() {
    var url string
    if err := rows.Scan(&url); err != nil {

      log.Fatal(err)
    }

    fmt.Println("url: ", url)
  }
  //fmt.Println("ok %q", rows)
}
