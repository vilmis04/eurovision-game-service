import { Button, Grid, TextField, Typography } from "@mui/material";
import FormField from "apps/eurovision-game-ui/src/components/FormField/FormField";
import { styles } from "./CountryEntry.styles";

interface ICountryEntryProps {
	country: string;
	isEditMode: boolean;
	removeCountry: () => void;
	isCountriesToAdd?: boolean;
	updateCountryObject?: (
		id: string,
		object: {
			name?: string;
			score?: number;
		}
	) => void;
}

const defaultUpdateCountryObject = (
	id: string,
	object: {
		name?: string;
		score?: number;
	}
) => {};

const CountryEntry: React.FC<ICountryEntryProps> = ({
	country,
	isEditMode,
	removeCountry,
	isCountriesToAdd = false,
	updateCountryObject = defaultUpdateCountryObject,
}) => (
	<Grid container gap={1} sx={styles.countryPoints}>
		<Grid container item xs={4}>
			{isEditMode && (
				<Button
					sx={styles.button}
					variant="text"
					onClick={removeCountry}
				>
					X
				</Button>
			)}
			{isCountriesToAdd ? (
				<TextField
					onChange={(event) =>
						updateCountryObject(country, {
							name: event.target.value,
						})
					}
				/>
			) : (
				<Typography variant="body1">{country}</Typography>
			)}
		</Grid>
		<Grid item xs={2}>
			{isCountriesToAdd ? (
				<TextField
					type="number"
					onChange={(event) =>
						updateCountryObject(country, {
							score: Number(event.target.value),
						})
					}
				/>
			) : (
				<FormField name={`countries.${country}`} type="number" />
			)}
		</Grid>
	</Grid>
);

export default CountryEntry;
