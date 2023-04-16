import { IAdminFormData, TCountries } from "@eurovision-game-monorepo/core";

export class Score {
	year: IAdminFormData["year"];
	type: IAdminFormData["type"];
	countries: TCountries;
}
