import {
	HttpMethods,
	IGetScoreRequest,
	IScoreFormData,
} from "@eurovision-game-monorepo/core";
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/dist/query/react";
import { paths } from "apps/eurovision-game-ui/src/paths";
import { UpdateResult } from "mongodb";
import { Tags } from "./admin.api";

export const scoreApi = createApi({
	reducerPath: "scoreApi",
	// TODO: move baseUrl to .env
	baseQuery: fetchBaseQuery({ baseUrl: "http://localhost:4200/api/" }),
	tagTypes: Object.values(Tags),
	endpoints: (builder) => ({
		createScore: builder.mutation<string, IScoreFormData>({
			query: (body) => ({
				method: HttpMethods.POST,
				url: paths.scores,
				credentials: "include",
				headers: [["Content-Type", "application/json"]],
				body,
			}),
			invalidatesTags: [Tags.SCORE],
		}),
		getScore: builder.query<IScoreFormData, IGetScoreRequest>({
			query: ({ year, type }) => ({
				method: HttpMethods.GET,
				url: `${paths.scores}/${year}/${type}`,
				credentials: "include",
				headers: [["Content-Type", "application/json"]],
			}),
			providesTags: [Tags.SCORE],
			forceRefetch: () => true,
		}),
		updateScore: builder.mutation<UpdateResult, IScoreFormData>({
			query: (body) => ({
				method: HttpMethods.PATCH,
				url: paths.scores,
				credentials: "include",
				headers: [["Content-Type", "application/json"]],
				body,
			}),
			invalidatesTags: [Tags.SCORE],
		}),
	}),
});

export const {
	useCreateScoreMutation,
	useGetScoreQuery,
	useUpdateScoreMutation,
} = scoreApi;
