import { Controller, Get, Body, Patch } from "@nestjs/common";
import { AdminService } from "./admin.service";
import { RootPaths } from "../../types/paths";
import { Admin } from "./entities/admin.entity";

// TODO: add a RoleGuard
@Controller(RootPaths.ADMIN)
export class AdminController {
	constructor(private readonly adminService: AdminService) {}

	@Get()
	async getAdminConfig() {
		return await this.adminService.getAdminConfig();
	}

	@Patch()
	async updateAdminConfig(@Body() body: Admin) {
		return await this.adminService.updateAdminConfig(body);
	}
}
