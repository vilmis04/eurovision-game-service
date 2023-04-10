import { GameTypes } from "@eurovision-game-monorepo/core";
import { CircularProgress, Grid } from "@mui/material";
import { useGetCountryListQuery } from "../@modules/country.api";

interface ICountryConfigFormProps {
	type: GameTypes;
	year: string;
}

const CountryConfigForm: React.FC<ICountryConfigFormProps> = ({
	year,
	type,
}) => {
	const { data: countryList } = useGetCountryListQuery({ type, year });

	console.log({ countryList });
	// TODO: add a spinner until the results
	return (
		<>{countryList ? <Grid>Render cards</Grid> : <CircularProgress />}</>
	);
};

export default CountryConfigForm;
