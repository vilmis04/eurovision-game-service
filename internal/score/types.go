package score

import "github.com/vilmis04/eurovision-game-service/internal/admin"

type Score struct {
	Country  string         `json:"country"`
	Year     uint16         `json:"year"`
	User     string         `json:"user"`
	GameType admin.GameType `json:"gameType"`
	InFinal  bool           `json:"inFinal"`
	Points   uint16         `json:"points"`
}

type UpadateScoreRequestBody struct {
	InFinal *bool   `json:"inFinal"`
	Points  *uint16 `json:"points"`
}

type ScoreResponse struct {
	Country string `json:"country"`
	InFinal bool   `json:"inFinal"`
	Points  uint16 `json:"points"`
}
