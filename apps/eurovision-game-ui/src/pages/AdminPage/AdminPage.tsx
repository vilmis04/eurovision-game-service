import { CountryTypes } from "@eurovision-game-monorepo/core";
import { Grid, Switch, Typography } from "@mui/material";
import { Field, Form, Formik } from "formik";
import { ChangeEvent } from "react";
import FormField from "../../components/FormField/FormField";
import { styles } from "./AdminPage.styles";

type TCountries = {
	[k in CountryTypes]: string | number;
};

interface IAdminFormData extends TCountries {
	isRegistrationEnabled: boolean;
}

const initialValues: IAdminFormData = {
	isRegistrationEnabled: true,
	ukraine: 5,
	united_kingdom: 5,
	spain: 5,
	sweden: 5,
	serbia: 5,
	norway: 5,
	belgium: 5,
	czech_republic: 5,
	romania: 5,
	portugal: 5,
	finland: 5,
	switzerland: 5,
	france: 5,
	armenia: 5,
	italy: 5,
	netherlands: 5,
	germany: 5,
	lithuania: 5,
	azerbaijan: 5,
	greece: 5,
	iceland: 5,
	moldova: 5,
	australia: 5,
	poland: 5,
	estonia: 5,
};

const AdminPage: React.FC = () => {
	const handleSubmit = (values: IAdminFormData) => {
		console.log("Submit these: ", values);
	};

	const { isRegistrationEnabled, ...countries } = initialValues;

	return (
		<Formik initialValues={initialValues} onSubmit={handleSubmit}>
			{({ values, setFieldValue, submitForm }) => {
				console.log(values);
				return (
					<Form>
						<Grid item container sx={styles.container}>
							{/* TODO: add sorting of maybe?*/}
							{Object.entries(countries).map(([country]) => (
								<Grid
									key={country}
									container
									gap={1}
									sx={styles.countryPoints}
								>
									<Grid item xs={4}>
										<Typography variant="body1">
											{country}
										</Typography>
									</Grid>
									<Grid item xs={2}>
										<FormField
											name={country}
											type="number"
											onBlur={submitForm}
										/>
									</Grid>
								</Grid>
							))}
							{/* TODO: add proper divider */}
							<Grid
								container
								sx={styles.field}
								borderTop="1px solid black"
							>
								<Typography>
									Registration posibility:
								</Typography>
								<Field
									name="isRegistrationEnabled"
									as={Switch}
									checked={values.isRegistrationEnabled}
									onChange={(
										_event: ChangeEvent,
										checked: boolean
									) => {
										setFieldValue(
											"isRegistrationEnabled",
											checked ? true : false
										);
										submitForm();
									}}
								/>
							</Grid>
						</Grid>
					</Form>
				);
			}}
		</Formik>
	);
};

export default AdminPage;
