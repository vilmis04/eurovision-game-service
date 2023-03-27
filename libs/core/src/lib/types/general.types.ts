import { UpdateResult } from "mongodb";
import { FetchBaseQueryError } from "@reduxjs/toolkit/dist/query";
import { SerializedError } from "@reduxjs/toolkit";

export type TSuccessResponse = {
	data: UpdateResult;
};

export type TErrorResponse = {
	error: FetchBaseQueryError | SerializedError;
};

export type TResponse = TSuccessResponse | TErrorResponse;
