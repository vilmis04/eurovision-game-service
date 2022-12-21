import { Card, Grid, Typography } from "@mui/material";
import VoteDropdownMenu from "./VoteDropdownMenu/VoteDropdownMenu";
import { styles } from "./CountryCard.styles";

export interface ICountryCardProps {
	country: string;
	artist: string;
	song: string;
	toggleModal: () => void;
	submitForm: () => Promise<void>;
}
// TODO: count already selected votes and allow to choose only available votes
const CountryCard: React.FC<ICountryCardProps> = ({
	country,
	artist,
	song,
	submitForm,
}) => (
	<Card variant="outlined" sx={styles.card}>
		<Grid container direction="column">
			<Grid item>
				<Typography variant="overline">{country}</Typography>
			</Grid>
			<Grid container item>
				<Grid item>
					<Typography variant="subtitle1">{artist}</Typography>
				</Grid>
				<Grid item sx={styles.gap}>
					<Typography variant="subtitle1" sx={styles.song}>
						{song}
					</Typography>
				</Grid>
			</Grid>
			<Grid item sx={styles.button}>
				<VoteDropdownMenu country={country} submitForm={submitForm} />
			</Grid>
		</Grid>
	</Card>
);

export default CountryCard;
