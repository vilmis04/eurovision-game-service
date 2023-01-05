import {
	isRejectedWithValue,
	Middleware,
	MiddlewareAPI,
} from "@reduxjs/toolkit";

export const authMiddleware: Middleware =
	(api: MiddlewareAPI) => (next) => (action) => {
		console.log("action intercepted");
		if (isRejectedWithValue(action) && action.payload.status === 304)
			window.location.replace("/login");

		return next(action);
	};
