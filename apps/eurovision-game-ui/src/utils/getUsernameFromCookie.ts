export const getUsernameFromCookie = () =>
	document.cookie
		.split("; ")
		.find((cookie) => cookie.startsWith("username"))
		?.split("=")[1];
