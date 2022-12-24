import { Card, Grid, Typography } from "@mui/material";
import VoteDropdownMenu from "./VoteDropdownMenu/VoteDropdownMenu";
import { styles } from "./CountryCard.styles";
import { countryData } from "./CountryCard.config";
import { CountryTypes } from "@eurovision-game-monorepo/core";

export interface ICountryCardProps {
	country: CountryTypes;
	submitForm: () => Promise<void>;
}
// TODO: count already selected votes and allow to choose only available votes
const CountryCard: React.FC<ICountryCardProps> = ({ country, submitForm }) => {
	const data = countryData[country];

	return (
		<Card variant="outlined" sx={styles.card}>
			<Grid container direction="column">
				<Grid item>
					<Typography variant="overline">{country}</Typography>
				</Grid>
				<Grid container item>
					<Grid item>
						<Typography variant="subtitle1">
							{data && data.artist}
						</Typography>
					</Grid>
					<Grid item sx={styles.gap}>
						<Typography variant="subtitle1" sx={styles.song}>
							{data && data.song}
						</Typography>
					</Grid>
				</Grid>
				<Grid item sx={styles.button}>
					<VoteDropdownMenu
						country={country}
						submitForm={submitForm}
					/>
				</Grid>
			</Grid>
		</Card>
	);
};

export default CountryCard;
