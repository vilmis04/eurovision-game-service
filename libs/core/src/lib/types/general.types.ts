import { FetchBaseQueryError } from "@reduxjs/toolkit/dist/query";
import { SerializedError } from "@reduxjs/toolkit";

export type TSuccessResponse<T> = {
	data: T;
};

export type TErrorResponse = {
	error: FetchBaseQueryError | SerializedError;
};

export type TResponse<T> = TSuccessResponse<T> | TErrorResponse;
