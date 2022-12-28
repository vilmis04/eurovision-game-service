import { Body, Controller, Post, Req } from "@nestjs/common";
import { AuthPaths } from "../../types/paths";

interface ILoginBody {
	username: string;
	password: string;
}

@Controller(AuthPaths.AUTH)
export class AuthController {
	@Post(AuthPaths.AUTH_LOGIN)
	async login(@Req() request: Request, @Body() body: ILoginBody) {
		console.log(request);
		console.log(body);
	}
}
