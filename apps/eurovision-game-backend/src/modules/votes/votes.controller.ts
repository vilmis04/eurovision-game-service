import { GameTypes, IUpdateVotesRequest } from "@eurovision-game-monorepo/core";
import { Body, Controller, Get, Param, Patch, Req } from "@nestjs/common";
import { RootPaths } from "../../types/paths";
import { VotesService } from "./votes.service";
import { Request } from "express";

@Controller(RootPaths.VOTES)
export class VotesController {
	constructor(private readonly votesService: VotesService) {}

	@Get(":year/:type")
	getVotes(
		@Req() req: Request,
		@Param("year") year: string,
		@Param("type") type: GameTypes
	) {
		return this.votesService.getVotes(req, year, type);
	}

	@Patch()
	updateVotes(@Req() req: Request, @Body() requestBody: IUpdateVotesRequest) {
		return this.votesService.updateVotes(req, requestBody);
	}
}
