import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import {
	HttpMethods,
	IPostLoginRequest,
	IPostLoginResponse,
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
	}),
});

export const { usePostLoginDetailsMutation } = authApi;
