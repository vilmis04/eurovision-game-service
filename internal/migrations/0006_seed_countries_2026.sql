-- +goose Up
-- Countries of the 2026 contest. Safe to re-run: existing rows (and the scores and
-- finalist flags already stored on them) are left untouched.
INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES
    ('Moldova', 'md', 2026, 'semi1', 0, false, 'Satoshi', 'Viva, Moldova!', 1, 0),
    ('Sweden', 'se', 2026, 'semi1', 0, false, 'FELICIA', 'My System', 2, 0),
    ('Croatia', 'hr', 2026, 'semi1', 0, false, 'LELEK', 'Andromeda', 3, 0),
    ('Greece', 'gr', 2026, 'semi1', 0, false, 'Akylas', 'Ferto', 4, 0),
    ('Portugal', 'pt', 2026, 'semi1', 0, false, 'Bandidos do Cante', 'Rosa', 5, 0),
    ('Georgia', 'ge', 2026, 'semi1', 0, false, 'Bzikebi', 'On Replay', 6, 0),
    ('Finland', 'fi', 2026, 'semi1', 0, false, 'Linda Lampenius x Pete Parkkonen', 'Liekinheitin', 7, 0),
    ('Montenegro', 'me', 2026, 'semi1', 0, false, 'Tamara Živković', 'Nova Zora', 8, 0),
    ('Estonia', 'ee', 2026, 'semi1', 0, false, 'Vanilla Ninja', 'Too Epic To Be True', 9, 0),
    ('Israel', 'il', 2026, 'semi1', 0, false, 'Noam Bettan', 'Michelle', 10, 0),
    ('Belgium', 'be', 2026, 'semi1', 0, false, 'ESSYLA', 'Dancing on the Ice', 11, 0),
    ('Lithuania', 'lt', 2026, 'semi1', 0, false, 'Lion Ceccah', 'Sólo Quiero Más', 12, 0),
    ('San Marino', 'sm', 2026, 'semi1', 0, false, 'SENHIT', 'Superstar', 13, 0),
    ('Poland', 'pl', 2026, 'semi1', 0, false, 'ALICJA', 'Pray', 14, 0),
    ('Serbia', 'rs', 2026, 'semi1', 0, false, 'LAVINA', 'Krah Mene', 15, 0),
    ('Bulgaria', 'bg', 2026, 'semi2', 0, false, 'DARA', 'Bangaranga', 1, 0),
    ('Azerbaijan', 'az', 2026, 'semi2', 0, false, 'JIVA', 'Just Go', 2, 0),
    ('Romania', 'ro', 2026, 'semi2', 0, false, 'Alexandra Căpitănescu', 'Choke Me', 3, 0),
    ('Luxembourg', 'lu', 2026, 'semi2', 0, false, 'Eva Marija', 'Mother Nature', 4, 0),
    ('Czechia', 'cz', 2026, 'semi2', 0, false, 'Daniel Zizka', 'CROSSROADS', 5, 0),
    ('Armenia', 'am', 2026, 'semi2', 0, false, 'SIMÓN', 'Paloma Rumba', 6, 0),
    ('Switzerland', 'ch', 2026, 'semi2', 0, false, 'Veronica Fusaro', 'Alice', 7, 0),
    ('Cyprus', 'cy', 2026, 'semi2', 0, false, 'Antigoni', 'JALLA', 8, 0),
    ('Latvia', 'lv', 2026, 'semi2', 0, false, 'Atvara', 'Ēnā', 9, 0),
    ('Australia', 'au', 2026, 'semi2', 0, false, 'Delta Goodrem', 'Exclipse', 10, 0),
    ('Ukraine', 'ua', 2026, 'semi2', 0, false, 'LELÉKA', 'Ridnym', 11, 0),
    ('Albania', 'al', 2026, 'semi2', 0, false, 'Alis', 'Nân', 12, 0),
    ('Malta', 'mt', 2026, 'semi2', 0, false, 'AIDAN', 'Bella', 13, 0),
    ('Norway', 'no', 2026, 'semi2', 0, false, 'JONAS LOVV', 'YA YA YA', 14, 0),
    ('Austria', 'at', 2026, 'final', 0, true, 'COSMÓ', 'Tanzschein', 0, 0),
    ('France', 'fr', 2026, 'final', 0, true, 'Monroe', 'Regarde !', 0, 0),
    ('Germany', 'de', 2026, 'final', 0, true, 'Sarah Engels', 'Fire', 0, 0),
    ('Italy', 'it', 2026, 'final', 0, true, 'Sal Da Vinci', 'Per Sempre Sì', 0, 0),
    ('United Kingdom', 'gb', 2026, 'final', 0, true, 'LOOK MUM NO COMPUTER', 'Eins, Zwei, Drei', 0, 0)
ON CONFLICT (year, name) DO NOTHING;

-- +goose Down
DELETE FROM country WHERE year = 2026;
