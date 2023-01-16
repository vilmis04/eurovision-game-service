import { MiddlewareConsumer, Module } from "@nestjs/common";
import { AuthMiddleware } from "../modules/auth/auth.middleware";
import { AuthModule } from "../modules/auth/auth.module";
import { UsersModule } from "../modules/users/users.module";
import { VotesController } from "../modules/votes/votes.controller";
import { VotesModule } from "../modules/votes/votes.module";

@Module({
	imports: [VotesModule, AuthModule, UsersModule],
})
export class AppModule {
	configure(consumer: MiddlewareConsumer) {
		consumer.apply(AuthMiddleware).forRoutes(VotesController);
	}
}
