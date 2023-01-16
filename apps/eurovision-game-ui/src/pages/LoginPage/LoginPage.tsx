import { Box, TextField } from "@mui/material";

// TODO: build proper UI and logic
const LoginPage: React.FC = () => (
	<Box>
		<Box>Username</Box>
		<TextField />
		<Box>Password</Box>
		<TextField type="password" />
	</Box>
);

export default LoginPage;
