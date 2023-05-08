// TODO: make it as close to a production server as possible
/**
 * This is not a production server yet!
 * This is only a minimal backend to get started.
 */

import { Logger } from "@nestjs/common";
import { NestFactory } from "@nestjs/core";

import { AppModule } from "./app/app.module";
import * as cookieParser from "cookie-parser";

async function bootstrap() {
	const app = await NestFactory.create(AppModule);
	const globalPrefix = "api";
	app.setGlobalPrefix(globalPrefix);
	const port = process.env.PORT || 3333;
	app.use(cookieParser());
	app.enableCors({
		credentials: true,
		// TODO: add real origins (after hosting)
		origin: [
			process.env.BASE_URL || "",
			"http://localhost:3000",
			"http://127.0.0.1:3000",
			"http://192.168.1.205:3000",
		],
	});
	await app.listen(port);
	Logger.log(
		`🚀 Application is running on: http://localhost:${port}/${globalPrefix}`
	);
}

bootstrap();
