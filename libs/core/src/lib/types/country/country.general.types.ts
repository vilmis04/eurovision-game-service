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
