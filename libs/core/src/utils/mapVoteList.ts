import { ICountry, IFinalsList, TCountries, VoteTypes } from "../lib";

export const mapSemiCountryList = (countryList: ICountry[] | undefined) =>
	(countryList || []).reduce(
		(result, { name, semiFinalScore }) =>
			semiFinalScore ? [...result, name] : result,
		[] as string[]
	);

export const mapSemiVotesList = (countryList: TCountries | undefined) =>
	Object.entries(countryList || {}).reduce(
		(arr, [country, score]) => (Boolean(score) ? [...arr, country] : arr),
		[] as string[]
	);

export const mapFinalVotesList = (
	countryList: TCountries | undefined
): IFinalsList[] =>
	Object.entries(countryList || {}).map(([country, score]) => ({
		name: country,
		position: score as VoteTypes,
	}));
