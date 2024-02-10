package country

import "github.com/vilmis04/eurovision-game-service/internal/admin"

type Country struct {
	Name      string
	Year      uint16
	GameType  admin.GameType
	Score     uint16
	IsInFinal bool
	Artist    string
	Song      string
}

type CreateCountryRequest struct {
	Name     *string         `json:"name"`
	GameType *admin.GameType `json:"gameType"`
	Artist   *string         `json:"artist"`
	Song     *string         `json:"Song"`
}

type UpdateCountryRequest struct {
	Score     *uint16
	IsInFinal *bool
}
