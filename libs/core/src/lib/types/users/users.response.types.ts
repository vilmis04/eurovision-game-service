import { IUser } from "./users.general.types";

export interface IGetUserResponse extends IUser {}

export interface IUserDataForToken extends Omit<IGetUserResponse, "password"> {}
