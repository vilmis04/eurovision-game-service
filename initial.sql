
CREATE DATABASE IF NOT EXISTS ev_game;

CREATE TABLE IF NOT EXISTS admin_config (
    id SERIAL PRIMARY KEY,
    year INT NOT NULL,
    gameType VARCHAR(255) NOT NULL,
    isVotingActive BOOLEAN NOT NULL,
    votingEnd TIMESTAMPTZ DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS "group" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(20) NOT NULL,
  "owner" VARCHAR(50) NOT NULL,
  "members" text ARRAY NOT NULL,
  "dateCreated" timestamp NOT NULL
);

-- Insert default values into admin_config
INSERT INTO admin_config (year, gameType, isVotingActive, votingEnd) VALUES
    (2024, 'semi1', false, NULL);


-- ------------------------------------------------------------
-- auth initial data for testing

CREATE DATABASE auth;

CREATE TABLE "auth" (
  "id" serial NOT NULL,
  PRIMARY KEY ("id"),
  "username" VARCHAR(50) UNIQUE NOT NULL,
  "password" TEXT NOT NULL
);

-- Insert default values into auth
INSERT INTO auth (username, password) VALUES
    ("test", "test");
