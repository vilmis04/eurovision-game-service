import { Module } from "@nestjs/common";
import { JwtUtils } from "../../utils/JwtUtils";
import { RepoClient } from "../../utils/RepoClient";
import { UsersController } from "./users.controller";
import { UsersService } from "./users.service";
import { AuthModule } from "../auth/auth.module";

@Module({
	imports: [AuthModule],
	providers: [UsersService, RepoClient, JwtUtils],
	controllers: [UsersController],
})
export class UsersModule {}
