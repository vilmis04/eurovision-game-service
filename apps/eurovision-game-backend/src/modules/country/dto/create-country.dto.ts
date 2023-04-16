import { ICountry } from "@eurovision-game-monorepo/core";

export class CreateCountryDto {
	year: ICountry["year"];
	type: ICountry["type"];
	name: ICountry["name"];
	artist: ICountry["artist"];
	song: ICountry["song"];
}
