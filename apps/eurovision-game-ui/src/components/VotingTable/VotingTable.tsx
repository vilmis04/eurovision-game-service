import { Grid } from "@mui/material"
import { styles } from "./VotingTable.styles"
import CountryCard from "components/CountryCard/CountryCard"
import { useEffect, useState } from "react"
import { Box } from "@mui/system"
import VotingModal from "components/VotingModal/VotingModal"

interface ICountry {
	country: string
	artist: string
	song: string
}

const VotingTable = () => {
	const [countries, setCountries] = useState([])
	const [isLoading, setIsLoading] = useState(countries.length === 0)
	const [isModalOpen, setIsModalOpen] = useState(false)

	useEffect(() => {
		fetch("http://localhost:3004/EV22")
			.then((res) => res.json())
			.then((res) => setCountries(res))
			.catch((err) => console.log(err))
			.finally(() => setIsLoading(false))
	}, [])

	const toggleModal = () => setIsModalOpen(!isModalOpen)

	return (
		<>
			{isModalOpen && (
				<VotingModal isOpen={isModalOpen} handleClose={toggleModal} />
			)}
			{isLoading ? (
				<Box>Loading</Box>
			) : (
				<Grid container sx={styles.tableContainer}>
					{countries.map((entry: ICountry) => (
						<Grid
							item
							md={4}
							sm={6}
							xs={12}
							key={entry.country}
							sx={styles.card}
						>
							<CountryCard
								country={entry.country}
								artist={entry.artist}
								song={entry.song}
								toggleModal={toggleModal}
							/>
						</Grid>
					))}
				</Grid>
			)}
		</>
	)
}

export default VotingTable
