CREATE DATABASE IF NOT EXISTS reader;
CREATE TABLE IF NOT EXISTS reader.story (
  Id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  Url STRING UNIQUE NOT NULL,
  Title STRING,
  Author STRING,
  NumChapters INT,
  INDEX url_idx (Url) STORING (Id)
);
CREATE USER story_service;
GRANT ALL ON DATABASE reader TO story_service;
GRANT ALL ON reader.story TO story_service;