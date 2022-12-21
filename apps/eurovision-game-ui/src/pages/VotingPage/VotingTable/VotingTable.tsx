import { Grid } from "@mui/material";
import { styles } from "./VotingTable.styles";
import { useEffect, useState } from "react";
import { Box } from "@mui/system";
import CountryCard from "./CountryCard/CountryCard";
import { Form, Formik } from "formik";
import { IVotingTableGetResponse } from "utils/global.types";

interface ICountry {
	country: string;
	artist: string;
	song: string;
}

const VotingTable: React.FC = () => {
	const [countries, setCountries] = useState([]);
	const [isLoading, setIsLoading] = useState(countries.length === 0);
	const [isModalOpen, setIsModalOpen] = useState(false);

	// TODO: replace with RTK query
	useEffect(() => {
		fetch("http://localhost:3004/EV22")
			.then((res) => res.json())
			.then((res) => setCountries(res))
			.catch((err) => console.log(err))
			.finally(() => setIsLoading(false));
	}, []);

	const toggleModal = (): void => setIsModalOpen(!isModalOpen);

	// TODO: move this formatting to backend
	// TODO: change to IGetVotesResponse
	const initialValues: IVotingTableGetResponse = countries.reduce(
		(prev, { country, vote }) => ({
			...prev,
			[country]: "",
		}),
		{}
	);

	console.log(initialValues);

	const onSubmit = (values: IVotingTableGetResponse): void => {
		console.log(values);
	};

	return isLoading ? (
		// TODO: replace with spinner or skeleton
		<Box>Loading</Box>
	) : (
		<Formik initialValues={initialValues} onSubmit={onSubmit}>
			{({ submitForm }): React.ReactNode => (
				<Form>
					<Grid container sx={styles.tableContainer}>
						{countries.map((entry: ICountry) => (
							<Grid
								item
								md={4}
								sm={6}
								xs={12}
								key={entry.country}
								sx={styles.card}
							>
								<CountryCard
									country={entry.country}
									artist={entry.artist}
									song={entry.song}
									toggleModal={toggleModal}
									submitForm={submitForm}
								/>
							</Grid>
						))}
					</Grid>
				</Form>
			)}
		</Formik>
	);
};

export default VotingTable;
