import { Module } from "@nestjs/common";
import { RepoClient } from "../../utils/RepoClient";
import { AuthController } from "./auth.controller";
import { AuthService } from "./auth.service";
import { JwtUtils } from "../../utils/JwtUtils";
import { GroupService } from "../group/group.service";
import { UsersService } from "../users/users.service";
import { VotesService } from "../votes/votes.service";
import { AdminService } from "../admin/admin.service";
import { CountryService } from "../country/country.service";

@Module({
	imports: [],
	controllers: [AuthController],
	providers: [
		AuthService,
		RepoClient,
		JwtUtils,
		GroupService,
		UsersService,
		VotesService,
		AdminService,
		CountryService,
	],
})
export class AuthModule {}
