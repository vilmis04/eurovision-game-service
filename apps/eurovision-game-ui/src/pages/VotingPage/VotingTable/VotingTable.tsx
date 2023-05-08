import { Box, CircularProgress, Grid, Typography } from "@mui/material";
import { styles } from "./VotingTable.styles";
import CountryCard from "./CountryCard/CountryCard";
import { Form, Formik } from "formik";
import {
	GameTypes,
	IAdminFormData,
	IVotes,
} from "@eurovision-game-monorepo/core";
import {
	useUpdateVotesMutation,
	useGetVotesQuery,
} from "../../@modules/votes.api";
import { useGetCountryListQuery } from "../../@modules/country.api";
import { TVoteFormData } from "./VotingTable.types";

interface IVotingTableProps
	extends Pick<IVotes, "year" | "type">,
		Pick<IAdminFormData, "isVotingDisabled"> {}

const VotingTable: React.FC<IVotingTableProps> = ({
	type,
	year,
	isVotingDisabled,
}) => {
	// TODO: add and combine isFetching for all queries
	const { data: votes, isFetching } = useGetVotesQuery({ type, year });
	const { data: countryList } = useGetCountryListQuery({ type, year });
	const initialValues =
		votes &&
		countryList?.reduce(
			(result, { name }) => ({
				...result,
				[name]: votes[name] || "",
			}),
			{}
		);

	const isFinal = type === GameTypes.FINAL;

	// TODO: move out of this file
	const title = {
		[GameTypes.FINAL]: "Final",
		[GameTypes.SEMI_1]: "Semi-final I",
		[GameTypes.SEMI_2]: "Semi-final II",
	};

	const [updateVotes] = useUpdateVotesMutation();

	const onSubmit = (values: TVoteFormData): void => {
		updateVotes({ type, year, votes: values });
	};

	return isFetching ? (
		// TODO: replace with spinner or skeleton
		<CircularProgress />
	) : (
		<Box>
			<Box>
				<Typography variant="h1" sx={{ fontSize: "32px" }}>
					{title[type]}
				</Typography>
			</Box>
			{countryList && initialValues && (
				<Formik initialValues={initialValues} onSubmit={onSubmit}>
					{(): React.ReactNode => (
						<Form>
							<Grid container sx={styles.tableContainer}>
								{countryList.map(
									({ name: country, artist, song }) => {
										return (
											<Grid
												item
												md={4}
												sm={6}
												xs={12}
												key={country}
												sx={styles.card}
											>
												<CountryCard
													country={country}
													artist={artist}
													song={song}
													isFinal={isFinal}
													isVotingDisabled={
														isVotingDisabled
													}
												/>
											</Grid>
										);
									}
								)}
							</Grid>
						</Form>
					)}
				</Formik>
			)}
		</Box>
	);
};

export default VotingTable;
