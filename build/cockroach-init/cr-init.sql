CREATE DATABASE IF NOT EXISTS reader;

CREATE TABLE IF NOT EXISTS reader.story (
  Id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  Url STRING UNIQUE NOT NULL,
  Title STRING,
  Author STRING,
  NumChapters INT,
  INDEX url_idx (Url) STORING (Id)
);
CREATE USER IF NOT EXISTS story_service;
GRANT ALL ON DATABASE reader TO story_service;
GRANT ALL ON reader.story TO story_service;

CREATE TABLE IF NOT EXISTS reader.chapter (
  Story UUID NOT NULL,
  Id INT NOT NULL,
  Url STRING,
  Title STRING,
  Html STRING,
  PRIMARY KEY (Story, Id),
  CONSTRAINT fk_story FOREIGN KEY (Story) REFERENCES reader.story
) INTERLEAVE IN PARENT reader.story (Story);
CREATE USER IF NOT EXISTS chapter_service;
GRANT ALL ON DATABASE reader to chapter_service;
GRANT ALL ON reader.chapter TO chapter_service;
GRANT SELECT ON reader.story TO chapter_service;