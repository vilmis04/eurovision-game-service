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
