enum RootPaths {
	HOME = "/",
	LOGIN = "login",
	SIGNUP = "signup",
}

enum AuthPaths {
	AUTH_LOGIN = "auth/login",
	AUTH_SIGNUP = "auth/signup",
}
export const paths = {
	home: RootPaths.HOME,
	login: RootPaths.LOGIN,
	signup: RootPaths.SIGNUP,
	authLogin: AuthPaths.AUTH_LOGIN,
	authSignup: AuthPaths.AUTH_SIGNUP,
};
