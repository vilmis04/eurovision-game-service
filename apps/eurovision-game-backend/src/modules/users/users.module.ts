import { Module } from "@nestjs/common";
import { RepoClient } from "../../utils/RepoClient";
import { AuthService } from "../auth/auth.service";
import { VotesService } from "../votes/votes.service";
import { UsersController } from "./users.controller";
import { UsersService } from "./users.service";

@Module({
	providers: [UsersService, RepoClient, VotesService, AuthService],
	controllers: [UsersController],
})
export class UsersModule {}
