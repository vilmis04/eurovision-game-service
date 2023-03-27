import { TErrorResponse, TResponse } from "../types/general.types";

export const isResponseError = (
	response: TResponse
): response is TErrorResponse => "error" in response;
