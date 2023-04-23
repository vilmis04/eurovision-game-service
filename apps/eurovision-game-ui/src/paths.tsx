enum RootPaths {
	HOME = "/",
	LOGIN = "login",
	SIGNUP = "signup",
	ADMIN = "admin",
	SCORES = "scores",
	COUNTRY = "country",
	FORBIDDEN = "/forbidden",
}

enum AuthPaths {
	AUTH_LOGIN = "auth/login",
	AUTH_SIGNUP = "users/signup",
	AUTH_LOGOUT = "auth/logout",
	AUTH_ROLES = "auth/roles",
}

export const paths = {
	home: RootPaths.HOME,
	login: RootPaths.LOGIN,
	signup: RootPaths.SIGNUP,
	authLogin: AuthPaths.AUTH_LOGIN,
	authSignup: AuthPaths.AUTH_SIGNUP,
	authLogOut: AuthPaths.AUTH_LOGOUT,
	authRoles: AuthPaths.AUTH_ROLES,
	admin: RootPaths.ADMIN,
	scores: RootPaths.SCORES,
	country: RootPaths.COUNTRY,
	forbidden: RootPaths.FORBIDDEN,
};
