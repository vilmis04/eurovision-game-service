-- +goose Up
-- Baseline schema. Every statement is IF NOT EXISTS so that a database that was
-- created earlier from initial.sql adopts this migration without changes.
CREATE TABLE IF NOT EXISTS admin_config (
    id             SERIAL PRIMARY KEY,
    year           INT NOT NULL,
    gameType       VARCHAR(255) NOT NULL,
    isVotingActive BOOLEAN NOT NULL,
    votingEnd      TIMESTAMPTZ DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS "group" (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(20) NOT NULL,
    owner       VARCHAR(50) NOT NULL,
    members     TEXT[] NOT NULL,
    datecreated TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS country (
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    code       VARCHAR(255) NOT NULL,
    year       INT NOT NULL,
    gametype   VARCHAR(50) NOT NULL,
    score      INT NOT NULL,
    isinfinal  BOOLEAN NOT NULL,
    artist     VARCHAR(255) NOT NULL,
    song       VARCHAR(255) NOT NULL,
    ordersemi  INT NOT NULL,
    orderfinal INT NOT NULL
);

CREATE TABLE IF NOT EXISTS score (
    id       SERIAL PRIMARY KEY,
    country  VARCHAR(255) NOT NULL,
    year     INT NOT NULL,
    gametype VARCHAR(255) NOT NULL,
    "user"   VARCHAR(255) NOT NULL,
    infinal  BOOLEAN NOT NULL,
    position INT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS score;
DROP TABLE IF EXISTS country;
DROP TABLE IF EXISTS "group";
DROP TABLE IF EXISTS admin_config;
