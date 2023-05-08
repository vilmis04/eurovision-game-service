import { Injectable, Req, Res } from "@nestjs/common";
import { RepoClient } from "../../utils/RepoClient";
import * as bcrypt from "bcrypt";
import { Response, Request } from "express";
import { JwtUtils } from "../../utils/JwtUtils";
import validator from "validator";
import { GroupService } from "../group/group.service";

export interface ILoginResponse {
	success: boolean;
	errors?: string[];
	access_token?: string;
}

@Injectable()
export class AuthService {
	constructor(
		private readonly repoClient: RepoClient,
		private readonly jwtUtils: JwtUtils,
		private readonly groupService: GroupService
	) {}

	async login(
		@Res({ passthrough: true }) response: Response,
		username: string,
		enteredPassword: string
	): Promise<ILoginResponse> {
		const user = await this.repoClient.getUserByUsername(username);

		if (!user) {
			response.status(400).send();
			throw new Error("Incorrect username");
		}
		if (!this.validateUser(enteredPassword, user.password)) {
			response.status(400).send();
			throw new Error("Incorrect password");
		}

		const { password, ...userData } = user;
		const access_token = await this.jwtUtils.generateToken(userData);

		response.cookie("jwt", access_token, {
			maxAge: 1000 * 24 * 3600,
			httpOnly: true,
		});

		response.cookie("username", username, {
			maxAge: 1000 * 24 * 3600,
		});

		return { success: true };
	}

	async validateUser(enteredPassword: string, userPassword: string) {
		return await bcrypt.compare(enteredPassword, userPassword);
	}

	async logout(@Res({ passthrough: true }) response: Response) {
		response.cookie("jwt", "", {
			maxAge: 1,
			httpOnly: true,
		});

		response.cookie("username", "", {
			maxAge: 1,
		});
	}

	async getRoles(@Req() request: Request) {
		const { roles } = this.jwtUtils.getUser(request);
		return roles ?? [];
	}

	async getAuthStatus(
		@Req() request: Request,
		@Res({ passthrough: true }) response: Response
	) {
		const isAuthenticated = this.jwtUtils.getAuthStatus(request);

		if (isAuthenticated) {
			await this.groupService.joinGroup(request, response);
		}

		return isAuthenticated;
	}

	async signUp(
		@Req() request: Request,
		@Res({ passthrough: true }) res: Response,
		username: string,
		enteredPassword: string
	): Promise<ILoginResponse> {
		const isUsernameTaken = Boolean(
			await this.repoClient.getUserByUsername(username)
		);

		if (isUsernameTaken) {
			res.status(400).send();
			throw new Error("Username already exists");
		}

		const salt = bcrypt.genSaltSync(10);
		const hash = bcrypt.hashSync(enteredPassword, salt);

		const { acknowledged: userAcknowledged, insertedId: userId } =
			await this.repoClient.createUser(validator.escape(username), hash);

		if (!userAcknowledged) {
			res.status(500).send();
			throw new Error("Internal server error");
		}

		const user = await this.repoClient.getUserById(userId);

		if (!user) {
			res.status(500).send();
			throw new Error("Internal server error");
		}

		const loggedinUser = await this.login(
			res,
			user.username,
			enteredPassword
		);

		await this.groupService.joinGroup(request, res);

		return loggedinUser;
	}
}
