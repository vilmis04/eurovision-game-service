import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import {
	HttpMethods,
	IPostLoginRequest,
	IPostLoginResponse,
	IPostSignUpRequest,
	IPostSignUpResponse,
	RoleTypes,
} from "@eurovision-game-monorepo/core";
import { paths } from "apps/eurovision-game-ui/src/paths";

export const authApi = createApi({
	reducerPath: "authApi",
	// TODO: move baseUrl to .env
	baseQuery: fetchBaseQuery({ baseUrl: "http://localhost:4200/api/" }),
	endpoints: (builder) => ({
		postLoginDetails: builder.mutation<
			IPostLoginResponse,
			IPostLoginRequest
		>({
			query: (body) => ({
				method: HttpMethods.POST,
				url: paths.authLogin,
				credentials: "include",
				headers: [["Content-Type", "application/json"]],
				body,
			}),
		}),
		createUser: builder.mutation<IPostSignUpResponse, IPostSignUpRequest>({
			query: (body) => ({
				method: HttpMethods.POST,
				url: paths.authSignup,
				credentials: "include",
				headers: [["Content-Type", "application/json"]],
				body,
			}),
		}),
		logOut: builder.mutation<void, void>({
			query: () => ({
				method: HttpMethods.POST,
				url: paths.authLogOut,
				credentials: "include",
				headers: [["Content-Type", "application/json"]],
			}),
		}),
		getRoles: builder.query<RoleTypes[], void>({
			query: () => ({
				method: HttpMethods.GET,
				url: paths.authRoles,
				credentials: "include",
				headers: [["Content-Type", "application/json"]],
			}),
		}),
	}),
});

export const {
	usePostLoginDetailsMutation,
	useCreateUserMutation,
	useLogOutMutation,
	useGetRolesQuery,
} = authApi;
