import { configureStore } from "@reduxjs/toolkit";
import { setupListeners } from "@reduxjs/toolkit/query";
import { adminApi } from "../pages/AdminPage/@modules/admin.api";
import { scoreApi } from "../pages/AdminPage/@modules/score.api";
import { authApi } from "../pages/LoginPage/@modules/auth.api";
import { votesApi } from "../pages/VotingPage/@modules/votes.api";
import { authMiddleware } from "./authMiddleware";

export const store = configureStore({
	reducer: {
		[votesApi.reducerPath]: votesApi.reducer,
		[authApi.reducerPath]: authApi.reducer,
		[scoreApi.reducerPath]: scoreApi.reducer,
		[adminApi.reducerPath]: adminApi.reducer,
	},
	middleware: (getDefaultMiddleware) =>
		getDefaultMiddleware().concat(
			votesApi.middleware,
			authApi.middleware,
			scoreApi.middleware,
			adminApi.middleware,
			authMiddleware
		),
});

setupListeners(store.dispatch);
