import { Controller, Get } from "@nestjs/common";
import { VotesService } from "./votes.service";

@Controller("votes")
export class VotesController {
	constructor(private readonly votesService: VotesService) {}

	@Get()
	getVotesByUsername() {
		return this.votesService.getVotesByUsername();
	}
}
