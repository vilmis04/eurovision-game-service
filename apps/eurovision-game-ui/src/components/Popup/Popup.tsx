import { Box, Button, CircularProgress, Dialog } from "@mui/material";
import { Form, Formik } from "formik";
import { PropsWithChildren } from "react";

interface IPopupProps<T> extends PropsWithChildren {
	isOpen: boolean;
	togglePopup: () => void;
	confirmLabel: React.ReactNode;
	cancelLabel: React.ReactNode;
	initialValues: T;
	handleSubmit: (values: T) => void;
	isLoading?: boolean;
	hideButtons?: boolean;
}

const Popup = <T extends {}>({
	isOpen,
	togglePopup,
	cancelLabel,
	confirmLabel,
	children,
	initialValues,
	handleSubmit,
	isLoading,
	hideButtons = false,
}: IPopupProps<T>) => (
	<Dialog open={isOpen} onClose={togglePopup}>
		<>
			{isLoading ? (
				<Box sx={{ display: "flex" }}>
					<CircularProgress sx={{ margin: 5 }} />
				</Box>
			) : (
				<Formik initialValues={initialValues} onSubmit={handleSubmit}>
					{() => (
						<Box component={Form} sx={{ padding: 3 }}>
							{children}
							{!hideButtons && (
								<Box
									sx={({ spacing }) => ({
										padding: spacing(2, 2, 0),
										display: "flex",
										justifyContent: "space-between",
									})}
								>
									<Button variant="contained" type="submit">
										{confirmLabel}
									</Button>
									<Button
										variant="outlined"
										onClick={togglePopup}
									>
										{cancelLabel}
									</Button>
								</Box>
							)}
						</Box>
					)}
				</Formik>
			)}
		</>
	</Dialog>
);

export default Popup;
