export const styles: ISx = {
	container: {
		alignItems: "center",
		flexDirection: "column",
		width: "100%",
		padding: 8,
	},
	title: ({ typography }) => ({
		fontSize: typography.h6,
	}),
	fieldGroup: ({ spacing }) => ({
		display: "flex",
		flexDirection: "column",
		justifyContent: "start",
		width: "100%",
		maxWidth: "235px",
		padding: spacing(2, 0),
	}),
};
