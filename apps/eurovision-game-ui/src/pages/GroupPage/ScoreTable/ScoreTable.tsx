import { Box, Grid, IconButton, Typography } from "@mui/material";
import DeleteIcon from "@mui/icons-material/Delete";
import EditIcon from "@mui/icons-material/Edit";
import LinkIcon from "@mui/icons-material/Link";
import LogoutIcon from "@mui/icons-material/Logout";
import ManageAccountsIcon from "@mui/icons-material/ManageAccounts";
import LooksOneIcon from "@mui/icons-material/LooksOne";
import LooksTwoIcon from "@mui/icons-material/LooksTwo";
import Looks3Icon from "@mui/icons-material/Looks3";
import { getUsernameFromCookie } from "apps/eurovision-game-ui/src/utils/getUsernameFromCookie";
import { IGroupForAction } from "../GroupPage.types";
import {
	IFinalsList,
	IGroup,
	TUserVoteResponse,
} from "@eurovision-game-monorepo/core";
import { WithId } from "mongodb";
import CenterSpinner from "apps/eurovision-game-ui/src/components/CenterSpinner/CenterSpinner";
import { useState } from "react";

interface IScoreTableProps {
	group: WithId<IGroup>;
	handleInvitation: (group: IGroupForAction) => Promise<void>;
	handleLeave: (group: IGroupForAction) => void;
	handleDelete: (group: IGroupForAction) => void;
	handleEdit: (group: IGroupForAction) => void;
	semiWinnersList1: string[];
	semiWinnersList2: string[];
	finalsList: IFinalsList[];
	userVotes?: TUserVoteResponse[];
}

const POINTS_PER_GUESS = 20;

const ScoreTable: React.FC<IScoreTableProps> = ({
	group: { owner, name, _id, members },
	handleInvitation,
	handleLeave,
	handleDelete,
	handleEdit,
	semiWinnersList1,
	semiWinnersList2,
	finalsList,
	userVotes,
}) => {
	const [totalScores, setTotalScores] = useState<{ [k: string]: number }>({});
	const username = getUsernameFromCookie();
	const isOwner = owner === username;

	if (!userVotes) return <CenterSpinner />;

	const getIcon = (score: number) => {
		const sorted = [...Object.values(totalScores)].sort((a, b) => b - a);
		const top3 = [...new Set(sorted)].slice(0, 2);

		const position = top3.indexOf(score) + 1;

		if (position === 1)
			return (
				<LooksOneIcon
					sx={{
						fontSize: "16px",
						marginBottom: -0.4,
						paddingRight: 0.5,
					}}
				/>
			);
		if (position === 2)
			return (
				<LooksTwoIcon
					sx={{
						fontSize: "16px",
						marginBottom: -0.4,
						paddingRight: 0.5,
					}}
				/>
			);
		if (position === 3)
			return (
				<Looks3Icon
					sx={{
						fontSize: "16px",
						marginBottom: -0.4,
						paddingRight: 0.5,
					}}
				/>
			);
	};

	const getSemiScore = ({
		user,
		semi,
		semiWinnersList,
	}: {
		user: string;
		semi: "semi1" | "semi2";
		semiWinnersList: string[];
	}) => {
		const selectedCountries = userVotes.find(
			({ member }) => user === member
		)?.[semi];
		const correctChoices = (selectedCountries || []).reduce(
			(count, country) => {
				return semiWinnersList.includes(country) ? ++count : count;
			},
			0
		);

		return correctChoices * POINTS_PER_GUESS;
	};

	const getFinalsScore = (user: string) => {
		const userFinalVotes = userVotes.find(
			({ member }) => user === member
		)?.finals;

		const correctChoices = (userFinalVotes || []).reduce(
			(count, { name, position }) => {
				const finalPosition = finalsList.find(
					({ name: country }) => country === name
				)?.position;

				return finalPosition === position ? ++count : count;
			},
			0
		);
		const score = correctChoices * POINTS_PER_GUESS;
		return score;
	};

	return (
		<>
			<Box
				sx={{
					display: "flex",
					justifyContent: "space-between",
				}}
			>
				<Box>
					{isOwner && (
						<IconButton
							sx={{ margin: 1 }}
							onClick={() =>
								handleInvitation({
									_id: `${_id}`,
									name,
								})
							}
						>
							<LinkIcon />
						</IconButton>
					)}
				</Box>
				<Box>
					{isOwner && (
						<IconButton
							sx={{ margin: 1 }}
							onClick={() =>
								handleEdit({
									_id: `${_id}`,
									name,
								})
							}
							value={name}
						>
							<EditIcon />
						</IconButton>
					)}
					<IconButton
						sx={{ margin: 1 }}
						onClick={() =>
							isOwner
								? handleDelete({
										_id: `${_id}`,
										name,
								  })
								: handleLeave({
										_id: `${_id}`,
										name,
								  })
						}
					>
						{isOwner ? <DeleteIcon /> : <LogoutIcon />}
					</IconButton>
				</Box>
			</Box>
			<Grid container sx={{ width: "100%" }} rowSpacing={1}>
				<Grid item container>
					<Grid item xs={4}></Grid>
					<Grid item xs={2}>
						<Typography fontWeight="bold">S-1</Typography>
					</Grid>
					<Grid item xs={2}>
						<Typography fontWeight="bold">S-2</Typography>
					</Grid>
					<Grid item xs={2}>
						<Typography fontWeight="bold">Final</Typography>
					</Grid>
					<Grid item xs={2}>
						<Typography fontWeight="bold">Total</Typography>
					</Grid>
				</Grid>
				{members.map((member) => {
					const semiScore1 = getSemiScore({
						user: member,
						semi: "semi1",
						semiWinnersList: semiWinnersList1,
					});
					const semiScore2 = getSemiScore({
						user: member,
						semi: "semi2",
						semiWinnersList: semiWinnersList2,
					});
					const finalsScore = getFinalsScore(member);
					const total = semiScore1 + semiScore2 + finalsScore;

					if (totalScores[member] !== total) {
						setTotalScores((totals) => ({
							...totals,
							[member]: total,
						}));
					}

					return (
						<Grid item container key={member}>
							<Grid item xs={4}>
								{getIcon(total)}
								{member}
								{!isOwner && owner === member && (
									<ManageAccountsIcon
										sx={{
											marginBottom: -0.25,
											fontSize: "16px",
											paddingLeft: 0.75,
										}}
									/>
								)}
							</Grid>
							<Grid item xs={2}>
								{semiScore1}
							</Grid>
							<Grid item xs={2}>
								{semiScore2}
							</Grid>
							<Grid item xs={2}>
								{finalsScore}
							</Grid>
							<Grid
								item
								xs={2}
								// onChange={() =>
								// 	setTotalScores((totals) => ({
								// 		...totals,
								// 		[member]: total,
								// 	}))
								// }
							>
								{total}
							</Grid>
						</Grid>
					);
				})}
			</Grid>
		</>
	);
};

export default ScoreTable;
