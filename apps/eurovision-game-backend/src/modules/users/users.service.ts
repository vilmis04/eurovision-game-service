import { Injectable, Res } from "@nestjs/common";
import { RepoClient } from "../../utils/RepoClient";
import { AuthService, ILoginResponse } from "../auth/auth.service";
import * as bcrypt from "bcrypt";
import { Response } from "express";
import validator from "validator";

interface ISignUpResponse extends ILoginResponse {}

@Injectable()
export class UsersService {
	constructor(
		private readonly repoClient: RepoClient,
		private readonly authService: AuthService
	) {}

	async signUp(
		@Res({ passthrough: true }) res: Response,
		username: string,
		enteredPassword: string
	): Promise<ISignUpResponse> {
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

		const loggedinUser = await this.authService.login(
			res,
			user.username,
			enteredPassword
		);

		return loggedinUser;
	}
}
