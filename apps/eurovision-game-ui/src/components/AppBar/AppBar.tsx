import {
	AppBar as MUIAppBar,
	Box,
	Button,
	IconButton,
	Toolbar,
	Typography,
} from "@mui/material";
import MenuIcon from "@mui/icons-material/Menu";
import { HttpMethods } from "@eurovision-game-monorepo/core";

const handleLogin = async () => {
	const response = await fetch("http://localhost:4200/api/auth/login", {
		method: HttpMethods.POST,
		headers: [["Content-Type", "application/json"]],
		credentials: "include",
		body: JSON.stringify({ username: "test_user2", password: "changeme" }),
	});
	const result = await response.json();

	console.log(result);
	return response;
};

const AppBar: React.FC = () => (
	<Box sx={{ flexGrow: 1 }}>
		<MUIAppBar position="sticky">
			<Toolbar>
				<IconButton
					size="large"
					edge="start"
					color="inherit"
					aria-label="menu"
					sx={{ mr: 2 }}
				>
					<MenuIcon />
				</IconButton>
				<Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
					EuroVision Guessing Game
				</Typography>
				<Button color="inherit" onClick={handleLogin}>
					Login
				</Button>
			</Toolbar>
		</MUIAppBar>
	</Box>
);

export default AppBar;
