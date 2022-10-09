import { Button } from "@mui/material"

interface IVoteButtonProps {
	country: string
	vote?: "1" | "2-5" | "5-10" | "10-15" | "15-20" | "20-25"
	handleClick: () => void
}

const VoteButton: React.FC<IVoteButtonProps> = ({
	country,
	vote,
	handleClick,
}) => {
	return (
		<Button
			variant="outlined"
			color={vote ? "success" : "error"}
			fullWidth
			onClick={handleClick}
		>
			{vote ?? "Vote"}
		</Button>
	)
}

export default VoteButton
