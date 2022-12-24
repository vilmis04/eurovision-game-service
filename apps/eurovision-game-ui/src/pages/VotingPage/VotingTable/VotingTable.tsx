import { Box, Grid } from "@mui/material";
import { styles } from "./VotingTable.styles";
import CountryCard from "./CountryCard/CountryCard";
import { Form, Formik } from "formik";
import {
	CountryTypes,
	IGetVotesResponse,
} from "@eurovision-game-monorepo/core";
import { useGetVotesByUsernameQuery } from "../@module/votes.api";

const VotingTable: React.FC = () => {
	const { data: countries, isFetching } = useGetVotesByUsernameQuery();

	const onSubmit = (values: IGetVotesResponse): void => {
		console.log(values);
	};

	return isFetching ? (
		// TODO: replace with spinner or skeleton
		<Box>Loading...</Box>
	) : (
		<Box>
			{countries && (
				<Formik initialValues={countries} onSubmit={onSubmit}>
					{({ submitForm }): React.ReactNode => (
						<Form>
							<Grid container sx={styles.tableContainer}>
								{/* {filterCountries(countries).map((country) => { */}
								{Object.entries(countries).map(([country]) => {
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
												country={
													country as CountryTypes
												}
												submitForm={submitForm}
											/>
										</Grid>
									);
								})}
							</Grid>
						</Form>
					)}
				</Formik>
			)}
		</Box>
	);
};

export default VotingTable;
