import { Card, Grid, Typography } from "@mui/material";
import VoteDropdownMenu from "./VoteDropdownMenu/VoteDropdownMenu";
import { styles } from "./CountryCard.styles";
import VoteSemiCheckbox from "./VoteSemiCheckbox/VoteSemiCheckbox";
import { IAdminFormData } from "@eurovision-game-monorepo/core";

export interface ICountryCardProps
	extends Pick<IAdminFormData, "isVotingDisabled"> {
	country: string;
	artist: string;
	song: string;
	isFinal: boolean;
}
// TODO: count already selected votes and allow to choose only available votes
const CountryCard: React.FC<ICountryCardProps> = ({
	country,
	artist,
	song,
	isFinal,
	isVotingDisabled,
}) => {
	return (
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
				{/* TODO: fix sx syntax to accept more than one class */}
				<Grid item sx={isFinal ? styles.button : {}}>
					{isFinal ? (
						<VoteDropdownMenu
							country={country}
							isVotingDisabled={isVotingDisabled}
						/>
					) : (
						<VoteSemiCheckbox
							country={country}
							isVotingDisabled={isVotingDisabled}
						/>
					)}
				</Grid>
			</Grid>
		</Card>
	);
};

export default CountryCard;
