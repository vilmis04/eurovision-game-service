import { useEffect, useState } from "react";
import {
	Accordion,
	AccordionDetails,
	AccordionSummary,
	Alert,
	Box,
	Button,
	Snackbar,
	Typography,
} from "@mui/material";
import ManageAccountsIcon from "@mui/icons-material/ManageAccounts";
import {
	ICreateGroupFormData,
	initialValues,
	isResponseSuccess,
} from "@eurovision-game-monorepo/core";
import {
	useCreateGroupMutation,
	useDeleteGroupMutation,
	useGenerateInvitationLinkMutation,
	useGetGroupUserVotesMutation,
	useGetGroupsQuery,
	useLeaveGroupMutation,
	useUpdateGroupMutation,
} from "../@modules/group.api";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import Popup from "../../components/Popup/Popup";
import { IGroupForAction, SubmitTypes } from "./GroupPage.types";
import { getPopupConfig } from "./GroupPage.config";
import ScoreTable from "./ScoreTable/ScoreTable";
import { getUsernameFromCookie } from "../../utils/getUsernameFromCookie";
import { useGetAdminConfigQuery } from "../@modules/admin.api";
import { useGetCountryList } from "./utils/useGetCountryList";
import CenterSpinner from "../../components/CenterSpinner/CenterSpinner";

const GroupPage: React.FC = () => {
	const { data: adminConfig } = useGetAdminConfigQuery();
	const { semiWinnersList1, semiWinnersList2, finalsList } =
		useGetCountryList(adminConfig?.type, adminConfig?.year);
	const {
		data: groups,
		isFetching: isFetchingGroups,
		isSuccess: isGetGroupsSuccess,
	} = useGetGroupsQuery({
		year: new Date().getFullYear().toString(),
	});
	const [showCreateGroupDialog, setShowCreateGroupDialog] = useState(false);
	const [submitType, setSubmitType] = useState(SubmitTypes.CREATE);
	const [link, setLink] = useState("");
	const [showLinkCopiedSnackbar, setShowLinkCopiedSnackbar] = useState(false);
	const [groupToEdit, setGroupToEdit] = useState<IGroupForAction | null>(
		null
	);
	const [
		createGroup,
		// @ts-ignore
		{ isSuccess: isCreateGroupSuccess, reset: resetCreate },
	] = useCreateGroupMutation();
	const [
		updateGroup,
		// @ts-ignore
		{ isSuccess: isUpdateGroupSuccess, reset: resetUpdate },
	] = useUpdateGroupMutation();
	const [
		deleteGroup,
		// @ts-ignore
		{ isSuccess: isDeleteGroupSuccess, reset: resetDelete },
	] = useDeleteGroupMutation();
	const [
		leaveGroup,
		// @ts-ignore
		{ isSuccess: isLeaveGroupSuccess, reset: resetLeave },
	] = useLeaveGroupMutation();
	const [generateInvitationLink] = useGenerateInvitationLinkMutation();
	const [
		getGroupUserVotes,
		// @ts-ignore
		{ isSuccess: isGetUserVotesSuccess, data: userVotes },
	] = useGetGroupUserVotesMutation();

	const toggleCreateGroupDialog = () => setShowCreateGroupDialog((s) => !s);

	const toggleLinkCopiedSnackbar = () => setShowLinkCopiedSnackbar((s) => !s);

	const handleClose = () => {
		setGroupToEdit(null);
		resetCreate();
		resetDelete();
		resetUpdate();
		resetLeave();
		toggleCreateGroupDialog();
	};

	const isSuccess =
		isCreateGroupSuccess ||
		isUpdateGroupSuccess ||
		isDeleteGroupSuccess ||
		isLeaveGroupSuccess;

	useEffect(() => {
		if (isSuccess) handleClose();
	}, [isSuccess]);

	useEffect(() => {
		if (isGetGroupsSuccess) {
			getGroupUserVotes(groups);
		}
	}, [isGetGroupsSuccess]);

	const handleSubmit = async ({ name, _id }: ICreateGroupFormData) => {
		if (submitType === SubmitTypes.CREATE) createGroup({ name });
		if (submitType === SubmitTypes.DELETE) deleteGroup({ id: `${_id}` });
		if (submitType === SubmitTypes.UPDATE)
			updateGroup({ id: `${_id}`, body: { name } });
		if (submitType === SubmitTypes.INVITE) {
			navigator.clipboard.writeText(link);
			toggleLinkCopiedSnackbar();
		}
		if (submitType === SubmitTypes.LEAVE) leaveGroup({ id: `${_id}` });
	};

	const handleCreate = () => {
		setSubmitType(SubmitTypes.CREATE);
		toggleCreateGroupDialog();
	};

	const handleEdit = (group: IGroupForAction) => {
		setSubmitType(SubmitTypes.UPDATE);
		setGroupToEdit(group);
		toggleCreateGroupDialog();
	};

	const handleDelete = (group: IGroupForAction) => {
		setSubmitType(SubmitTypes.DELETE);
		setGroupToEdit(group);
		toggleCreateGroupDialog();
	};

	const handleLeave = (group: IGroupForAction) => {
		setSubmitType(SubmitTypes.LEAVE);
		setGroupToEdit(group);
		toggleCreateGroupDialog();
	};

	const handleInvitation = async (group: IGroupForAction) => {
		setSubmitType(SubmitTypes.INVITE);
		setGroupToEdit(group);
		const response = await generateInvitationLink({
			id: `${group._id}`,
		});
		isResponseSuccess(response) && setLink(`${response.data}`);
		toggleCreateGroupDialog();
	};

	const popupConfig = getPopupConfig({ groupToEdit, link });
	const username = getUsernameFromCookie();

	// TODO: move all sx
	return (
		<>
			{isFetchingGroups ? (
				<CenterSpinner />
			) : (
				<Box>
					<Box>
						<Button
							variant="outlined"
							onClick={handleCreate}
							sx={{ margin: 1 }}
						>
							Create group
						</Button>
					</Box>
					<Box>
						{groups && groups.length > 0 ? (
							groups.map((group) => (
								<Box key={`${group._id}`} sx={{ margin: 2 }}>
									<Accordion sx={{ background: "#9dd2fa" }}>
										<AccordionSummary
											expandIcon={<ExpandMoreIcon />}
										>
											{group.name}
											{group.owner === username && (
												<ManageAccountsIcon
													sx={{
														fontSize: "16px",
														paddingLeft: 1,
													}}
												/>
											)}
										</AccordionSummary>
										<AccordionDetails
											sx={{
												paddingTop: 0,
												marginTop: -2,
											}}
										>
											<ScoreTable
												group={group}
												handleInvitation={
													handleInvitation
												}
												handleLeave={handleLeave}
												handleDelete={handleDelete}
												handleEdit={handleEdit}
												semiWinnersList1={
													semiWinnersList1
												}
												semiWinnersList2={
													semiWinnersList2
												}
												finalsList={finalsList}
												userVotes={userVotes}
											/>
										</AccordionDetails>
									</Accordion>
								</Box>
							))
						) : (
							<Typography variant="subtitle1">
								No groups to show
							</Typography>
						)}
					</Box>
					<Popup
						cancelLabel={popupConfig[submitType].cancelLabel}
						confirmLabel={popupConfig[submitType].confirmLabel}
						isOpen={showCreateGroupDialog}
						togglePopup={handleClose}
						initialValues={groupToEdit || initialValues}
						handleSubmit={handleSubmit}
						isLoading={popupConfig[submitType].isLoading}
					>
						{popupConfig[submitType].component}
					</Popup>
					<Snackbar
						open={showLinkCopiedSnackbar}
						autoHideDuration={5000}
						onClose={toggleLinkCopiedSnackbar}
					>
						<Alert
							onClose={toggleLinkCopiedSnackbar}
							severity="success"
							sx={{ width: "100%" }}
						>
							Link copied!
						</Alert>
					</Snackbar>
				</Box>
			)}
		</>
	);
};

export default GroupPage;
