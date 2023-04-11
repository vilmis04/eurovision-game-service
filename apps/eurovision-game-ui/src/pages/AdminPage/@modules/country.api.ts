import {
	GameTypes,
	HttpMethods,
	ICountry,
} from "@eurovision-game-monorepo/core";
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/dist/query/react";
import { paths } from "apps/eurovision-game-ui/src/paths";
import { DeleteResult, InsertOneResult, UpdateResult } from "mongodb";

export enum Tags {
	ADMIN = "admin",
	SCORE = "score",
	COUNTRY = "country",
}

// TODO: move types to different file
type TCountryListParams = { year: string; type: GameTypes };
type TCountryParams = { year: string; name: string };
type TCreateCountryRequestBody = Omit<
	ICountry,
	"semiFinalScore" | "finalScore"
>;
interface IUpdateCountryRequestBody extends Partial<ICountry> {
	name: ICountry["name"];
	year: ICountry["year"];
}

export const countryApi = createApi({
	reducerPath: "countryApi",
	// TODO: move baseUrl to .env
	baseQuery: fetchBaseQuery({ baseUrl: "http://localhost:4200/api/" }),
	tagTypes: Object.values(Tags),
	endpoints: (builder) => ({
		getCountryList: builder.query<ICountry[], TCountryListParams>({
			query: ({ year, type }) => ({
				method: HttpMethods.GET,
				url: `${paths.country}/${year}/${type}`,
				credentials: "include",
				headers: [["Content-Type", "application/json"]],
			}),
			providesTags: [Tags.COUNTRY],
		}),
		createCountry: builder.mutation<
			InsertOneResult<ICountry>,
			TCreateCountryRequestBody
		>({
			query: (body) => ({
				method: HttpMethods.POST,
				url: paths.country,
				credentials: "include",
				headers: [["Content-Type", "application/json"]],
				body,
			}),
			invalidatesTags: [Tags.COUNTRY],
		}),
		updateCountry: builder.mutation<
			UpdateResult,
			IUpdateCountryRequestBody
		>({
			query: (body) => ({
				method: HttpMethods.PATCH,
				url: paths.country,
				credentials: "include",
				headers: [["Content-Type", "application/json"]],
				body,
			}),
			invalidatesTags: [Tags.COUNTRY],
		}),
		deleteCountry: builder.mutation<DeleteResult, TCountryParams>({
			query: ({ year, name }) => ({
				method: HttpMethods.DELETE,
				url: `${paths.country}/${year}/${name}`,
				credentials: "include",
				headers: [["Content-Type", "application/json"]],
			}),
			invalidatesTags: [Tags.COUNTRY],
		}),
	}),
});

export const {
	useGetCountryListQuery,
	useCreateCountryMutation,
	useUpdateCountryMutation,
	useDeleteCountryMutation,
} = countryApi;
