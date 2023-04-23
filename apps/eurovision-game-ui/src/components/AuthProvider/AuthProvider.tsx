import { useEffect, FC } from "react";
import { useNavigate } from "react-router-dom";
import { useGetRolesQuery } from "../../pages/LoginPage/@modules/auth.api";
import PageLayout from "../PageLayout/PageLayout";
import { RoleTypes } from "@eurovision-game-monorepo/core";

const AuthProvider: FC = () => {
	const { data: roles = [], isSuccess } = useGetRolesQuery();

	const navigate = useNavigate();

	useEffect(() => {
		if (roles.length > 0 && !roles.includes(RoleTypes.ADMIN)) {
			navigate("/forbidden");
		}
	}, [isSuccess]);

	const isAdmin = roles.includes(RoleTypes.ADMIN);

	return isAdmin ? <PageLayout /> : null;
};

export default AuthProvider;
