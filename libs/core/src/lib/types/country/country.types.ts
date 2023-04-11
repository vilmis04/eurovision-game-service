import { GameTypes } from "../admin";

export interface ICountry {
	year: string;
	type: GameTypes;
	name: string;
	artist: string;
	song: string;
	semiFinalScore: boolean;
	finalScore: number;
}

export interface ICreateCountryFormData
	extends Pick<ICountry, "name" | "artist" | "song"> {}
