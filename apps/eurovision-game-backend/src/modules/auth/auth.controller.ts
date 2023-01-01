import { IGetUserResponse } from "@eurovision-game-monorepo/core";
import { Body, Controller, Post, Res } from "@nestjs/common";
import { Response } from "express";
import { AuthPaths } from "../../types/paths";
import { AuthService } from "./auth.service";

@Controller(AuthPaths.AUTH)
export class AuthController {
	constructor(private readonly authService: AuthService) {}

	@Post(AuthPaths.LOGIN)
	async login(
		@Res({ passthrough: true }) response: Response,
		@Body() { username, password }: IGetUserResponse
	) {
		const result = await this.authService.login(username, password);
		if (!result.success) return result;

		response.cookie("jwt", result.access_token, {
			maxAge: 1000 * 24 * 3600,
			httpOnly: true,
		});

		return { success: true };
	}
}
