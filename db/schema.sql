DROP TABLE IF EXISTS requests;

CREATE TABLE requests (
  id SERIAL PRIMARY KEY,
  created_at timestamp DEFAULT now(),
  url TEXT,
  host TEXT,
  remote_addr TEXT,
  protocol VARCHAR(255),
  req_uri TEXT,
  method VARCHAR(20),
  content_length INTEGER,
  raw_request TEXT
);
