export interface IGetUserResponse {
	username: string;
	password: string;
}

export interface IUserDataForToken extends Omit<IGetUserResponse, "password"> {}
