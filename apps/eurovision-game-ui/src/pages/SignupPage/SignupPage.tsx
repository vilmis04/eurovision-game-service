import { Alert, Button, Grid, Snackbar, Typography } from "@mui/material";
import { Form, Formik } from "formik";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import FormField from "../../components/FormField/FormField";
import { paths } from "../../paths";
import { useCreateUserMutation } from "../@modules/auth.api";
import { initialValues } from "./SignupPage.configs";
import { styles } from "./SignupPage.styles";
import { ISignupFormData } from "./SignupPage.types";
import { signupValidationSchema } from "./SignupPage.validation.schema";

const SignupPage: React.FC = () => {
	// TODO: isError handling
	// @ts-ignore
	const [createUser, { isSuccess, isError, error }] = useCreateUserMutation();
	const navigate = useNavigate();
	const [showSnackbar, setShowSnackbar] = useState(false);

	useEffect(() => {
		if (isSuccess) {
			navigate(paths.home);
		}
	}, [isSuccess]);

	useEffect(() => {
		if (isError) {
			const { status } = error as { status: number; data: unknown };
			status === 400 && setShowSnackbar(true);
		}
	}, [isError]);

	const toggleSnackbar = () => setShowSnackbar((s) => !s);

	const handleSubmit = (values: ISignupFormData) => {
		const { repeatPassword, ...body } = values;
		createUser(body);
	};

	return (
		<>
			<Formik
				initialValues={initialValues}
				validationSchema={signupValidationSchema}
				onSubmit={handleSubmit}
			>
				{({ errors, touched }) => (
					<Grid container sx={styles.container}>
						<Form>
							<Grid item sx={styles.fieldGroup}>
								<Typography sx={styles.title} component="h1">
									Create new account
								</Typography>
							</Grid>
							<Grid item sx={styles.fieldGroup}>
								<Grid item>
									<Typography>Username</Typography>
								</Grid>

								<FormField
									name="username"
									errors={errors}
									touched={touched}
								/>
							</Grid>
							<Grid item sx={styles.fieldGroup}>
								<Grid item>
									<Typography>Password</Typography>
								</Grid>
								<Grid item>
									<FormField
										name="password"
										errors={errors}
										touched={touched}
										type="password"
									/>
								</Grid>
							</Grid>
							<Grid item sx={styles.fieldGroup}>
								<Grid item>
									<Typography>Repeat password</Typography>
								</Grid>
								<Grid item>
									<FormField
										name="repeatPassword"
										errors={errors}
										touched={touched}
										type="password"
									/>
								</Grid>
							</Grid>
							<Grid item sx={styles.fieldGroup}>
								<Button
									variant="outlined"
									fullWidth
									type="submit"
								>
									Sign up
								</Button>
							</Grid>
						</Form>
					</Grid>
				)}
			</Formik>
			<Snackbar
				open={showSnackbar}
				onClose={toggleSnackbar}
				autoHideDuration={10000}
			>
				<Alert
					onClose={toggleSnackbar}
					severity="error"
					sx={{ width: "100%" }}
				>
					Username already exists!
				</Alert>
			</Snackbar>
		</>
	);
};

export default SignupPage;
