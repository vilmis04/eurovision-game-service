import { Box } from "@mui/material";
import { useGetAdminConfigQuery } from "../@modules/admin.api";
import VotingTable from "./VotingTable/VotingTable";

const VotingPage: React.FC = () => {
	const { data: adminConfig } = useGetAdminConfigQuery();
	if (!adminConfig) return <Box />;

	return <VotingTable type={adminConfig.type} year={adminConfig.year} />;
};

export default VotingPage;
