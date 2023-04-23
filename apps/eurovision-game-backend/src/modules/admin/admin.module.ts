import { Module } from "@nestjs/common";
import { AdminService } from "./admin.service";
import { AdminController } from "./admin.controller";
import { RepoClient } from "../../utils/RepoClient";
import { AdminRoleGuard } from "./adminRole.guard";
import { JwtUtils } from "../../utils/JwtUtils";
import { CountryService } from "../country/country.service";

@Module({
	controllers: [AdminController],
	providers: [
		AdminService,
		RepoClient,
		AdminRoleGuard,
		JwtUtils,
		CountryService,
	],
})
export class AdminModule {}
