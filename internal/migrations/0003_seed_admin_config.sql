-- +goose Up
-- The single configuration row the API reads. Existing configuration is kept.
INSERT INTO admin_config (id, year, gameType, isVotingActive, votingEnd)
VALUES (1, 2026, 'semi1', false, NULL)
ON CONFLICT (id) DO NOTHING;

-- +goose Down
DELETE FROM admin_config WHERE id = 1;
