import { Module } from "@nestjs/common";
import { CountryService } from "./country.service";
import { CountryController } from "./country.controller";
import { RepoClient } from "../../utils/RepoClient";
import { AdminRoleGuard } from "../admin/adminRole.guard";
import { JwtUtils } from "../../utils/JwtUtils";

@Module({
	controllers: [CountryController],
	providers: [CountryService, RepoClient, AdminRoleGuard, JwtUtils],
})
export class CountryModule {}
