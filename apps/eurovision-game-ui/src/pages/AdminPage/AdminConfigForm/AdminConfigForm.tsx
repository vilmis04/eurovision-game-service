import { GameTypes, IAdminFormData } from "@eurovision-game-monorepo/core";
import {
	Button,
	Grid,
	MenuItem,
	Select,
	Switch,
	Typography,
} from "@mui/material";
import { Field, FieldProps, Form, Formik, FormikHelpers } from "formik";
import { ChangeEvent } from "react";
import { styles } from "../AdminPage.styles";

interface IAdminConfigFormProps {
	initialValues: IAdminFormData;
	handleAdminConfigUpdate: (
		values: IAdminFormData,
		{ resetForm }: FormikHelpers<IAdminFormData>
	) => Promise<void>;
}

const AdminConfigForm: React.FC<IAdminConfigFormProps> = ({
	initialValues,
	handleAdminConfigUpdate,
}) => (
	<Formik initialValues={initialValues} onSubmit={handleAdminConfigUpdate}>
		{({ values, setFieldValue, dirty, resetForm }) => (
			<Form>
				<Grid item container sx={styles.container}>
					{/* TODO: add proper divider */}
					<Grid
						container
						sx={styles.field} // TODO: move to a separate file
						borderTop="1px solid black"
					>
						<Typography>Registration posibility:</Typography>
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
							}}
						/>
					</Grid>
					<Grid
						container
						// TODO: move styling
						sx={{
							margin: "0 auto",
							width: "250px",
						}}
					>
						<Grid container>
							<Typography>Year:</Typography>
						</Grid>
						<Grid item container>
							<Field name="year">
								{({ field }: FieldProps<string>) => (
									<>
										<Select
											fullWidth
											sx={styles.inputField}
											{...field}
										>
											{/* TODO: get options from database */}
											<MenuItem value="2022">
												2022
											</MenuItem>
											<MenuItem value="2023">
												2023
											</MenuItem>
										</Select>
									</>
								)}
							</Field>
						</Grid>
					</Grid>
					<Grid
						container
						sx={{
							margin: "0 auto",
							width: "250px",
						}}
					>
						<Grid container>
							<Typography>Game type:</Typography>
						</Grid>
						<Grid item container>
							<Field name="type">
								{({ field }: FieldProps<GameTypes>) => (
									<>
										<Select
											fullWidth
											sx={styles.inputField}
											{...field}
										>
											{Object.values(GameTypes).map(
												(type) => (
													<MenuItem
														key={type}
														value={type}
													>
														{type}
													</MenuItem>
												)
											)}
										</Select>
									</>
								)}
							</Field>
						</Grid>
					</Grid>
					<Grid container sx={styles.saveButtonContainer}>
						<Button
							type="submit"
							disabled={!dirty}
							variant="contained"
							fullWidth
							sx={styles.inputField}
						>
							Save config
						</Button>
					</Grid>
					<Grid container sx={styles.cancelButtonContainer}>
						<Button
							onClick={() => resetForm()}
							disabled={!dirty}
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

export default AdminConfigForm;
