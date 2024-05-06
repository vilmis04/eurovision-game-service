
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
  "datecreated" timestamp NOT NULL
);

CREATE TABLE "country" (
  "id"          SERIAL PRIMARY KEY,
	"name"        VARCHAR(255) NOT NULL,
	"code"        VARCHAR(255) NOT NULL,
	"year"        INT NOT NULL,
	"gametype"    VARCHAR(50) NOT NULL,
	"score"       INT NOT NULL,
	"isinfinal"   BOOLEAN NOT NULL,
	"artist"      VARCHAR(255) NOT NULL,
	"song"        VARCHAR(255) NOT NULL,
	"ordersemi"   INT NOT NULL,
	"orderfinal"  INT NOT NULL
);

CREATE TABLE score (
  "id"       SERIAL PRIMARY KEY,
	"country"  VARCHAR(255) NOT NULL,
	"year"     INT NOT NULL,
	"gametype" VARCHAR(255) NOT NULL,
	"user"     VARCHAR(255) NOT NULL,         
	"infinal"  BOOLEAN NOT NULL,           
	"position"   INT NOT NULL         
);

-- Insert countries into table country
INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Cyprus', 'cy', 2024, 'semi1', 0, false, 'Silia Kapsis', 'Liar', 1, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Serbia', 'rs', 2024, 'semi1', 0, false, 'TEYA DORA', 'RAMONDA', 2, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Lithuania', 'lt', 2024, 'semi1', 0, false, 'Silvester Belt', 'Luktelk', 3, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Ireland', 'ie', 2024, 'semi1', 0, false, 'Bambie Thug', 'Doomsday Blue', 4, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Ukraine', 'ua', 2024, 'semi1', 0, false, 'alyona alyona & Jerry Heil', 'Teresa & Maria', 5, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Poland', 'pl', 2024, 'semi1', 0, false, 'LUNA', 'The Tower', 6, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Croatia', 'hr', 2024, 'semi1', 0, false, 'Baby Lasagna', 'Rim Tim Tagi Dim', 7, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Iceland', 'is', 2024, 'semi1', 0, false, 'Hera Björk', 'Scared of Heights', 8, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Slovenia', 'si', 2024, 'semi1', 0, false, 'Raiven', 'Veronika', 9, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Finland', 'fi', 2024, 'semi1', 0, false, 'Windows95man', 'No Rules!', 10, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Moldova', 'md', 2024, 'semi1', 0, false, 'Natalia Barbu', 'In The Middle', 11, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Azerbaijan', 'az', 2024, 'semi1', 0, false, 'FAHREE feat. Ilkin Dovlatov', 'Özünlə Apar', 12, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Australia', 'au', 2024, 'semi1', 0, false, 'Electric Fields', 'One Milkali (One Blood)', 13, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Portugal', 'pt', 2024, 'semi1', 0, false, 'iolanda', 'Grito', 14, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Luxembourg', 'lu', 2024, 'semi1', 0, false, 'TALI', 'Fighter', 15, 0);



INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Malta', 'mt', 2024, 'semi2', 0, false, 'Sarah Bonnici', 'Loop', 1, 0);
		
INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Albania', 'al', 2024, 'semi2', 0, false, 'Besa', 'Titan', 2, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Greece', 'gr', 2024, 'semi2', 0, false, 'Marina Satti', 'ZARI', 3, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Switzerland', 'ch', 2024, 'semi2', 0, false, 'Nemo', 'The Code', 4, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Czechia', 'cz', 2024, 'semi2', 0, false, 'Aiko', 'Pedestal', 5, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Austria', 'at', 2024, 'semi2', 0, false, 'Kaleen', 'We Will Rave', 6, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Denmark', 'dk', 2024, 'semi2', 0, false, 'SABA', 'SAND', 7, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Armenia', 'am', 2024, 'semi2', 0, false, 'LADANIVA', 'Jako', 8, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Latvia', 'lv', 2024, 'semi2', 0, false, 'Dons', 'Hollow', 9, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('San Marino', 'sm', 2024, 'semi2', 0, false, 'MEGARA', '11:11', 10, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Georgia', 'ge', 2024, 'semi2', 0, false, 'Nutsa Buzaladze', 'Firefighter', 11, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Belgium', 'be', 2024, 'semi2', 0, false, 'Mustii', 'Before the Party''s Over', 12, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Estonia', 'ee', 2024, 'semi2', 0, false, '5MIINUST x Puuluup', '(nendest) narkootikumidest ei tea me (küll) midagi', 13, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Israel', 'il', 2024, 'semi2', 0, false, 'Eden Golan', 'Hurricane', 14, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Norway', 'no', 2024, 'semi2', 0, false, 'Gåte', 'Ulveham', 15, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Netherlands', 'nl', 2024, 'semi2', 0, false, 'Joost Klein', 'Europapa', 16, 0);



INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Germany', 'de', 2024, 'final', 0, true, 'ISAAK', 'Always On The Run', 0, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('United Kingdom', 'gb', 2024, 'final', 0, true, 'Olly Alexander', 'Dizzy', 0, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('France', 'fr', 2024, 'final', 0, true, 'Slimane', 'Mon amour', 0, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Italy', 'it', 2024, 'final', 0, true, 'Angelina Mango', 'La noia', 0, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Spain', 'es', 2024, 'final', 0, true, 'Nebulossa', 'ZORRA', 0, 0);

INSERT INTO country (name, code, year, gameType, score, isInFinal, artist, song, orderSemi, orderFinal)
VALUES ('Sweden', 'se', 2024, 'final', 0, true, 'Marcus & Martinus', 'Unforgettable', 0, 1);


-- Insert default values into admin_config
INSERT INTO admin_config (year, gameType, isVotingActive, votingEnd) VALUES
    (2024, 'semi1', false, NULL);


-- ------------------------------------------------------------
-- auth initial data for testing

CREATE TABLE "auth" (
  "id" serial NOT NULL,
  PRIMARY KEY ("id"),
  "username" VARCHAR(50) UNIQUE NOT NULL,
  "password" TEXT NOT NULL
);
