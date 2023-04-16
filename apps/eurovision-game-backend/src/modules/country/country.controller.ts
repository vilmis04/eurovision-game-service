import {
	Controller,
	Get,
	Post,
	Body,
	Patch,
	Param,
	Delete,
} from "@nestjs/common";
import { CountryService } from "./country.service";
import { UpdateCountryDto } from "./dto/update-country.dto";
import { CreateCountryDto } from "./dto/create-country.dto";
import { GameTypes } from "@eurovision-game-monorepo/core";

@Controller("country")
export class CountryController {
	constructor(private readonly countryService: CountryService) {}

	@Post()
	create(@Body() createCountryDto: CreateCountryDto) {
		return this.countryService.create(createCountryDto);
	}

	@Get(":year/:type")
	findAll(@Param("type") type: GameTypes, @Param("year") year: string) {
		return this.countryService.findCountries({ year, type });
	}

	@Patch()
	update(@Body() updateCountryDto: UpdateCountryDto) {
		return this.countryService.update(updateCountryDto);
	}

	@Get(":year/:name")
	findOne(@Param("year") year: string, @Param("name") name: string) {
		return this.countryService.findOne(year, name);
	}

	@Delete(":year/:name")
	remove(@Param("year") year: string, @Param("name") name: string) {
		return this.countryService.remove(year, name);
	}
}
