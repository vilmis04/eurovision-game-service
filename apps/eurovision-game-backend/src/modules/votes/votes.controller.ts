import { IGetVotesResponse } from "@eurovision-game-monorepo/core";
import { Body, Controller, Get, Patch, Req } from "@nestjs/common";
import { RootPaths } from "../../types/paths";
import { VotesService } from "./votes.service";
import { Request } from "express";
@Controller(RootPaths.VOTES)
export class VotesController {
	constructor(private readonly votesService: VotesService) {}

	@Get()
	getVotesByUsername(@Req() req: Request) {
		return this.votesService.getVotesByUsername(req);
	}
	@Patch()
	editVotesByUsername(@Req() req: Request, @Body() votes: IGetVotesResponse) {
		return this.votesService.editVotesByUsername(req, votes);
	}
}
