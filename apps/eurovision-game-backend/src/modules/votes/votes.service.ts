import { Injectable } from "@nestjs/common";

import { IGetVotesResponse } from "@eurovision-game-monorepo/core";
import { RepoClient } from "../../utils/RepoClient";
import { filterCountries } from "../../utils/filterCountries";

@Injectable()
export class VotesService {
	constructor(private repoClient: RepoClient) {}

	async getVotesByUsername(): Promise<IGetVotesResponse> {
		const username = "test_user1";
		// TODO: fix typing coming from mongoDB
		const values = (await this.repoClient.getVotesByUserName({
			username,
		})) as unknown as IGetVotesResponse;

		return filterCountries(values) as unknown as IGetVotesResponse;
	}
}
