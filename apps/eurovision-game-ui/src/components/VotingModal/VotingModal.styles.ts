export const styles: ISx = {
	container: ({ spacing, breakpoints }) => ({
		position: "absolute",
		top: "50%",
		left: "50%",
		transform: "translate(-50%, -50%)",
		width: "60%",
		backgroundColor: "background.paper",
		padding: spacing(4),
		[breakpoints.down("md")]: {
			width: "80%",
		},
		[breakpoints.down("sm")]: {
			width: "80%",
		},
	}),
}
