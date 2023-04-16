import { IAdminFormData } from "@eurovision-game-monorepo/core";

export class Admin {
	isRegistrationEnabled: IAdminFormData["isRegistrationEnabled"];
	year: IAdminFormData["year"];
	type: IAdminFormData["type"];
}
