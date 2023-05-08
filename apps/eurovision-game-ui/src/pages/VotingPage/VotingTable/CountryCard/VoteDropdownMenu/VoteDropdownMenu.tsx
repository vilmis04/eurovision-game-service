import { IAdminFormData, VoteTypes } from "@eurovision-game-monorepo/core";
import {
	FormControl,
	FormHelperText,
	InputLabel,
	MenuItem,
	Select,
	SelectChangeEvent,
} from "@mui/material";
import { useField, useFormikContext } from "formik";
import { ICountryCardProps } from "../CountryCard";
import { TVoteFormData } from "../../VotingTable.types";
import { useState } from "react";

// TODO: move to a constants file
const LIMITS: { [k in VoteTypes]: number } = {
	[VoteTypes.P1]: 1,
	[VoteTypes.P2_5]: 4,
	[VoteTypes.P6_10]: 5,
	[VoteTypes.P11_15]: 5,
	[VoteTypes.P16_20]: 5,
	[VoteTypes.P21_25]: 5,
};

interface IVoteDropdownMenuProps
	extends Pick<ICountryCardProps, "country">,
		Pick<IAdminFormData, "isVotingDisabled"> {}

const VoteDropdownMenu: React.FC<IVoteDropdownMenuProps> = ({
	country,
	isVotingDisabled,
}) => {
	const { submitForm, values } = useFormikContext<TVoteFormData>();
	const [field, _meta, { setValue }] = useField(country);
	const [shouldShowHelperText, setShouldShowHelperText] = useState(false);

	const handleChange = (event: SelectChangeEvent): void => {
		setValue(event.target.value);
		submitForm();
	};

	/** Return false when there are spots left */
	const checkLimits = (vote: VoteTypes) => {
		const selected = Object.values(values).filter(
			(selectedVote) => selectedVote === vote
		).length;
		return !(selected < LIMITS[vote]);
	};

	const color = field.value ? "success" : "error";

	const handleClick = () => {
		if (isVotingDisabled) toggleHelperText();
	};

	const toggleHelperText = () => setShouldShowHelperText((s) => !s);

	return (
		<FormControl fullWidth>
			<InputLabel id="voting-select-label">Vote</InputLabel>
			<Select
				{...field}
				onChange={(event) => handleChange(event)}
				fullWidth
				label="Vote"
				labelId="voting-select-label"
				error={!field.value}
				color={color}
				disabled={isVotingDisabled}
				onClick={handleClick}
			>
				<MenuItem value="">Vote</MenuItem>
				{Object.values(VoteTypes).map((item) => (
					<MenuItem
						disabled={checkLimits(item)}
						key={item}
						value={item}
					>
						{item}
					</MenuItem>
				))}
			</Select>
			{isVotingDisabled && shouldShowHelperText && (
				// TODO: move sx to a different file
				<FormHelperText error sx={{ paddingX: 1 }}>
					Voting is closed!
				</FormHelperText>
			)}
		</FormControl>
	);
};

export default VoteDropdownMenu;
