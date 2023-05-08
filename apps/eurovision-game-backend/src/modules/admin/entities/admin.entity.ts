import { IAdminFormData } from "@eurovision-game-monorepo/core";

export class Admin implements IAdminFormData {
	isRegistrationEnabled: IAdminFormData["isRegistrationEnabled"];
	year: IAdminFormData["year"];
	type: IAdminFormData["type"];
	isVotingDisabled: IAdminFormData["isVotingDisabled"];
	calculateFinalsScore: IAdminFormData["calculateFinalsScore"];
}
