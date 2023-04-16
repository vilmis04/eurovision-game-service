import {
	IAdminFormData,
	IToggleRegistrationRequest,
} from "@eurovision-game-monorepo/core";
import {
	BaseQueryFn,
	FetchArgs,
	FetchBaseQueryError,
	FetchBaseQueryMeta,
	MutationDefinition,
} from "@reduxjs/toolkit/dist/query";
import { MutationTrigger } from "@reduxjs/toolkit/dist/query/react/buildHooks";

export const mapFormDataToResponse = (values: IAdminFormData) => {
	return { countries: values.countries };
};

export const submitToggleRegistration = (
	checked: boolean,
	toggleRegistration: MutationTrigger<
		MutationDefinition<
			IToggleRegistrationRequest,
			BaseQueryFn<
				string | FetchArgs,
				unknown,
				FetchBaseQueryError,
				{},
				FetchBaseQueryMeta
			>,
			never,
			boolean,
			"adminApi"
		>
	>
) => {
	toggleRegistration({ isRegistrationEnabled: checked });
};
