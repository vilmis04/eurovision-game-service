import { HttpMethods, IAdminFormData } from "@eurovision-game-monorepo/core";
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/dist/query/react";
import { UpdateResult } from "mongodb";
import { paths } from "../../paths";

export enum Tags {
	ADMIN = "admin",
}

export const adminApi = createApi({
	reducerPath: "adminApi",
	// TODO: move baseUrl to .env
	baseQuery: fetchBaseQuery({ baseUrl: "http://localhost:4200/api/" }),
	tagTypes: Object.values(Tags),
	endpoints: (builder) => ({
		getAdminConfig: builder.query<IAdminFormData, void>({
			query: () => ({
				method: HttpMethods.GET,
				url: paths.admin,
				credentials: "include",
				headers: [["Content-Type", "application/json"]],
			}),
			providesTags: [Tags.ADMIN],
		}),
		updateAdminConfig: builder.mutation<UpdateResult, IAdminFormData>({
			query: (body) => ({
				method: HttpMethods.PATCH,
				url: paths.admin,
				credentials: "include",
				headers: [["Content-Type", "application/json"]],
				body,
			}),
			invalidatesTags: [Tags.ADMIN],
		}),
		promoteToFinal: builder.mutation<void, void>({
			query: () => ({
				method: HttpMethods.POST,
				url: paths.promote,
				credentials: "include",
				headers: [["Content-Type", "application/json"]],
			}),
		}),
		submitFinalScore: builder.mutation<UpdateResult, void>({
			query: () => ({
				method: HttpMethods.POST,
				url: paths.submitFinal,
				credentials: "include",
				headers: [["Content-Type", "application/json"]],
			}),
		}),
	}),
});

export const {
	useGetAdminConfigQuery,
	useUpdateAdminConfigMutation,
	usePromoteToFinalMutation,
	useSubmitFinalScoreMutation,
} = adminApi;
