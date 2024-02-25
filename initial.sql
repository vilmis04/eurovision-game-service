
CREATE DATABASE IF NOT EXISTS ev_game;

CREATE TABLE IF NOT EXISTS admin_config (
    id SERIAL PRIMARY KEY,
    year INT NOT NULL,
    gameType VARCHAR(255) NOT NULL,
    isVotingActive BOOLEAN NOT NULL
);

-- Insert default values into admin_config
INSERT INTO admin_config (year, gameType, isVotingActive) VALUES
    (2024, 'semi1', false);


-- ------------------------------------------------------------
-- auth initial data for testing

CREATE DATABASE IF NOT EXISTS ev_game;

CREATE TABLE IF NOT EXISTS auth (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
);

-- Insert default values into auth
INSERT INTO auth (username, password) VALUES
    ("test", "test");
