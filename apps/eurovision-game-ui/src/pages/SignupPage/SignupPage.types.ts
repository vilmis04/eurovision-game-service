import { IPostSignUpRequest } from "@eurovision-game-monorepo/core";

export interface ISignupFormData extends IPostSignUpRequest {
	repeatPassword: string;
}
