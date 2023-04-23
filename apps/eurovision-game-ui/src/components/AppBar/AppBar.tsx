import {
	AppBar as MUIAppBar,
	Box,
	Button,
	IconButton,
	Toolbar,
	Typography,
} from "@mui/material";
import MenuIcon from "@mui/icons-material/Menu";
import { useLogOutMutation } from "../../pages/LoginPage/@modules/auth.api";
import { useNavigate } from "react-router-dom";
import { useEffect } from "react";
import { paths } from "../../paths";

const AppBar: React.FC = () => {
	const [logout, { isSuccess }] = useLogOutMutation();
	const navigate = useNavigate();

	useEffect(() => {
		if (isSuccess) {
			navigate(`/${paths.login}`);
		}
	}, [isSuccess]);

	const handleLogout = () => logout();

	return (
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
					<Typography
						variant="h6"
						component="div"
						sx={{ flexGrow: 1 }}
					>
						EuroVision game
					</Typography>
					<Button color="inherit" onClick={handleLogout}>
						Log out
					</Button>
				</Toolbar>
			</MUIAppBar>
		</Box>
	);
};

export default AppBar;
