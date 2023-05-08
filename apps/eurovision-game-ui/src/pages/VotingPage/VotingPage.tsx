import { Box } from "@mui/material";
import { useGetAdminConfigQuery } from "../@modules/admin.api";
import VotingTable from "./VotingTable/VotingTable";

const VotingPage: React.FC = () => {
	const { data: adminConfig } = useGetAdminConfigQuery(undefined, {
		// 10 min polling interval (calculation is ~15 min)
		pollingInterval: 600000,
	});
	if (!adminConfig) return <Box />;

	return (
		<VotingTable
			type={adminConfig.type}
			year={adminConfig.year}
			isVotingDisabled={adminConfig.isVotingDisabled}
		/>
	);
};

export default VotingPage;
