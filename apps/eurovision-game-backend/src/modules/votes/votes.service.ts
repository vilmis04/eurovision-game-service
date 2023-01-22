import { Injectable, Req } from "@nestjs/common";
import { Request } from "express";

import { IGetVotesResponse } from "@eurovision-game-monorepo/core";
import { RepoClient } from "../../utils/RepoClient";
import { InsertOneResult, UpdateResult } from "mongodb";
import { JwtUtils } from "../../utils/JwtUtils";

@Injectable()
export class VotesService {
	constructor(private repoClient: RepoClient, private jwtUtils: JwtUtils) {}

	async getVotesByUsername(@Req() req: Request): Promise<IGetVotesResponse> {
		const username = this.jwtUtils.getUsername(req);
		const values = await this.repoClient.getVotesByUserName({
			username,
		});

		if (values) {
			const { _id, username, ...countries } = values;
			return countries;
		}

		throw new Error("No values found");
	}
	async editVotesByUsername(
		@Req() req: Request,
		votes: IGetVotesResponse
	): Promise<UpdateResult> {
		const username = this.jwtUtils.getUsername(req);
		const response = await this.repoClient.editVotesByUserName({
			username,
			votes,
		});

		return response;
	}
	async createNewUserVotes(username: string): Promise<InsertOneResult> {
		const userVotes = await this.repoClient.createNewUserVotes(username);

		return userVotes;
	}
}
