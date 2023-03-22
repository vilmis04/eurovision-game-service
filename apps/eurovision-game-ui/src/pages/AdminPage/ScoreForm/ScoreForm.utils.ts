import { IScoreFormData } from "@eurovision-game-monorepo/core";
import { TCountriesToAdd } from "./ScoreForm.types";

interface IMapCountryChangesToValuesArgs {
	values: IScoreFormData;
	countriesToAdd: TCountriesToAdd;
	countries: string[];
}

export const mapCountryChangesToValues = ({
	countriesToAdd = {},
	values,
	countries,
}: IMapCountryChangesToValuesArgs): IScoreFormData => {
	const { countries: formikCountries } = values;
	const addedCountries = Object.values(countriesToAdd).reduce(
		(list, { name, score }) => ({ ...list, [name]: score }),
		{}
	);
	const reducedCountries = Object.entries(formikCountries).reduce(
		(updatedList, countryEntry) => {
			const [country] = countryEntry;
			if (!countries.includes(country)) {
				return updatedList;
			}

			return {
				...updatedList,
				...{ [country]: values.countries[country] },
			};
		},
		{}
	);
	return {
		...values,
		countries: { ...reducedCountries, ...addedCountries },
	};
};
