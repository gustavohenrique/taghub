PRAGMA auto_vacuum = 1;

DROP TABLE IF EXISTS repos;
CREATE TABLE repos (
  id TEXT NOT NULL PRIMARY KEY,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  name TEXT NOT NULL,
  description TEXT NOT NULL DEFAULT '',
  url TEXT NOT NULL DEFAULT '',
  homepage TEXT NOT NULL DEFAULT ''
);

CREATE INDEX idx_repos_name ON repos(name);


DROP TABLE IF EXISTS tags;
CREATE TABLE tags (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL
);

CREATE UNIQUE INDEX idx_name ON tags(name);


DROP TABLE IF EXISTS mapping;
CREATE TABLE mapping (
  repo_id TEXT NOT NULL,
  tag_id INTEGER,
  FOREIGN KEY(repo_id) REFERENCES repos(id) ON DELETE CASCADE,
  FOREIGN KEY(tag_id) REFERENCES tags(id) ON DELETE CASCADE
);

CREATE UNIQUE INDEX idx_mapping ON mapping(repo_id, tag_id);
