import { Module } from "@nestjs/common";
import { GroupService } from "./group.service";
import { GroupController } from "./group.controller";
import { RepoClient } from "../../utils/RepoClient";
import { JwtUtils } from "../../utils/JwtUtils";
import { UsersService } from "../users/users.service";
import { UsersModule } from "../users/users.module";

@Module({
	imports: [UsersModule],
	controllers: [GroupController],
	providers: [GroupService, RepoClient, JwtUtils, UsersService],
})
export class GroupModule {}
