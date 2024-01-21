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

type adminConfigRequestBody struct {
	Year           *uint16   `json:"year"`
	GameType       *GameType `json:"gameType"`
	IsVotingAcitve *bool     `json:"isVotingActive"`
}
