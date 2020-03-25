PRAGMA auto_vacuum = 1;

DROP TABLE IF EXISTS repos;
CREATE TABLE repos (
  id TEXT NOT NULL PRIMARY KEY,
  created_at TIMESTAMP DEFAULT now,
  updated_at TIMESTAMP,
  name TEXT NOT NULL,
  description TEXT NOT NULL DEFAULT '',
  url TEXT NOT NULL DEFAULT '',
  homepage TEXT NOT NULL DEFAULT ''
);

CREATE INDEX idx_repos_name ON repos(name);
