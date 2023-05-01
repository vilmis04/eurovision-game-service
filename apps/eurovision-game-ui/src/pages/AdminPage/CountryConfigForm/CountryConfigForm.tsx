import {
	GameTypes,
	ICountry,
	ICreateCountryFormData,
} from "@eurovision-game-monorepo/core";
import {
	Box,
	Button,
	Card,
	Checkbox,
	CircularProgress,
	Dialog,
	Grid,
	Typography,
} from "@mui/material";
import {
	useCreateCountryMutation,
	useDeleteCountryMutation,
	useGetCountryListQuery,
	useUpdateCountryMutation,
} from "../../@modules/country.api";
import { ChangeEvent, useEffect, useState } from "react";
import { Field, Form, Formik } from "formik";
import FormField from "apps/eurovision-game-ui/src/components/FormField/FormField";
import SaveIcon from "@mui/icons-material/Save";
import DeleteOutlineIcon from "@mui/icons-material/DeleteOutline";
import { usePromoteToFinalMutation } from "../../@modules/admin.api";

interface ICountryConfigFormProps {
	type: GameTypes;
	year: string;
}

enum CountryActions {
	UPDATE = "UPDATE",
	DELETE = "DELETE",
}

const CountryConfigForm: React.FC<ICountryConfigFormProps> = ({
	year,
	type,
}) => {
	const { data: countryList } = useGetCountryListQuery({ type, year });
	const [createCountry, { isLoading }] = useCreateCountryMutation();
	const [updateCountry] = useUpdateCountryMutation();
	const [deleteCountry] = useDeleteCountryMutation();
	// @ts-ignore
	const [promoteToFinal, { isSuccess: isPromoteToFinalSuccess }] =
		usePromoteToFinalMutation();

	const [showAddCountryPopup, setShowAddCountryPopup] = useState(false);
	const [showConfirmationPopup, setShowConfirmationPopup] = useState(false);
	const [countryAction, setCountryAction] = useState<CountryActions>(
		CountryActions.UPDATE
	);

	useEffect(() => {
		if (isPromoteToFinalSuccess) {
			toggleConfirmationPopup();
		}
	}, [isPromoteToFinalSuccess]);

	const isFinal = type === GameTypes.FINAL;

	const toggleAddCountryPopup = () => setShowAddCountryPopup((s) => !s);

	const toggleConfirmationPopup = () => setShowConfirmationPopup((s) => !s);

	const handleCreateCountry = async (values: ICreateCountryFormData) => {
		const requestBody = {
			...values,
			type,
			year,
		};

		await createCountry(requestBody);

		toggleAddCountryPopup();
	};

	const initialAddCountryValues: ICreateCountryFormData = {
		artist: "",
		name: "",
		song: "",
	};

	const handleUpdateCountry = async (values: ICountry) => {
		const { year, name, finalScore, semiFinalScore } = values;

		if (countryAction === CountryActions.UPDATE) {
			await updateCountry({ year, name, finalScore, semiFinalScore });
		}

		if (countryAction === CountryActions.DELETE) {
			await deleteCountry({ year, name });
		}
	};

	const handleUpdateClick = () => setCountryAction(CountryActions.UPDATE);
	const handleDeleteClick = () => setCountryAction(CountryActions.DELETE);

	const handlePromoteToFinal = () => {
		promoteToFinal();
	};
	const handleSubmitFinalScore = () => console.log("submit final score");

	const handleNextStage = isFinal
		? handleSubmitFinalScore
		: handlePromoteToFinal;

	// TODO: add a spinner until the results
	// TODO: move all sx
	return (
		<Box>
			{countryList ? (
				<Grid container spacing={2} sx={{ padding: 2 }}>
					{countryList.map((country) => (
						<Grid item container sm={12} key={country.name}>
							<Card sx={{ padding: 1, width: "100%" }}>
								<Box
									sx={{
										display: "flex",
										justifyContent: "space-between",
									}}
								>
									<Typography
										sx={{ width: "30%" }}
										variant="button"
									>
										{country.name}
									</Typography>
									<Formik
										initialValues={country}
										onSubmit={handleUpdateCountry}
									>
										{({
											values,
											setFieldValue,
											submitForm,
										}) => (
											<Form>
												{isFinal ? (
													<FormField
														// TODO: move sx
														sx={{
															maxWidth: "100px",
														}}
														name="finalScore"
														label="Final Result"
													/>
												) : (
													<Field
														name="semiFinalScore"
														as={Checkbox}
														checked={
															values.semiFinalScore
														}
														onChange={(
															_event: ChangeEvent,
															checked: boolean
														) => {
															setFieldValue(
																"semiFinalScore",
																checked
																	? true
																	: false
															);
															setCountryAction(
																CountryActions.UPDATE
															);
															submitForm();
														}}
													/>
												)}
												{isFinal && (
													<Button
														onClick={
															handleUpdateClick
														}
														type="submit"
													>
														<SaveIcon />
													</Button>
												)}
												<Button
													onClick={handleDeleteClick}
													type="submit"
												>
													<DeleteOutlineIcon />
												</Button>
											</Form>
										)}
									</Formik>
								</Box>
							</Card>
						</Grid>
					))}
				</Grid>
			) : (
				<CircularProgress />
			)}
			<Button onClick={toggleAddCountryPopup}>Add country</Button>
			<Button onClick={toggleConfirmationPopup}>
				{isFinal ? "Calculate score" : "Promote to final"}
			</Button>
			{/* TODO: Move to a separate component */}
			<Dialog
				open={showConfirmationPopup}
				onClose={toggleConfirmationPopup}
			>
				<Box
					sx={{
						padding: 2,
						display: "flex",
						flexDirection: "column",
						alignItems: "center",
					}}
				>
					<Typography
						variant="body1"
						sx={{ fontSize: "20px" }}
					>{`Confirm ${
						isFinal ? "final score" : "country promotion"
					}?`}</Typography>
					<Box>
						<Button onClick={handleNextStage}>Confirm</Button>
						<Button onClick={toggleConfirmationPopup}>
							Cancel
						</Button>
					</Box>
				</Box>
			</Dialog>
			{/* Move to separate component */}
			<Dialog open={showAddCountryPopup} onClose={toggleAddCountryPopup}>
				<Box>
					<Formik
						initialValues={initialAddCountryValues}
						onSubmit={handleCreateCountry}
					>
						{() => (
							<Box
								component={Form}
								sx={{
									display: "flex",
									flexDirection: "column",
									padding: 2,
								}}
							>
								<Box sx={{ margin: 1 }}>
									<Typography variant="h4">
										Add country
									</Typography>
								</Box>
								<FormField
									name="name"
									label="Country name"
									sx={{ margin: 1 }}
								/>
								<FormField
									name="song"
									label="Song title"
									sx={{ margin: 1 }}
								/>
								<FormField
									name="artist"
									label="Artist(s)"
									sx={{ margin: 1 }}
								/>
								<Button
									variant="contained"
									type="submit"
									color="primary"
									sx={{ margin: 1 }}
								>
									{isLoading ? (
										<CircularProgress />
									) : (
										"Create country"
									)}
								</Button>
								<Button
									variant="outlined"
									color="secondary"
									sx={{ margin: 1 }}
								>
									Cancel
								</Button>
							</Box>
						)}
					</Formik>
				</Box>
			</Dialog>
		</Box>
	);
};

export default CountryConfigForm;
