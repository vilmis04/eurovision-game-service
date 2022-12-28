import { Injectable } from "@nestjs/common";

import { IGetVotesResponse } from "@eurovision-game-monorepo/core";
import { RepoClient } from "../../utils/RepoClient";
import { UpdateResult } from "mongodb";

@Injectable()
export class VotesService {
	constructor(private repoClient: RepoClient) {}

	async getVotesByUsername(): Promise<IGetVotesResponse> {
		// TODO: get username from auth session
		const username = "test_user1";
		const values = await this.repoClient.getVotesByUserName({
			username,
		});

		if (values) {
			const { _id, username, ...countries } = values;
			return countries;
		}

		throw new Error("No values found");
	}
	async editVotesByUsername(votes: IGetVotesResponse): Promise<UpdateResult> {
		// TODO: get username from auth session
		const username = "test_user1";
		const response = await this.repoClient.editVotesByUserName({
			username,
			votes,
		});

		return response;
	}
}
