import { PartialType } from "@nestjs/mapped-types";
import { Country } from "../entities/country.entity";

export class UpdateCountryDto extends PartialType(Country) {
	name: Country["name"];
	year: Country["year"];
}
