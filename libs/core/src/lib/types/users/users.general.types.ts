export interface IUser {
	username: string;
	password: string;
	roles: RoleTypes[];
	groups: string[];
}

export enum RoleTypes {
	ADMIN = "admin",
	USER = "user",
}
