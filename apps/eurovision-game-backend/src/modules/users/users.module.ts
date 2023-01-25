import { Module } from "@nestjs/common";
import { JwtUtils } from "../../utils/JwtUtils";
import { RepoClient } from "../../utils/RepoClient";
import { AuthService } from "../auth/auth.service";
import { VotesService } from "../votes/votes.service";
import { UsersController } from "./users.controller";
import { UsersService } from "./users.service";

@Module({
	providers: [UsersService, RepoClient, VotesService, AuthService, JwtUtils],
	controllers: [UsersController],
})
export class UsersModule {}
