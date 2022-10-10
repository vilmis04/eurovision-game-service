import { FormControl, InputLabel, MenuItem, Select } from "@mui/material";
import { useState } from "react";
import { VoteTypes } from "utils/votes";

interface IProps {
	onChange?: (newValue: string) => void;
}

const VoteDropdownMenu: React.FC<IProps> = ({ onChange }) => {
	const [value, setValue] = useState("");

	const handleChange = (newValue: string) => {
		if (onChange) onChange(newValue);
		setValue(newValue);
		console.log(value, newValue);
	};

	const color = value ? "success" : "error";

	return (
		<FormControl fullWidth>
			<InputLabel id="voting-select-label">Vote</InputLabel>
			<Select
				labelId="voting-select-label"
				id="voting-select"
				value={value}
				label="Vote"
				fullWidth
				onChange={(event) => handleChange(event.target.value)}
				error={!value}
				color={color}
			>
				<MenuItem value="">Vote</MenuItem>
				{Object.values(VoteTypes).map((item) => (
					<MenuItem key={item} value={item}>
						{item}
					</MenuItem>
				))}
			</Select>
		</FormControl>
	);
};

export default VoteDropdownMenu;
