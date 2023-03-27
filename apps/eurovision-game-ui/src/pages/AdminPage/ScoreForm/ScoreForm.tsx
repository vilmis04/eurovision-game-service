import {
	GameTypes,
	IScoreFormData,
	isResponseError,
} from "@eurovision-game-monorepo/core";
import { Button, Grid } from "@mui/material";
import { Form, Formik, FormikHelpers, FormikState } from "formik";
import { useEffect, useState } from "react";
import { v4 as uuid } from "uuid";
import { TCountriesToAdd } from "./ScoreForm.types";
import { mapCountryChangesToValues } from "./ScoreForm.utils";
import CountryEntry from "./CountryEntry/CountryEntry";
import { styles } from "./ScoreForm.styles";
import {
	useGetScoreQuery,
	useUpdateScoreMutation,
} from "../@modules/score.api";

interface IScoreFormProps {
	type: GameTypes;
	year: string;
	setIsFetching?: React.Dispatch<React.SetStateAction<boolean>>;
}

const ScoreForm: React.FC<IScoreFormProps> = ({
	type,
	year,
	setIsFetching,
}) => {
	const [isEditMode, setIsEditMode] = useState(false);
	const [countries, setCountries] = useState<string[]>([]);
	const [countriesToAdd, setCountriesToAdd] = useState<TCountriesToAdd>({});

	const { data: scoreData, isFetching } = useGetScoreQuery({ type, year });

	const [updateScore] = useUpdateScoreMutation({ fixedCacheKey: "score" });

	useEffect(() => {
		scoreData && setCountries(Object.keys(scoreData.countries));
	}, [scoreData?.countries]);

	useEffect(() => {
		setIsFetching && setIsFetching(isFetching);
	}, [isFetching]);

	const toggleEditMode = () => setIsEditMode((s) => !s);

	const handleCancel = (
		resetForm: (next?: Partial<FormikState<IScoreFormData>>) => void
	) => {
		resetForm();
		setIsEditMode(false);
		scoreData && setCountries(Object.keys(scoreData.countries));
		setCountriesToAdd({});
	};

	const removeCountry = (countryToRemove: string) => {
		setCountries(
			countries.filter((country) => country !== countryToRemove)
		);
	};

	const removeCountryFromToAdd = (index: string) => {
		setCountriesToAdd((list) => {
			const { [index]: _, ...restCountries } = list;
			return restCountries;
		});
	};

	const addCountry = () => {
		setCountriesToAdd((list) => ({
			...list,
			[uuid()]: { name: "", score: 0 },
		}));
	};

	const handleCountryButtonClick = isEditMode ? addCountry : toggleEditMode;

	const updateCountryObject = (
		id: string,
		object: { name?: string; score?: number }
	) => {
		setCountriesToAdd((list) => ({
			...list,
			[id]: { ...list[id], ...object },
		}));
	};

	const handleSubmit = async (
		values: IScoreFormData,
		{ setFieldValue }: FormikHelpers<IScoreFormData>
	): Promise<void> => {
		const updatedValues = mapCountryChangesToValues({
			countries,
			countriesToAdd,
			values,
		});
		setIsEditMode(false);
		setCountriesToAdd({});
		setFieldValue("countries", updatedValues.countries);

		const response = await updateScore(updatedValues);

		if (isResponseError(response)) {
			// TODO: do something with an error
			console.log(response.error);
			return;
		}
	};

	return (
		<>
			{scoreData && (
				<Formik initialValues={scoreData} onSubmit={handleSubmit}>
					{({ dirty, resetForm }) => (
						<Form>
							<Grid item container sx={styles.container}>
								{countries.map((country) => (
									<CountryEntry
										key={country}
										country={country}
										isEditMode={isEditMode}
										removeCountry={() =>
											removeCountry(country)
										}
									/>
								))}
								{Object.entries(countriesToAdd).map(([id]) => (
									<CountryEntry
										key={id}
										country={id}
										isEditMode={isEditMode}
										removeCountry={() =>
											removeCountryFromToAdd(id)
										}
										isCountriesToAdd
										updateCountryObject={
											updateCountryObject
										}
									/>
								))}
								<Grid container sx={styles.saveButtonContainer}>
									<Button
										color="secondary"
										onClick={handleCountryButtonClick}
										variant="contained"
										fullWidth
										sx={styles.inputField}
									>
										{isEditMode
											? "Add country"
											: "Edit countries"}
									</Button>
								</Grid>
								<Grid
									container
									sx={styles.cancelButtonContainer}
								>
									<Button
										type="submit"
										disabled={!dirty && !isEditMode}
										variant="contained"
										fullWidth
										sx={styles.inputField}
									>
										Save
									</Button>
								</Grid>
								<Grid
									container
									sx={styles.cancelButtonContainer}
								>
									<Button
										onClick={() => handleCancel(resetForm)}
										disabled={!dirty && !isEditMode}
										variant="outlined"
										fullWidth
										sx={styles.inputField}
									>
										Cancel
									</Button>
								</Grid>
							</Grid>
						</Form>
					)}
				</Formik>
			)}
		</>
	);
};

export default ScoreForm;
