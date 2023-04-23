import { GameTypes, TCountries } from "../admin";

export interface IVotes {
	username: string;
	type: GameTypes;
	year: string;
	votes: TCountries;
}

export enum VoteTypes {
	P1 = "1",
	P2_5 = "2-5",
	P6_10 = "6-10",
	P11_15 = "11-15",
	P16_20 = "16-20",
	P21_25 = "21-25",
}
