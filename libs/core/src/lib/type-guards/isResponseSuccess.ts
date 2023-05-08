import { TResponse, TSuccessResponse } from "../types/general.types";

export const isResponseSuccess = <T>(
	response: TResponse<T>
): response is TSuccessResponse<T> => "data" in response;
