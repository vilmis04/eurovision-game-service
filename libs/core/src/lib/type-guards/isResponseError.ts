import { TErrorResponse, TResponse } from "../types/general.types";

export const isResponseError = <T>(
	response: TResponse<T>
): response is TErrorResponse => "error" in response;
