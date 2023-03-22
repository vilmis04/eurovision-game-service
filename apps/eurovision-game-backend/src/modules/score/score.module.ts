import { Module } from "@nestjs/common";
import { ScoreService } from "./score.service";
import { ScoreController } from "./score.controller";
import { RepoClient } from "../../utils/RepoClient";

@Module({
	controllers: [ScoreController],
	providers: [ScoreService, RepoClient],
})
export class ScoreModule {}
