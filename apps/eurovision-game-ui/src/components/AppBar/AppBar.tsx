import {
	AppBar as MUIAppBar,
	Box,
	Button,
	IconButton,
	Toolbar,
	Typography,
} from "@mui/material";
import MenuIcon from "@mui/icons-material/Menu";

// TODO: add logout functionality
const handleLogout = async () => {};

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
					EuroVision game
				</Typography>
				<Button color="inherit" onClick={handleLogout}>
					Log out
				</Button>
			</Toolbar>
		</MUIAppBar>
	</Box>
);

export default AppBar;
