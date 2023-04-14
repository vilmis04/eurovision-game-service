export type TCountries = {
	[k: string]: string | number;
};

export interface IAdminFormData {
	isRegistrationEnabled: boolean;
	year: string;
	type: GameTypes;
}

export enum GameTypes {
	SEMI_1 = "SEMI_1",
	SEMI_2 = "SEMI_2",
	FINAL = "FINAL",
}

export const deprecated_initialValues = {
	ukraine: 0,
	united_kingdom: 0,
	spain: 0,
	sweden: 0,
	serbia: 0,
	norway: 0,
	belgium: 0,
	czech_republic: 0,
	romania: 0,
	portugal: 0,
	finland: 0,
	switzerland: 0,
	france: 0,
	armenia: 0,
	italy: 0,
	netherlands: 0,
	germany: 0,
	lithuania: 0,
	azerbaijan: 0,
	greece: 0,
	iceland: 0,
	moldova: 0,
	australia: 0,
	poland: 0,
	estonia: 0,
};
