import {
	GameTypes,
	IAdminFormData,
	IScoreFormData,
} from "@eurovision-game-monorepo/core";
import { FormikHelpers } from "formik";
import { useEffect } from "react";
import AdminConfigForm from "./AdminConfigForm/AdminConfigForm";
import {
	useGetAdminConfigQuery,
	useUpdateAdminConfigMutation,
} from "./@modules/admin.api";
import ScoreForm from "./ScoreForm/ScoreForm";

// TODO: CONTINUE
/*
4. Fix score backend
5. Fix AdminConfigData code formatting and styles
6. AuthGuard
7. RoleGuard
*/

const scoreData: IScoreFormData = {
	type: GameTypes.FINAL,
	year: "2022",
	countries: {
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
	},
};

const AdminPage: React.FC = () => {
	// TODO: add effect to handle success / errors (snackbar?)

	const handleScoreFormSubmit = (values: IScoreFormData) => {
		// const requestBody = mapFormDataToResponse(values);
		// updateCountriesScores(requestBody);
	};

	const { data: adminData } = useGetAdminConfigQuery();
	const [updateAdminConfig, { isSuccess: isAdminConfigUpdateSuccess }] =
		useUpdateAdminConfigMutation();

	useEffect(() => {
		// TODO: show something for success (snackbar?)
		// TODO: we really need it ^^^
	}, [isAdminConfigUpdateSuccess]);

	const handleAdminConfigUpdate = async (
		values: IAdminFormData,
		{ resetForm }: FormikHelpers<IAdminFormData>
	) => {
		const { acknowledged } = await updateAdminConfig(values).unwrap();
		if (!acknowledged) {
			// TODO: inform about errors (snackbar?)

			return;
		}

		resetForm({ values });
	};

	return (
		<>
			{scoreData && adminData && (
				<>
					<ScoreForm
						initialValues={scoreData}
						handleScoreFormSubmit={handleScoreFormSubmit}
					/>
					<AdminConfigForm
						initialValues={adminData}
						handleAdminConfigUpdate={handleAdminConfigUpdate}
					/>
				</>
			)}
		</>
	);
};

export default AdminPage;
