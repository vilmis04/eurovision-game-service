import { TResponse, TSuccessResponse } from "../types/general.types";

export const isResponseSuccess = (
	response: TResponse
): response is TSuccessResponse => "data" in response;
