import { Module } from "@nestjs/common";
import { RepoClient } from "../../utils/RepoClient";
import { AuthController } from "./auth.controller";
import { AuthService } from "./auth.service";

@Module({
	imports: [],
	controllers: [AuthController],
	providers: [AuthService, RepoClient],
})
export class AuthModule {}
