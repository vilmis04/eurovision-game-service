import {
	AppBar as MUIAppBar,
	Box,
	Button,
	IconButton,
	Toolbar,
	Typography,
	Drawer,
} from "@mui/material";
import MenuIcon from "@mui/icons-material/Menu";
import {
	useGetRolesQuery,
	useLogOutMutation,
} from "../../pages/@modules/auth.api";
import { useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";
import { paths } from "../../paths";
import CloseIcon from "@mui/icons-material/Close";
import { RoleTypes } from "@eurovision-game-monorepo/core";
import { getUsernameFromCookie } from "../../utils/getUsernameFromCookie";

const AppBar: React.FC = () => {
	// @ts-ignore
	const [logout, { isSuccess }] = useLogOutMutation();
	const navigate = useNavigate();
	const { data: roles = [] } = useGetRolesQuery();
	const [isDrawerOpen, setIsDrawerOpen] = useState(false);

	const toggleDrawer = () => setIsDrawerOpen((s) => !s);

	const isAdmin = roles.includes(RoleTypes.ADMIN);

	const username = getUsernameFromCookie();

	useEffect(() => {
		if (isSuccess) {
			navigate(`/${paths.login}`);
		}
	}, [isSuccess]);

	useEffect(() => {
		setIsDrawerOpen(false);
	}, []);

	const handleLogout = () => {
		logout();
		setIsDrawerOpen(false);
	};

	const navigateTo = (path: string) => {
		navigate(path);
		setIsDrawerOpen(false);
	};

	return (
		<>
			<Box sx={{ flexGrow: 1 }}>
				<MUIAppBar position="sticky">
					<Toolbar>
						<IconButton
							size="large"
							edge="start"
							color="inherit"
							aria-label="menu"
							sx={{ mr: 2 }}
							onClick={toggleDrawer}
						>
							<MenuIcon />
						</IconButton>
						<Typography
							variant="h6"
							component="div"
							sx={{ flexGrow: 1 }}
						>
							Vote For The Winners
						</Typography>
					</Toolbar>
				</MUIAppBar>
			</Box>
			<Drawer open={isDrawerOpen} onClose={toggleDrawer}>
				<Box
					sx={{
						background: "#daedf5",
						padding: 2,
						width: "250px",
						height: "100%",
						display: "flex",
						flexDirection: "column",
						justifyContent: "space-between",
					}}
				>
					<Box
						sx={{
							display: "flex",
							flexDirection: "column",
							alignItems: "start",
						}}
					>
						<Box
							sx={{
								display: "flex",
								width: "100%",
								justifyContent: "space-between",
								paddingBottom: 3,
							}}
						>
							<Box sx={{ display: "flex", alignItems: "center" }}>
								<Typography>{username}</Typography>
							</Box>
							<IconButton
								aria-label="Close menu"
								onClick={toggleDrawer}
							>
								<CloseIcon />
							</IconButton>
						</Box>
						<Button onClick={() => navigateTo(paths.home)}>
							Vote
						</Button>
						<Button onClick={() => navigateTo(`/${paths.groups}`)}>
							Groups
						</Button>
						{isAdmin && (
							<Button
								onClick={() => navigateTo(`/${paths.admin}`)}
							>
								Admin
							</Button>
						)}
					</Box>
					<Button onClick={handleLogout}>Log out</Button>
				</Box>
			</Drawer>
		</>
	);
};

export default AppBar;
