-- +goose Up
-- Countries of the 2024 contest. Safe to re-run: existing rows (and the scores and
-- finalist flags already stored on them) are left untouched.
INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES
    ('Cyprus', 'cy', 2024, 'semi1', 0, false, 'Silia Kapsis', 'Liar', 1, 0),
    ('Serbia', 'rs', 2024, 'semi1', 0, false, 'TEYA DORA', 'RAMONDA', 2, 0),
    ('Lithuania', 'lt', 2024, 'semi1', 0, false, 'Silvester Belt', 'Luktelk', 3, 0),
    ('Ireland', 'ie', 2024, 'semi1', 0, false, 'Bambie Thug', 'Doomsday Blue', 4, 0),
    ('Ukraine', 'ua', 2024, 'semi1', 0, false, 'alyona alyona & Jerry Heil', 'Teresa & Maria', 5, 0),
    ('Poland', 'pl', 2024, 'semi1', 0, false, 'LUNA', 'The Tower', 6, 0),
    ('Croatia', 'hr', 2024, 'semi1', 0, false, 'Baby Lasagna', 'Rim Tim Tagi Dim', 7, 0),
    ('Iceland', 'is', 2024, 'semi1', 0, false, 'Hera Björk', 'Scared of Heights', 8, 0),
    ('Slovenia', 'si', 2024, 'semi1', 0, false, 'Raiven', 'Veronika', 9, 0),
    ('Finland', 'fi', 2024, 'semi1', 0, false, 'Windows95man', 'No Rules!', 10, 0),
    ('Moldova', 'md', 2024, 'semi1', 0, false, 'Natalia Barbu', 'In The Middle', 11, 0),
    ('Azerbaijan', 'az', 2024, 'semi1', 0, false, 'FAHREE feat. Ilkin Dovlatov', 'Özünlə Apar', 12, 0),
    ('Australia', 'au', 2024, 'semi1', 0, false, 'Electric Fields', 'One Milkali (One Blood)', 13, 0),
    ('Portugal', 'pt', 2024, 'semi1', 0, false, 'iolanda', 'Grito', 14, 0),
    ('Luxembourg', 'lu', 2024, 'semi1', 0, false, 'TALI', 'Fighter', 15, 0),
    ('Malta', 'mt', 2024, 'semi2', 0, false, 'Sarah Bonnici', 'Loop', 1, 0),
    ('Albania', 'al', 2024, 'semi2', 0, false, 'Besa', 'Titan', 2, 0),
    ('Greece', 'gr', 2024, 'semi2', 0, false, 'Marina Satti', 'ZARI', 3, 0),
    ('Switzerland', 'ch', 2024, 'semi2', 0, false, 'Nemo', 'The Code', 4, 0),
    ('Czechia', 'cz', 2024, 'semi2', 0, false, 'Aiko', 'Pedestal', 5, 0),
    ('Austria', 'at', 2024, 'semi2', 0, false, 'Kaleen', 'We Will Rave', 6, 0),
    ('Denmark', 'dk', 2024, 'semi2', 0, false, 'SABA', 'SAND', 7, 0),
    ('Armenia', 'am', 2024, 'semi2', 0, false, 'LADANIVA', 'Jako', 8, 0),
    ('Latvia', 'lv', 2024, 'semi2', 0, false, 'Dons', 'Hollow', 9, 0),
    ('San Marino', 'sm', 2024, 'semi2', 0, false, 'MEGARA', '11:11', 10, 0),
    ('Georgia', 'ge', 2024, 'semi2', 0, false, 'Nutsa Buzaladze', 'Firefighter', 11, 0),
    ('Belgium', 'be', 2024, 'semi2', 0, false, 'Mustii', 'Before the Party''s Over', 12, 0),
    ('Estonia', 'ee', 2024, 'semi2', 0, false, '5MIINUST x Puuluup', '(nendest) narkootikumidest ei tea me (küll) midagi', 13, 0),
    ('Israel', 'il', 2024, 'semi2', 0, false, 'Eden Golan', 'Hurricane', 14, 0),
    ('Norway', 'no', 2024, 'semi2', 0, false, 'Gåte', 'Ulveham', 15, 0),
    ('Netherlands', 'nl', 2024, 'semi2', 0, false, 'Joost Klein', 'Europapa', 16, 0),
    ('Germany', 'de', 2024, 'final', 0, true, 'ISAAK', 'Always On The Run', 0, 0),
    ('United Kingdom', 'gb', 2024, 'final', 0, true, 'Olly Alexander', 'Dizzy', 0, 0),
    ('France', 'fr', 2024, 'final', 0, true, 'Slimane', 'Mon amour', 0, 0),
    ('Italy', 'it', 2024, 'final', 0, true, 'Angelina Mango', 'La noia', 0, 0),
    ('Spain', 'es', 2024, 'final', 0, true, 'Nebulossa', 'ZORRA', 0, 0),
    ('Sweden', 'se', 2024, 'final', 0, true, 'Marcus & Martinus', 'Unforgettable', 0, 1)
ON CONFLICT (year, name) DO NOTHING;

-- +goose Down
DELETE FROM country WHERE year = 2024;
