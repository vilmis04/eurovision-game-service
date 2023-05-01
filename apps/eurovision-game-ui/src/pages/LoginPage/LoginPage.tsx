import { IPostLoginRequest } from "@eurovision-game-monorepo/core";
import { Button, Grid, Typography } from "@mui/material";
import { Form, Formik } from "formik";
import { initialValues } from "./LoginPage.configs";
import { styles } from "./LoginPage.styles";
import {
	useGetAuthStatusQuery,
	usePostLoginDetailsMutation,
} from "../@modules/auth.api";
import FormField from "../../components/FormField/FormField";
import { loginValidationSchema } from "./LoginPage.validation.schema";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { paths } from "../../paths";
import { useJoinGroupMutation } from "../@modules/group.api";

const LoginPage: React.FC = () => {
	// TODO: isError handling
	// @ts-ignore
	const [postLoginDetails, { isSuccess }] = usePostLoginDetailsMutation();
	const [joinGroup] = useJoinGroupMutation();
	const { data: isAuthenticated } = useGetAuthStatusQuery();

	const navigate = useNavigate();

	useEffect(() => {
		if (isSuccess || isAuthenticated) {
			navigate(paths.home);
		}
	}, [isSuccess, isAuthenticated]);

	const handleSubmit = async (values: IPostLoginRequest) => {
		await postLoginDetails(values);
		await joinGroup();
	};

	const navigateToSignupPage = () => navigate(`/${paths.signup}`);

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
							<Button variant="outlined" fullWidth type="submit">
								Login
							</Button>
						</Grid>
						<Grid item sx={styles.fieldGroup}>
							{/* TODO: add sign up logic */}
							<Button
								variant="text"
								onClick={navigateToSignupPage}
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
