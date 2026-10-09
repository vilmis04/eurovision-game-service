-- +goose Up
-- Countries of the 2025 contest. Safe to re-run: existing rows (and the scores and
-- finalist flags already stored on them) are left untouched.
INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES
    ('Iceland', 'is', 2025, 'semi1', 0, false, 'VÆB', 'RÓA', 1, 0),
    ('Poland', 'pl', 2025, 'semi1', 0, false, 'Justyna Steczkowska', 'GAJA', 2, 0),
    ('Slovenia', 'si', 2025, 'semi1', 0, false, 'Klemen', 'How Much Time Do We Have Left', 3, 0),
    ('Estonia', 'ee', 2025, 'semi1', 0, false, 'Tommy Cash', 'Espresso Macchiato', 4, 0),
    ('Ukraine', 'ua', 2025, 'semi1', 0, false, 'Ziferblat', 'Bird of Pray', 5, 0),
    ('Sweden', 'se', 2025, 'semi1', 0, false, 'KAJ', 'Bara Bada Bastu', 6, 0),
    ('Portugal', 'pt', 2025, 'semi1', 0, false, 'NAPA', 'Deslocado', 7, 0),
    ('Norway', 'no', 2025, 'semi1', 0, false, 'Kyle Alessandro', 'Lighter', 8, 0),
    ('Belgium', 'be', 2025, 'semi1', 0, false, 'Red Sebastian', 'Strobe Lights', 9, 0),
    ('Azerbaijan', 'az', 2025, 'semi1', 0, false, 'Mamagama', 'Run With U', 10, 0),
    ('San Marino', 'sm', 2025, 'semi1', 0, false, 'Gabry Ponte', 'Tutta L''Italia', 11, 0),
    ('Albania', 'al', 2025, 'semi1', 0, false, 'Shkodra Elektronike', 'Zjerm', 12, 0),
    ('Netherlands', 'nl', 2025, 'semi1', 0, false, 'Claude', 'C''est La Vie', 13, 0),
    ('Croatia', 'hr', 2025, 'semi1', 0, false, 'Marko Bošnjak', 'Poison Cake', 14, 0),
    ('Cyprus', 'cy', 2025, 'semi1', 0, false, 'Theo Evan', 'Shh', 15, 0),
    ('Australia', 'au', 2025, 'semi2', 0, false, 'Go-Jo', 'Milkshake Man', 1, 0),
    ('Montenegro', 'me', 2025, 'semi2', 0, false, 'Nina Žižić', 'Dobrodošli', 2, 0),
    ('Ireland', 'ie', 2025, 'semi2', 0, false, 'EMMY', 'Laika Party', 3, 0),
    ('Latvia', 'lv', 2025, 'semi2', 0, false, 'Tautumeitas', 'Bur Man Laimi', 4, 0),
    ('Armenia', 'am', 2025, 'semi2', 0, false, 'PARG', 'SURVIVOR', 5, 0),
    ('Austria', 'at', 2025, 'semi2', 0, false, 'JJ', 'Wasted Love', 6, 0),
    ('Greece', 'gr', 2025, 'semi2', 0, false, 'Klavdia', 'Asteromáta', 7, 0),
    ('Lithuania', 'lt', 2025, 'semi2', 0, false, 'Katarsis', 'Tavo Akys', 8, 0),
    ('Malta', 'mt', 2025, 'semi2', 0, false, 'Miriana Conte', 'SERVING', 9, 0),
    ('Georgia', 'ge', 2025, 'semi2', 0, false, 'Mariam Shengelia', 'Freedom', 10, 0),
    ('Denmark', 'dk', 2025, 'semi2', 0, false, 'Sissal', 'Hallucination', 11, 0),
    ('Czechia', 'cz', 2025, 'semi2', 0, false, 'ADONXS', 'Kiss Kiss Goodbye', 12, 0),
    ('Luxembourg', 'lu', 2025, 'semi2', 0, false, 'Laura Thorn', 'La Poupée Monte Le Son', 13, 0),
    ('Israel', 'il', 2025, 'semi2', 0, false, 'Yuval Raphael', 'New Day Will Rise', 14, 0),
    ('Serbia', 'rs', 2025, 'semi2', 0, false, 'Princ', 'Mila', 15, 0),
    ('Finland', 'fi', 2025, 'semi2', 0, false, 'Erika Vikman', 'ICH KOMME', 16, 0),
    ('Italy', 'it', 2025, 'final', 0, true, 'Lucio Corsi', 'Volevo Essere Un Duro', 0, 0),
    ('France', 'fr', 2025, 'final', 0, true, 'Louane', 'maman', 0, 0),
    ('Germany', 'de', 2025, 'final', 0, true, 'Abor & Tynna', 'Baller', 0, 0),
    ('Spain', 'es', 2025, 'final', 0, true, 'Melody', 'ESA DIVA', 0, 0),
    ('United Kingdom', 'gb', 2025, 'final', 0, true, 'Remember Monday', 'What The Hell Just Happened?', 0, 0),
    ('Switzerland', 'ch', 2025, 'final', 0, true, 'Zoë Më', 'Voyage', 0, 0)
ON CONFLICT (year, name) DO NOTHING;

-- +goose Down
DELETE FROM country WHERE year = 2025;
