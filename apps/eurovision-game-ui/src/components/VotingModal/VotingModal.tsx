import { Box, Typography, Modal } from "@mui/material/"
import { styles } from "./VotingModal.styles"

interface IVotingModalProps {
	isOpen: boolean
	handleClose: () => void
}

const VotingModal: React.FC<IVotingModalProps> = ({ isOpen, handleClose }) => (
	<Modal
		open={isOpen}
		onClose={handleClose}
		aria-labelledby="modal-modal-title"
		aria-describedby="modal-modal-description"
	>
		<Box sx={styles.container}>
			<Typography id="modal-modal-title" variant="h6" component="h2">
				Text in a modal
			</Typography>
			<Typography id="modal-modal-description" sx={{ mt: 2 }}>
				Duis mollis, est non commodo luctus, nisi erat porttitor ligula.
			</Typography>
		</Box>
	</Modal>
)

export default VotingModal
