import { Controller, Get, Body, Patch, UseGuards, Post } from "@nestjs/common";
import { AdminService } from "./admin.service";
import { AdminPaths, RootPaths } from "../../types/paths";
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

	@UseGuards(AdminRoleGuard)
	@Post(AdminPaths.PROMOTE)
	async promoteToFinal() {
		return await this.adminService.promoteToFinal();
	}

	@UseGuards(AdminRoleGuard)
	@Post(AdminPaths.SUBMIT_FINAL)
	async submitFinalScore() {
		return await this.adminService.submitFinalScore();
	}
}
