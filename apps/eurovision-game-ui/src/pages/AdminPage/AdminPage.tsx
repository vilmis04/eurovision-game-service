import { IAdminFormData } from "@eurovision-game-monorepo/core";
import { FormikHelpers } from "formik";
import { useEffect } from "react";
import AdminConfigForm from "./AdminConfigForm/AdminConfigForm";
import {
	useGetAdminConfigQuery,
	useUpdateAdminConfigMutation,
} from "../@modules/admin.api";
import { CircularProgress } from "@mui/material";
import CountryConfigForm from "./CountryConfigForm/CountryConfigForm";

const AdminPage: React.FC = () => {
	// TODO: add effect to handle success / errors (snackbar or banner?)

	const { data: adminData } = useGetAdminConfigQuery();
	// igonring TS because the it shows no isSuccess when there is one
	// @ts-ignore
	const [updateAdminConfig, { isSuccess: isAdminConfigUpdateSuccess }] =
		useUpdateAdminConfigMutation();

	useEffect(() => {
		// TODO: show something for success (snackbar?)
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
			{adminData ? (
				<>
					<CountryConfigForm
						year={adminData.year}
						type={adminData.type}
					/>
					<AdminConfigForm
						initialValues={adminData}
						handleAdminConfigUpdate={handleAdminConfigUpdate}
					/>
				</>
			) : (
				<CircularProgress />
			)}
		</>
	);
};

export default AdminPage;
