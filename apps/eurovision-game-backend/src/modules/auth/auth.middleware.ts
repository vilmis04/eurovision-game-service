import {
	HttpException,
	HttpStatus,
	Injectable,
	NestMiddleware,
} from "@nestjs/common";
import { Request, Response, NextFunction } from "express";
import * as jwt from "jsonwebtoken";

const SECRET_KEY = process.env.SECRET_KEY!;

@Injectable()
export class AuthMiddleware implements NestMiddleware {
	use(req: Request, res: Response, next: NextFunction) {
		const jwtToken = req.cookies?.jwt;

		if (jwtToken == null) {
			throw new HttpException(
				"User not authorized!",
				HttpStatus.UNAUTHORIZED
			);
		}

		const user = jwt.verify(jwtToken, SECRET_KEY);
		if (user == null) {
			throw new HttpException(
				"User not authorized!",
				HttpStatus.UNAUTHORIZED
			);
		}

		next();
	}
}
