import { Body, Controller, Patch, Req } from "@nestjs/common";
import { Request } from "express";
import { UsersService } from "./users.service";
import { RootPaths } from "../../types/paths";
import { User } from "./entities/user.entity";

@Controller(RootPaths.USERS)
export class UsersController {
	constructor(private readonly userService: UsersService) {}

	@Patch()
	async updateUser(
		@Req() request: Request,
		@Body() requestBody: Partial<User>
	) {
		return await this.userService.updateUser(request, requestBody);
	}
}
