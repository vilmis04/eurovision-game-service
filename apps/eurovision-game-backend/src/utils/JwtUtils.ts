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
}
