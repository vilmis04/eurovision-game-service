import { ICountry } from "@eurovision-game-monorepo/core";

export class Country {
	year: ICountry["year"];
	type: ICountry["type"];
	name: ICountry["name"];
	artist: ICountry["artist"];
	song: ICountry["song"];
	semiFinalScore: ICountry["semiFinalScore"];
	finalScore: ICountry["finalScore"];
}
