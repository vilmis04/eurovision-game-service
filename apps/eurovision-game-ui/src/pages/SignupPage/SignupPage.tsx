import { Button, Grid, Typography } from "@mui/material";
import { Form, Formik } from "formik";
import { useEffect } from "react";
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
	const [createUser, { isSuccess }] = useCreateUserMutation();
	const navigate = useNavigate();

	useEffect(() => {
		if (isSuccess) {
			navigate(paths.home);
		}
	}, [isSuccess]);

	const handleSubmit = (values: ISignupFormData) => {
		const { repeatPassword, ...body } = values;
		createUser(body);
	};

	return (
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
							<Button variant="outlined" fullWidth type="submit">
								Sign up
							</Button>
						</Grid>
					</Form>
				</Grid>
			)}
		</Formik>
	);
};

export default SignupPage;
