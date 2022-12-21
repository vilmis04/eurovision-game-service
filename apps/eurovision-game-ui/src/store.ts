import { configureStore } from "@reduxjs/toolkit";
import { setupListeners } from "@reduxjs/toolkit/query";
import { votesApi } from "pages/VotingPage/@module/votes.api";

export const store = configureStore({
	reducer: {
		[votesApi.reducerPath]: votesApi.reducer,
	},
	middleware: (getDefaultMiddleware) =>
		getDefaultMiddleware().concat(votesApi.middleware),
});

setupListeners(store.dispatch);
