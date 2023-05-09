import { Box, FormControlLabel, TextField, Typography } from "@mui/material";
import { IGroupForAction, SubmitTypes } from "./GroupPage.types";
import FormField from "../../components/FormField/FormField";

interface IGetPopupConfigParams {
	groupToEdit: IGroupForAction | null;
	link?: string;
}

type TPopupConfig = {
	[k in SubmitTypes]: {
		confirmLabel: React.ReactNode;
		cancelLabel: React.ReactNode;
		component: React.ReactNode;
		isLoading?: boolean;
		hideButtons?: boolean;
	};
};

export const getPopupConfig = ({
	groupToEdit,
	link = "Preparing link",
}: IGetPopupConfigParams): TPopupConfig => ({
	[SubmitTypes.DELETE]: {
		confirmLabel: "Delete",
		cancelLabel: "Cancel",
		isLoading: !groupToEdit,
		component: (
			<Box>
				<Typography>
					{`Are you sure you want to delete ${groupToEdit?.name} group?`}
				</Typography>
			</Box>
		),
	},
	[SubmitTypes.UPDATE]: {
		confirmLabel: "Update",
		cancelLabel: "Cancel",
		component: (
			<FormControlLabel
				control={<FormField name="name" />}
				label={
					<Typography
						variant="body1"
						sx={{
							width: "100%",
						}}
					>
						Group name:
					</Typography>
				}
				labelPlacement="top"
			/>
		),
	},
	[SubmitTypes.CREATE]: {
		confirmLabel: "Create",
		cancelLabel: "Cancel",
		component: (
			<FormControlLabel
				control={<FormField name="name" />}
				label={
					<Typography
						variant="body1"
						sx={{
							width: "100%",
						}}
					>
						Group name:
					</Typography>
				}
				labelPlacement="top"
			/>
		),
	},
	[SubmitTypes.INVITE]: {
		confirmLabel: "Copy",
		cancelLabel: "Close",
		component: (
			<FormControlLabel
				control={<TextField value={link} />}
				label={
					<Typography
						variant="body1"
						sx={{
							width: "100%",
						}}
					>
						Invitation link:
					</Typography>
				}
				labelPlacement="top"
			/>
		),
	},
	[SubmitTypes.LEAVE]: {
		confirmLabel: "Leave",
		cancelLabel: "Cancel",
		isLoading: !groupToEdit,
		component: (
			<Box>
				<Typography>
					{`Are you sure you want to leave ${groupToEdit?.name} group?`}
				</Typography>
			</Box>
		),
	},
});
