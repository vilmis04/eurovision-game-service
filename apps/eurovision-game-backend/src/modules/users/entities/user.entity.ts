import { IUser } from "@eurovision-game-monorepo/core";

export class User {
	username: IUser["username"];
	password: IUser["password"];
	roles: IUser["roles"];
	groups: IUser["groups"];
}
