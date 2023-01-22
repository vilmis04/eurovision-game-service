import { HttpException, HttpStatus, Injectable } from "@nestjs/common";
import * as jwt from "jsonwebtoken";
import { Request } from "express";

const SECRET_KEY = process.env.SECRET_KEY || "";

interface IUserFromToken {
	username: string;
}

@Injectable()
export class JwtUtils {
	public getUsername(req: Request): string {
		const token = req.cookies.jwt;
		const user = jwt.verify(token, SECRET_KEY) as IUserFromToken | null;

		if (user == null) {
			throw new HttpException("Unauthorized", HttpStatus.UNAUTHORIZED);
		}

		return user.username;
	}
}
