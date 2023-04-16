export const styles: ISx = {
	field: {
		alignItems: "center",
		justifyContent: "center",
	},
	divider: {
		borderTop: "1px solid black",
	},
	container: ({ spacing }) => ({
		padding: spacing(2, 0, 6),
	}),
	saveButtonContainer: {
		justifyContent: "center",
		padding: 2,
	},
	cancelButtonContainer: {
		justifyContent: "center",
		padding: 2,
		paddingTop: 0,
	},
	inputField: {
		maxWidth: "250px",
		display: "block",
	},
	fieldTitle: {
		margin: "0 auto",
		width: "250px",
	},
};
