import {
	isRejectedWithValue,
	Middleware,
	MiddlewareAPI,
} from "@reduxjs/toolkit";
import { paths } from "../paths";

export const authMiddleware: Middleware =
	(api: MiddlewareAPI) => (next) => (action) => {
		if (isRejectedWithValue(action) && action.payload.status === 401) {
			window.location.replace(`/${paths.login}`);
		}

		return next(action);
	};
