import { IPostLoginRequest } from "@eurovision-game-monorepo/core";
import { Button, Grid, Typography } from "@mui/material";
import { Form, Formik } from "formik";
import { initialValues } from "./LoginPage.configs";
import { styles } from "./LoginPage.styles";
import { usePostLoginDetailsMutation } from "./modules/auth.api";
import FormTextField from "../../components/FormTextField/FormTextField";
import { loginValidationSchema } from "./LoginPage.validation.schema";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { paths } from "../../paths";

const LoginPage: React.FC = () => {
	const [postLoginDetails, { isSuccess }] = usePostLoginDetailsMutation();
	const navigate = useNavigate();

	useEffect(() => {
		if (isSuccess) {
			navigate(paths.home);
		}
	}, [isSuccess]);

	const handleSubmit = (values: IPostLoginRequest) => {
		postLoginDetails(values);
	};

	return (
		<Formik
			initialValues={initialValues}
			validationSchema={loginValidationSchema}
			onSubmit={handleSubmit}
		>
			{({ errors, touched }) => (
				<Grid container sx={styles.container}>
					<Form>
						<Grid item sx={styles.fieldGroup}>
							<Typography sx={styles.title} component="h1">
								Login to your account
							</Typography>
						</Grid>
						<Grid item sx={styles.fieldGroup}>
							<Grid item>
								<Typography>Username</Typography>
							</Grid>

							<FormTextField
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
								<FormTextField
									name="password"
									errors={errors}
									touched={touched}
									type="password"
								/>
							</Grid>
						</Grid>
						<Grid item sx={styles.fieldGroup}>
							<Button variant="outlined" fullWidth type="submit">
								Login
							</Button>
						</Grid>
						<Grid item sx={styles.fieldGroup}>
							{/* TODO: add sign up logic */}
							<Button
								variant="text"
								onClick={() => {}}
								fullWidth
								type="button"
							>
								Sign up
							</Button>
						</Grid>
					</Form>
				</Grid>
			)}
		</Formik>
	);
};

export default LoginPage;
