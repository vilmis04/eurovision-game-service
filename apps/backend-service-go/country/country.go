package country

import "github.com/vilmis04/eurovision-game-monorepo/tree/main/apps/backend-service-go/admin"

type Country struct {
	Country   string         `json:"country"`
	Year      uint16         `json:"year"`
	GameType  admin.GameType `json:"gameType"`
	Score     uint16         `json:"score"`
	IsInFinal bool           `json:"isInFinal"`
}
