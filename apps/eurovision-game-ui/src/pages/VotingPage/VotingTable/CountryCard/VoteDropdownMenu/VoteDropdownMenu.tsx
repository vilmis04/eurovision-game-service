import { VoteTypes } from "@eurovision-game-monorepo/core";
import {
	FormControl,
	InputLabel,
	MenuItem,
	Select,
	SelectChangeEvent,
} from "@mui/material";
import { useField, useFormikContext } from "formik";
import { ICountryCardProps } from "../CountryCard";
import { TVoteFormData } from "../../VotingTable.types";

// TODO: move to a constants file
const LIMITS: { [k in VoteTypes]: number } = {
	[VoteTypes.P1]: 1,
	[VoteTypes.P2_5]: 4,
	[VoteTypes.P6_10]: 5,
	[VoteTypes.P11_15]: 5,
	[VoteTypes.P16_20]: 5,
	[VoteTypes.P21_25]: 5,
};

interface IVoteDropdownMenuProps extends Pick<ICountryCardProps, "country"> {}

const VoteDropdownMenu: React.FC<IVoteDropdownMenuProps> = ({ country }) => {
	const { submitForm, values } = useFormikContext<TVoteFormData>();
	const [field, _meta, { setValue }] = useField(country);

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
		</FormControl>
	);
};

export default VoteDropdownMenu;
