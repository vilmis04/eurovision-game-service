import { Module } from "@nestjs/common";
import { CountryService } from "./country.service";
import { CountryController } from "./country.controller";
import { RepoClient } from "../../utils/RepoClient";

@Module({
	controllers: [CountryController],
	providers: [CountryService, RepoClient],
})
export class CountryModule {}
