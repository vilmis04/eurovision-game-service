enum RootPaths {
	HOME = "/",
	LOGIN = "login",
	SIGNUP = "signup",
	ADMIN = "admin",
	SCORES = "scores",
}

enum AuthPaths {
	AUTH_LOGIN = "auth/login",
	AUTH_SIGNUP = "users/signup",
	AUTH_LOGOUT = "auth/logout",
}

export const paths = {
	home: RootPaths.HOME,
	login: RootPaths.LOGIN,
	signup: RootPaths.SIGNUP,
	authLogin: AuthPaths.AUTH_LOGIN,
	authSignup: AuthPaths.AUTH_SIGNUP,
	authLogOut: AuthPaths.AUTH_LOGOUT,
	admin: RootPaths.ADMIN,
	scores: RootPaths.SCORES,
};
