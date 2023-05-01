import { configureStore } from "@reduxjs/toolkit";
import { setupListeners } from "@reduxjs/toolkit/query";
import { adminApi } from "../pages/@modules/admin.api";
import { scoreApi } from "../pages/@modules/score.api";
import { authApi } from "../pages/@modules/auth.api";
import { votesApi } from "../pages/@modules/votes.api";
import { authMiddleware } from "./authMiddleware";
import { countryApi } from "../pages/@modules/country.api";
import { groupApi } from "../pages/@modules/group.api";

export const store = configureStore({
	reducer: {
		[votesApi.reducerPath]: votesApi.reducer,
		[authApi.reducerPath]: authApi.reducer,
		[adminApi.reducerPath]: adminApi.reducer,
		[scoreApi.reducerPath]: scoreApi.reducer,
		[countryApi.reducerPath]: countryApi.reducer,
		[groupApi.reducerPath]: groupApi.reducer,
	},
	middleware: (getDefaultMiddleware) =>
		getDefaultMiddleware().concat(
			votesApi.middleware,
			authApi.middleware,
			scoreApi.middleware,
			adminApi.middleware,
			countryApi.middleware,
			groupApi.middleware,
			authMiddleware
		),
});

setupListeners(store.dispatch);
