import { IGetVotesResponse } from "@eurovision-game-monorepo/core";
import { Body, Controller, Get, Patch } from "@nestjs/common";
import { VotesService } from "./votes.service";

@Controller("votes")
export class VotesController {
	constructor(private readonly votesService: VotesService) {}

	@Get()
	getVotesByUsername() {
		return this.votesService.getVotesByUsername();
	}
	@Patch()
	editVotesByUsername(@Body() votes: IGetVotesResponse) {
		return this.votesService.editVotesByUsername(votes);
	}
}
