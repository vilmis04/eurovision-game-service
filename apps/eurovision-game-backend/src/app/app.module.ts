import { Module } from "@nestjs/common";
import { VotesModule } from "../modules/votes/votes.module";

import { AppController } from "./app.controller";
import { AppService } from "./app.service";

@Module({
	imports: [VotesModule],
	controllers: [AppController],
	providers: [AppService],
})
export class AppModule {}
