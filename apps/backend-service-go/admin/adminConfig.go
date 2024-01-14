package admin

type AdminConfig struct {
	Year           uint16   `json:"year"`
	GameType       GameType `json:"gameType"`
	IsVotingAcitve bool     `json:"isVotingActive"`
}

type GameType string

const (
	GameTypeSemi1 GameType = "semi1"
	GameTypeSemi2 GameType = "semi2"
	GameTypeFinal GameType = "final"
)

type CountryDetails struct {
	Country   string   `json:"country"`
	Year      uint16   `json:"year"`
	GameType  GameType `json:"gameType"`
	Score     uint16   `json:"score"`
	IsInFinal bool     `json:"isInFinal"`
}
