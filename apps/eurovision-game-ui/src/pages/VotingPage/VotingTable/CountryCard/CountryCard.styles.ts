export const styles: ISx = {
	card: ({ spacing }) => ({
		padding: spacing(1),
		height: "100%",
	}),
	gap: ({ spacing }) => ({
		paddingLeft: spacing(1),
	}),
	song: { fontWeight: "bold" },
	button: ({ spacing }) => ({ padding: "0 25%", paddingTop: spacing(1) }),
};
