export const styles: ISx = {
	tableContainer: ({ spacing }) => ({
		display: "flex",
		flexDirection: "columns",
		alignItems: "center",
		padding: spacing(3, 3, 0),
		maxWidth: 1200,
		margin: "auto",
	}),
	card: ({ spacing }) => ({
		padding: spacing(1),
	}),
}
