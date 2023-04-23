import { Injectable } from "@nestjs/common";
import { RepoClient } from "../../utils/RepoClient";
import { Admin } from "./entities/admin.entity";
import { CountryService } from "../country/country.service";
import { GameTypes } from "@eurovision-game-monorepo/core";

@Injectable()
export class AdminService {
	constructor(
		private readonly repoClient: RepoClient,
		private readonly countryService: CountryService
	) {}

	async getAdminConfig() {
		return await this.repoClient.getAdminConfig();
	}

	async updateAdminConfig(adminConfig: Admin) {
		return this.repoClient.updateAdminConfig(adminConfig);
	}

	async promoteToFinal() {
		const { type, year } = await this.getAdminConfig();
		const semiFinalists = await this.countryService.findCountries({
			type,
			year,
		});
		const finalists = semiFinalists
			.filter(({ semiFinalScore }) => semiFinalScore)
			.map((country) => {
				const { _id, ...restCountry } = country;
				return { ...restCountry, type: GameTypes.FINAL };
			});

		await this.repoClient.createManyCountries(finalists);
	}

	async submitFinalScore() {
		// TODO: add submit final score logic
	}
}
