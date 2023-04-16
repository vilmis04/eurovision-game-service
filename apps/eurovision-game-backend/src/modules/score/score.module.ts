import { Module } from "@nestjs/common";
import { ScoreService } from "./score.service";
import { ScoreController } from "./score.controller";
import { RepoClient } from "../../utils/RepoClient";
import { AdminRoleGuard } from "../admin/adminRole.guard";
import { JwtUtils } from "../../utils/JwtUtils";

@Module({
	controllers: [ScoreController],
	providers: [ScoreService, RepoClient, AdminRoleGuard, JwtUtils],
})
export class ScoreModule {}
