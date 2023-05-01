import { useNavigate, useParams } from "react-router-dom";
import { useSetGroupToJoinQuery } from "../../pages/@modules/group.api";
import { useEffect } from "react";
import { paths } from "../../paths";

const JoinGroup = () => {
	const { token1, token2, token3 } = useParams();
	const token = `${token1}.${token2}.${token3}`;
	const navigate = useNavigate();
	const { isSuccess } = useSetGroupToJoinQuery({ token });

	useEffect(() => {
		if (isSuccess) navigate(`/${paths.login}`);
	}, [isSuccess]);

	return null;
};

export default JoinGroup;
