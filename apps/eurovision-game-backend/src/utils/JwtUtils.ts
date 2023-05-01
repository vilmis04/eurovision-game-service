import { HttpException, HttpStatus, Injectable } from "@nestjs/common";
import * as jwt from "jsonwebtoken";
import { Request } from "express";
import { IUser } from "@eurovision-game-monorepo/core";

const SECRET_KEY = process.env.SECRET_KEY || "";

interface IUserFromToken extends Omit<IUser, "password"> {}

@Injectable()
export class JwtUtils {
	public getUser(req: Request): IUserFromToken {
		const token = req.cookies.jwt;
		const user = jwt.verify(token, SECRET_KEY) as IUserFromToken | null;

		if (user == null) {
			throw new HttpException("Unauthorized", HttpStatus.UNAUTHORIZED);
		}

		return user;
	}

	public verifyToken(token: string): IUserFromToken {
		const user = jwt.verify(token, SECRET_KEY) as IUserFromToken | null;

		if (user == null) {
			throw new HttpException("Unauthorized", HttpStatus.UNAUTHORIZED);
		}

		return user;
	}

	async generateToken(data: string | object, maxAge: number = 24 * 3600) {
		const token = jwt.sign(data, SECRET_KEY, {
			expiresIn: maxAge,
		});

		return token;
	}

	public async encryptLink(id: string) {
		return await this.generateToken({ id });
	}

	public async decryptLink(token: string) {
		const linkData = jwt.verify(token, SECRET_KEY) as { id: string };
		return linkData.id;
	}

	public getAuthStatus(req: Request): boolean {
		const token = req.cookies?.jwt;
		if (!token) return false;
		const user = jwt.verify(token, SECRET_KEY) as IUserFromToken | null;
		if (!user) return false;

		return true;
	}
}
