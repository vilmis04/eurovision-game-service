import { IGetVotesResponse } from "@eurovision-game-monorepo/core";

// TODO: write tests
export const filterCountries = (votes: IGetVotesResponse) => {
	const values = Object.entries(votes);
	const filtered = values.filter(
		([key]) => key !== "_id" && key !== "username"
	);
	console.log(Object.fromEntries(filtered));
	return Object.fromEntries(filtered);
};
