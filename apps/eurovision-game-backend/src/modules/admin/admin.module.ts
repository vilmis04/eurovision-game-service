import { Module } from "@nestjs/common";
import { AdminService } from "./admin.service";
import { AdminController } from "./admin.controller";
import { RepoClient } from "../../utils/RepoClient";

@Module({
	controllers: [AdminController],
	providers: [AdminService, RepoClient],
})
export class AdminModule {}
