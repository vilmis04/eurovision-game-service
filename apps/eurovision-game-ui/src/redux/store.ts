import { configureStore } from "@reduxjs/toolkit";
import { setupListeners } from "@reduxjs/toolkit/query";
import { votesApi } from "../pages/VotingPage/@module/votes.api";
import { authMiddleware } from "./authMiddleware";

export const store = configureStore({
	reducer: {
		[votesApi.reducerPath]: votesApi.reducer,
	},
	middleware: (getDefaultMiddleware) =>
		getDefaultMiddleware().concat(votesApi.middleware, authMiddleware),
});

setupListeners(store.dispatch);
