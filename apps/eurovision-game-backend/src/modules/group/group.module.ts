import { Module } from "@nestjs/common";
import { GroupService } from "./group.service";
import { GroupController } from "./group.controller";
import { RepoClient } from "../../utils/RepoClient";
import { JwtUtils } from "../../utils/JwtUtils";
import { UsersService } from "../users/users.service";
import { UsersModule } from "../users/users.module";
import { AdminService } from "../admin/admin.service";
import { VotesService } from "../votes/votes.service";
import { CountryService } from "../country/country.service";

@Module({
	imports: [UsersModule],
	controllers: [GroupController],
	providers: [
		GroupService,
		RepoClient,
		JwtUtils,
		UsersService,
		AdminService,
		VotesService,
		CountryService,
	],
})
export class GroupModule {}
