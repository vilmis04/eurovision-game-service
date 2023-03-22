import { IScoreFormData } from "@eurovision-game-monorepo/core";
import { Button, Grid } from "@mui/material";
import { Form, Formik, FormikState } from "formik";
import { useEffect, useState } from "react";
import { styles } from "../AdminPage.styles";
import { v4 as uuid } from "uuid";
import { TCountriesToAdd } from "./ScoreForm.types";
import { mapCountryChangesToValues } from "./ScoreForm.utils";
import CountryEntry from "./CountryEntry/CountryEntry";

// TODO: CONTINUE
/*
3. integrate with backend
*/

interface IScoreFormProps {
	initialValues: IScoreFormData;
	handleScoreFormSubmit: (values: IScoreFormData) => void;
}

const ScoreForm: React.FC<IScoreFormProps> = ({
	initialValues,
	handleScoreFormSubmit,
}) => {
	const [isEditMode, setIsEditMode] = useState(false);
	const [countries, setCountries] = useState<string[]>([]);
	const [countriesToAdd, setCountriesToAdd] = useState<TCountriesToAdd>({});

	useEffect(() => {
		setCountries(Object.keys(initialValues.countries));
	}, [initialValues]);

	const toggleEditMode = () => setIsEditMode((s) => !s);

	const handleCancel = (
		resetForm: (next?: Partial<FormikState<IScoreFormData>>) => void
	) => {
		resetForm();
		setIsEditMode(false);
		setCountries(Object.keys(initialValues.countries));
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

	const handleSubmit = (values: IScoreFormData): void => {
		const updatedValues = mapCountryChangesToValues({
			countries,
			countriesToAdd,
			values,
		});

		handleScoreFormSubmit(updatedValues);
	};

	return (
		<Formik initialValues={initialValues} onSubmit={handleSubmit}>
			{({ dirty, resetForm }) => (
				<Form>
					<Grid item container sx={styles.container}>
						{countries.map((country) => (
							<CountryEntry
								key={country}
								country={country}
								isEditMode={isEditMode}
								removeCountry={() => removeCountry(country)}
							/>
						))}
						{Object.entries(countriesToAdd).map(([id]) => (
							<CountryEntry
								key={id}
								country={id}
								isEditMode={isEditMode}
								removeCountry={() => removeCountryFromToAdd(id)}
								isCountriesToAdd
								updateCountryObject={updateCountryObject}
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
								{isEditMode ? "Add country" : "Edit countries"}
							</Button>
						</Grid>
						<Grid container sx={styles.cancelButtonContainer}>
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
						<Grid container sx={styles.cancelButtonContainer}>
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
	);
};

export default ScoreForm;
