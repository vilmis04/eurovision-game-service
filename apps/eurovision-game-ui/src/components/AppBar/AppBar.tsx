import {
	AppBar,
	Box,
	Button,
	IconButton,
	Toolbar,
	Typography,
} from "@mui/material";
import MenuIcon from "@mui/icons-material/Menu";

const ButtonAppBar: React.FC = () => (
	<Box sx={{ flexGrow: 1 }}>
		<AppBar position="sticky">
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
				<Button color="inherit">Login</Button>
			</Toolbar>
		</AppBar>
	</Box>
);

export default ButtonAppBar;
