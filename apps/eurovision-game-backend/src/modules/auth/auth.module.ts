import { Module } from "@nestjs/common";
import { RepoClient } from "../../utils/RepoClient";
import { AuthController } from "./auth.controller";
import { AuthService } from "./auth.service";
import { JwtUtils } from "../../utils/JwtUtils";

@Module({
	imports: [],
	controllers: [AuthController],
	providers: [AuthService, RepoClient, JwtUtils],
})
export class AuthModule {}
