import { IUserDataForToken } from "@eurovision-game-monorepo/core";
import { Injectable } from "@nestjs/common";
import { RepoClient } from "../../utils/RepoClient";
import * as jwt from "jsonwebtoken";

interface ILoginResponse {
	success: boolean;
	errors?: string[];
	access_token?: string;
}

@Injectable()
export class AuthService {
	constructor(private readonly repoClient: RepoClient) {}
	async login(
		username: string,
		enteredPassword: string
	): Promise<ILoginResponse> {
		// TODO: sanitize inputs
		const user = await this.repoClient.getUserByUsername(username);

		// TODO: add error codes (?)
		if (!user) {
			return { success: false, errors: ["no user found"] };
		}
		if (!this.validateUser(enteredPassword, user.password)) {
			return { success: false, errors: ["invalid password"] };
		}

		const { password, ...userData } = user;
		const access_token = await this.generateToken(userData);

		return { success: true, access_token };
	}

	async validateUser(enteredPassword: string, userPassword: string) {
		// TODO: add bcrypt here
		return enteredPassword === userPassword;
	}

	async generateToken(user: IUserDataForToken) {
		const MAX_AGE = 24 * 3600; // 1 day in seconds
		const access_token = jwt.sign(user, process.env.SECRET_KEY!, {
			expiresIn: MAX_AGE,
		});
		return access_token;
	}
}
