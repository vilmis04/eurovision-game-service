import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { IGetVotesResponse } from "./votes.types";

export const votesApi = createApi({
	reducerPath: "votesApi",
	baseQuery: fetchBaseQuery({ baseUrl: "http://localhost:4200/" }),
	endpoints: (builder) => ({
		getVotesByUsername: builder.query<IGetVotesResponse, void>({
			query: () => `votes`,
		}),
		editVotesByUsername: builder.query<IGetVotesResponse, void>({
			query: () => `votes`,
		}),
	}),
});

export const { useGetVotesByUsernameQuery, useEditVotesByUsernameQuery } =
	votesApi;
