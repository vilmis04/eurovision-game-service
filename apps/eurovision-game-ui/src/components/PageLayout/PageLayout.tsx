import { Outlet } from "react-router-dom";

import AppBar from "../AppBar/AppBar";

const PageLayout: React.FC = () => (
	<>
		<AppBar />
		<Outlet />
	</>
);
export default PageLayout;
