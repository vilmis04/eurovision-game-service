import { GameTypes, TCountries } from "../admin";

export interface IScoreFormData {
	year: string;
	type: GameTypes;
	countries: TCountries;
}

export interface IGetScoreRequest {
	year: IScoreFormData["year"];
	type: IScoreFormData["type"];
}
