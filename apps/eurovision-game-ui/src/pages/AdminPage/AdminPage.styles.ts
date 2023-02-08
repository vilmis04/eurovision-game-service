export const styles: ISx = {
	container: ({ spacing }) => ({
		padding: spacing(2, 0, 6),
	}),
	field: {
		alignItems: "center",
		justifyContent: "center",
	},
	countryPoints: ({ spacing }) => ({
		padding: spacing(1, 0),
		alignItems: "center",
		justifyContent: "center",
	}),
	saveButtonContainer: {
		justifyContent: "center",
		padding: 2,
	},
	saveButton: {
		maxWidth: "250px",
	},
};
