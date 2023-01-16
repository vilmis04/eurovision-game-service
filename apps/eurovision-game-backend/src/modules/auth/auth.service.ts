import { IUserDataForToken } from "@eurovision-game-monorepo/core";
import { Injectable, Res } from "@nestjs/common";
import { RepoClient } from "../../utils/RepoClient";
import * as jwt from "jsonwebtoken";
import * as bcrypt from "bcrypt";
import { Response } from "express";

export interface ILoginResponse {
	success: boolean;
	errors?: string[];
	access_token?: string;
}

@Injectable()
export class AuthService {
	constructor(private readonly repoClient: RepoClient) {}

	async login(
		@Res({ passthrough: true }) res: Response,
		username: string,
		enteredPassword: string
	): Promise<ILoginResponse> {
		const user = await this.repoClient.getUserByUsername(username);

		if (!user) {
			res.status(400).send();
			throw new Error("Incorrect username");
		}
		if (!this.validateUser(enteredPassword, user.password)) {
			res.status(400).send();
			throw new Error("Incorrect password");
		}

		const { password, ...userData } = user;
		const access_token = await this.generateToken(userData);

		return { success: true, access_token };
	}

	async validateUser(enteredPassword: string, userPassword: string) {
		return await bcrypt.compare(enteredPassword, userPassword);
	}

	async generateToken(user: IUserDataForToken) {
		const MAX_AGE = 24 * 3600; // 1 day in seconds
		const access_token = jwt.sign(user, process.env.SECRET_KEY!, {
			expiresIn: MAX_AGE,
		});
		return access_token;
	}
}
