import { IGetUserResponse } from "@eurovision-game-monorepo/core";
import { Body, Controller, Post, Res } from "@nestjs/common";
import { Response } from "express";
import { UsersService } from "./users.service";

enum UserPaths {
	SIGN_UP = "signUp",
}

const MAX_AGE = 1000 * 24 * 3600;

@Controller("users")
export class UsersController {
	constructor(private readonly userService: UsersService) {}

	@Post(UserPaths.SIGN_UP)
	async signUp(
		@Res({ passthrough: true }) response: Response,
		@Body() { username, password }: IGetUserResponse
	) {
		console.log("reached users controller");
		const result = await this.userService.signUp(
			response,
			username,
			password
		);
		if (!result.success) return result;

		response.cookie("jwt", result.access_token, {
			maxAge: MAX_AGE,
			httpOnly: true,
		});

		return { success: true };
	}
}
