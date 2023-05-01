import { Injectable } from "@nestjs/common";
import { RepoClient } from "../../utils/RepoClient";
import { User } from "./entities/user.entity";

@Injectable()
export class UsersService {
	constructor(private readonly repoClient: RepoClient) {}

	async updateUser(username: string, requestBody: Partial<User>) {
		const userToUpdate = await this.repoClient.getUserByUsername(username);
		if (!userToUpdate) throw new Error("No user found");

		return await this.repoClient.updateUser(username, requestBody);
	}
}
