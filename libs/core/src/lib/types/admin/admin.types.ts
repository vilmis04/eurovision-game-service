import { VoteTypes } from "../votes";

export type TCountries = {
	[k: string]: VoteTypes | string | number | boolean;
};

export interface IAdminFormData {
	isRegistrationEnabled: boolean;
	year: string;
	type: GameTypes;
	isVotingDisabled: boolean;
	calculateFinalsScore: boolean;
}

export enum GameTypes {
	SEMI_1 = "SEMI_1",
	SEMI_2 = "SEMI_2",
	FINAL = "FINAL",
}
