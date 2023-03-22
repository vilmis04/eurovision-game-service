import { HttpMethods, IAdminFormData } from "@eurovision-game-monorepo/core";
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/dist/query/react";
import { paths } from "apps/eurovision-game-ui/src/paths";
import { UpdateResult } from "mongodb";

export const adminApi = createApi({
	reducerPath: "adminApi",
	// TODO: move baseUrl to .env
	baseQuery: fetchBaseQuery({ baseUrl: "http://localhost:4200/api/" }),
	endpoints: (builder) => ({
		getAdminConfig: builder.query<IAdminFormData, void>({
			query: () => ({
				method: HttpMethods.GET,
				url: paths.admin,
				credentials: "include",
				headers: [["Content-Type", "application/json"]],
			}),
		}),
		updateAdminConfig: builder.mutation<UpdateResult, IAdminFormData>({
			// updateAdminConfig: builder.mutation<IAdminFormData, IAdminFormData>({
			query: (body) => ({
				method: HttpMethods.PATCH,
				url: paths.admin,
				credentials: "include",
				headers: [["Content-Type", "application/json"]],
				body,
			}),
		}),
	}),
});

export const { useGetAdminConfigQuery, useUpdateAdminConfigMutation } =
	adminApi;
