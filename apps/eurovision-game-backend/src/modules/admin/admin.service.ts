import { Injectable } from "@nestjs/common";
import { RepoClient } from "../../utils/RepoClient";
import { Admin } from "./entities/admin.entity";

@Injectable()
export class AdminService {
	constructor(private repoClient: RepoClient) {}

	async getAdminConfig() {
		return await this.repoClient.getAdminConfig();
	}

	async updateAdminConfig(adminConfig: Admin) {
		return this.repoClient.updateAdminConfig(adminConfig);
	}
}
