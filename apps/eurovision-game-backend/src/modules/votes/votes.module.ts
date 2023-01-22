import { Module } from "@nestjs/common";
import { JwtUtils } from "../../utils/JwtUtils";
import { RepoClient } from "../../utils/RepoClient";
import { VotesController } from "./votes.controller";
import { VotesService } from "./votes.service";

@Module({
	imports: [],
	controllers: [VotesController],
	providers: [VotesService, RepoClient, JwtUtils],
})
export class VotesModule {}
