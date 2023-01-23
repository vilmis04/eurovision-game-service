import { configureStore } from "@reduxjs/toolkit";
import { setupListeners } from "@reduxjs/toolkit/query";
import { authApi } from "../pages/LoginPage/modules/auth.api";
import { votesApi } from "../pages/VotingPage/modules/votes.api";
import { authMiddleware } from "./authMiddleware";

export const store = configureStore({
	reducer: {
		[votesApi.reducerPath]: votesApi.reducer,
		[authApi.reducerPath]: authApi.reducer,
	},
	middleware: (getDefaultMiddleware) =>
		getDefaultMiddleware().concat(
			votesApi.middleware,
			authApi.middleware,
			authMiddleware
		),
});

setupListeners(store.dispatch);
