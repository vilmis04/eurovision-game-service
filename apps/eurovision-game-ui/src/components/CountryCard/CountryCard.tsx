import { Card, Grid, Typography } from "@mui/material";
import VoteDropdownMenu from "components/VoteDropdownMenu/VoteDropdownMenu";
import { styles } from "./CountryCard.styles";

interface ICountryCardProps {
	country: string;
	artist: string;
	song: string;
	toggleModal: () => void;
}

const CountryCard: React.FC<ICountryCardProps> = ({
	country,
	artist,
	song,
	toggleModal,
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
				<VoteDropdownMenu />
			</Grid>
		</Grid>
	</Card>
);

export default CountryCard;
