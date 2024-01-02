package admin

type AdminConfig struct {
	Year           uint16
	GameType       GameType
	IsVotingAcitve bool
}

type GameType string

const (
	GameTypeSemi1 GameType = "semi1"
	GameTypeSemi2 GameType = "semi2"
	GameTypeFinal GameType = "final"
)

type CountryDetails struct {
	Country   string
	Year      uint16
	GameType  GameType
	Score     uint16
	IsInFinal bool
}
