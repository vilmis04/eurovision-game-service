
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
