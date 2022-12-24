import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { IGetVotesResponse } from "@eurovision-game-monorepo/core";
import { HttpMethods } from "@eurovision-game-monorepo/core";

interface IGetVotesRequest {
	votes: IGetVotesResponse;
}

export const votesApi = createApi({
	reducerPath: "votesApi",
	baseQuery: fetchBaseQuery({ baseUrl: "http://localhost:4200/api/" }),
	endpoints: (builder) => ({
		getVotesByUsername: builder.query<IGetVotesResponse, void>({
			query: () => `votes`,
		}),
		editVotesByUsername: builder.mutation<
			IGetVotesResponse,
			IGetVotesRequest
		>({
			query: ({ votes }) => ({
				url: `votes`,
				method: HttpMethods.PATCH,
				body: votes,
			}),
		}),
	}),
});

export const { useGetVotesByUsernameQuery, useEditVotesByUsernameMutation } =
	votesApi;
