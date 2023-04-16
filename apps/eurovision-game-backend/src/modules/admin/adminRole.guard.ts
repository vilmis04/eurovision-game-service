import { Injectable, CanActivate, ExecutionContext } from "@nestjs/common";
import { Observable } from "rxjs";
import { Request } from "express";
import { JwtUtils } from "../../utils/JwtUtils";
import { RoleTypes } from "@eurovision-game-monorepo/core";

@Injectable()
export class AdminRoleGuard implements CanActivate {
	constructor(private readonly jwtUtils: JwtUtils) {}
	canActivate(
		context: ExecutionContext
	): boolean | Promise<boolean> | Observable<boolean> {
		const request: Request = context.switchToHttp().getRequest();
		const { roles } = this.jwtUtils.getUser(request);

		return roles.includes(RoleTypes.ADMIN);
	}
}
