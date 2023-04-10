import { Injectable } from "@nestjs/common";
import { UpdateCountryDto } from "./dto/update-country.dto";
import { RepoClient } from "../../utils/RepoClient";
import { CreateCountryDto } from "./dto/create-country.dto";
import { Country } from "./entities/country.entity";
import { GameTypes } from "@eurovision-game-monorepo/core";

@Injectable()
export class CountryService {
	constructor(private readonly repoClient: RepoClient) {}

	create(createCountryDto: CreateCountryDto) {
		const requestBody: Country = {
			...createCountryDto,
			semiFinalScore: false,
			finalScore: 0,
		};
		return this.repoClient.createCountry(requestBody);
	}

	async findCountries({ year, type }: { year: string; type: GameTypes }) {
		const countryList = await this.repoClient.findCountries(year, type);
		console.log({ countryList });

		return countryList;
	}

	async update(updateCountryDto: UpdateCountryDto) {
		const { name, year } = updateCountryDto;
		const countryToUpdate = await this.findOne(year, name);

		return this.repoClient.updateCountry({
			...countryToUpdate,
			...updateCountryDto,
		});
	}

	findOne(year: string, name: string) {
		return this.repoClient.getOneCountry(year, name);
	}

	remove(year: string, name: string) {
		return this.repoClient.removeCountry(year, name);
	}
}
