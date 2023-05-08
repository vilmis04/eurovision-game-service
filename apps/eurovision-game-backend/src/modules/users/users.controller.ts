import { Body, Controller, Patch, Req } from "@nestjs/common";
import { Request } from "express";
import { UsersService } from "./users.service";
import { RootPaths } from "../../types/paths";
import { User } from "./entities/user.entity";
import { JwtUtils } from "../../utils/JwtUtils";

@Controller(RootPaths.USERS)
export class UsersController {
	constructor(
		private readonly userService: UsersService,
		private readonly jwtUtils: JwtUtils
	) {}

	@Patch()
	async updateUser(
		@Req() request: Request,
		@Body() requestBody: Partial<User>
	) {
		const { username } = this.jwtUtils.getUser(request);
		return await this.userService.updateUser(username, requestBody);
	}
}
