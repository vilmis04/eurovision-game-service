import { Module } from "@nestjs/common";
import { AuthModule } from "../modules/auth/auth.module";
import { VotesModule } from "../modules/votes/votes.module";

@Module({
	imports: [VotesModule, AuthModule],
})
export class AppModule {}
