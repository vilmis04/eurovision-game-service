import { Module } from "@nestjs/common";
import { RepoClient } from "../../utils/RepoClient";
import { VotesController } from "./votes.controller";
import { VotesService } from "./votes.service";

@Module({
	imports: [],
	controllers: [VotesController],
	providers: [VotesService, RepoClient],
})
export class VotesModule {}
