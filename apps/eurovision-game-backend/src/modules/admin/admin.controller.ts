import { Controller, Get, Body, Patch, UseGuards } from "@nestjs/common";
import { AdminService } from "./admin.service";
import { RootPaths } from "../../types/paths";
import { Admin } from "./entities/admin.entity";
import { AdminRoleGuard } from "./adminRole.guard";

// TODO: add a RoleGuard
@Controller(RootPaths.ADMIN)
export class AdminController {
	constructor(private readonly adminService: AdminService) {}

	@Get()
	async getAdminConfig() {
		return await this.adminService.getAdminConfig();
	}

	@UseGuards(AdminRoleGuard)
	@Patch()
	async updateAdminConfig(@Body() body: Admin) {
		return await this.adminService.updateAdminConfig(body);
	}
}
