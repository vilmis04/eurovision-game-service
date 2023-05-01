import { IGetUserResponse } from "@eurovision-game-monorepo/core";
import { Body, Controller, Get, Param, Post, Req, Res } from "@nestjs/common";
import { Response, Request } from "express";
import { AuthPaths, RootPaths } from "../../types/paths";
import { AuthService } from "./auth.service";
import { GroupService } from "../group/group.service";

@Controller(AuthPaths.AUTH)
export class AuthController {
	constructor(
		private readonly authService: AuthService,
		private readonly groupService: GroupService
	) {}

	@Post(AuthPaths.LOGIN)
	async login(
		@Res({ passthrough: true }) response: Response,
		@Body() { username, password }: IGetUserResponse
	) {
		return await this.authService.login(response, username, password);
	}

	@Post(AuthPaths.LOGOUT)
	async logout(@Res({ passthrough: true }) response: Response) {
		return await this.authService.logout(response);
	}

	@Get(AuthPaths.ROLES)
	async getRoles(@Req() request: Request) {
		return await this.authService.getRoles(request);
	}

	@Get(`${RootPaths.GROUPS}/join/:token`)
	async addGroupToJoin(
		@Res({ passthrough: true }) response: Response,
		@Param("token") token: string
	) {
		return await this.groupService.addGroupToJoin(response, token);
	}

	@Get()
	async getAuthStatus(
		@Req() request: Request,
		@Res({ passthrough: true }) response: Response
	) {
		return await this.authService.getAuthStatus(request, response);
	}

	@Post("signup")
	async signUp(
		@Res({ passthrough: true }) response: Response,
		@Req() request: Request,
		@Body() { username, password }: IGetUserResponse
	) {
		const result = await this.authService.signUp(
			request,
			response,
			username,
			password
		);
		if (!result.success) return result;

		return { success: true };
	}
}
