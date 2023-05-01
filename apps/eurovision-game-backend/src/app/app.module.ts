import { MiddlewareConsumer, Module } from "@nestjs/common";
import { AdminModule } from "../modules/admin/admin.module";
import { AuthMiddleware } from "../modules/auth/auth.middleware";
import { AuthModule } from "../modules/auth/auth.module";
import { UsersModule } from "../modules/users/users.module";
import { VotesModule } from "../modules/votes/votes.module";
import { RootPaths } from "../types/paths";
import { CountryModule } from "../modules/country/country.module";
import { GroupModule } from "../modules/group/group.module";

@Module({
	imports: [
		VotesModule,
		AuthModule,
		UsersModule,
		AdminModule,
		CountryModule,
		GroupModule,
	],
})
export class AppModule {
	configure(consumer: MiddlewareConsumer) {
		consumer
			.apply(AuthMiddleware)
			.forRoutes(
				RootPaths.VOTES,
				RootPaths.ADMIN,
				RootPaths.GROUPS,
				RootPaths.USERS
			);
	}
}
