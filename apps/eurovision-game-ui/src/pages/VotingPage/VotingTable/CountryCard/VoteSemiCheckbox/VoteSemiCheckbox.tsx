import {
	Box,
	Checkbox,
	FormControlLabel,
	FormHelperText,
	Typography,
} from "@mui/material";
import { useField, useFormikContext } from "formik";
import { TVoteFormData } from "../../VotingTable.types";
import { useState } from "react";
import { IAdminFormData } from "@eurovision-game-monorepo/core";

// TODO: move to a constants file
const SPOTS_IN_FINAL = 10;

interface IVoteSemiCheckboxProps
	extends Pick<IAdminFormData, "isVotingDisabled"> {
	country: string;
}

const VoteSemiCheckbox: React.FC<IVoteSemiCheckboxProps> = ({
	country,
	isVotingDisabled,
}) => {
	const [field, _, { setValue }] = useField(country);
	const { submitForm, values } = useFormikContext<TVoteFormData>();
	const [shouldShowHelperText, setShouldShowHelperText] = useState(false);

	const handleChange = () => {
		setValue(!Boolean(field.value));
		submitForm();
	};

	const toggleHelperText = () => setShouldShowHelperText((s) => !s);

	/** Return false when there are spots left */
	const checkFinalSpotLimits = (isChecked: boolean) => {
		const selected = Object.values(values).filter((vote) => vote).length;
		return !isChecked && !(selected < SPOTS_IN_FINAL);
	};

	const isDisabled =
		isVotingDisabled || checkFinalSpotLimits(Boolean(field.value));
	const isHelperTextVisible = isDisabled && shouldShowHelperText;

	const handleClick = () => {
		if (isDisabled) toggleHelperText();
	};

	return (
		<Box onClick={handleClick}>
			<FormControlLabel
				// TODO: move styles to a different file
				sx={{ width: "100%" }}
				control={
					<Checkbox
						{...field}
						disabled={isDisabled}
						checked={Boolean(field.value)}
						onChange={handleChange}
						onClick={handleClick}
					/>
				}
				label={
					<Typography variant="body1">Through to final?</Typography>
				}
				labelPlacement="start"
			/>
			{isHelperTextVisible && (
				// TODO: move sx to a different file
				<FormHelperText error sx={{ paddingX: 1 }}>
					{isVotingDisabled
						? "Voting is closed!"
						: `All of the final spots are taken. Remove existing vote to select ${country}`}
				</FormHelperText>
			)}
		</Box>
	);
};

export default VoteSemiCheckbox;
