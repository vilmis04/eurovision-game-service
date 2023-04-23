import { Injectable, Req } from "@nestjs/common";
import { Request } from "express";

import {
	GameTypes,
	IGetVotesRequest,
	IUpdateVotesRequest,
	TCountries,
} from "@eurovision-game-monorepo/core";
import { RepoClient } from "../../utils/RepoClient";
import { UpdateResult } from "mongodb";
import { JwtUtils } from "../../utils/JwtUtils";
import { Votes } from "./entities/Votes";

@Injectable()
export class VotesService {
	constructor(private repoClient: RepoClient, private jwtUtils: JwtUtils) {}

	async getVotes(
		@Req() req: Request,
		year: string,
		type: GameTypes
	): Promise<TCountries> {
		const { username } = this.jwtUtils.getUser(req);
		const votesInstance = await this.repoClient.getVotes({
			username,
			year,
			type,
		});

		const voteData =
			votesInstance || (await this.createVotes({ username, year, type }));

		return voteData.votes;
	}

	async updateVotes(
		@Req() req: Request,
		requestBody: IUpdateVotesRequest
	): Promise<UpdateResult> {
		const { username } = this.jwtUtils.getUser(req);
		const response = await this.repoClient.updateVotes({
			username,
			...requestBody,
		});

		return response;
	}

	async createVotes({
		username,
		type,
		year,
	}: IGetVotesRequest): Promise<Votes> {
		const userVotes = await this.repoClient.createVotes({
			username,
			type,
			year,
		});

		return userVotes;
	}
}
