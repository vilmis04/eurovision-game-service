import { Injectable, NestMiddleware, Req, Res } from "@nestjs/common";
import { Request, Response, NextFunction } from "express";
import * as jwt from "jsonwebtoken";

const SECRET_KEY = process.env.SECRET_KEY!;

@Injectable()
export class AuthMiddleware implements NestMiddleware {
	use(
		@Req() req: Request,
		@Res({ passthrough: true }) res: Response,
		next: NextFunction
	) {
		const jwtToken = req.cookies.jwt;

		if (jwtToken == null) res.status(304).send();
		const user = jwt.verify(jwtToken, SECRET_KEY);
		if (user == null) res.status(304).send();
		next();
	}
}
