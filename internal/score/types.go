package score

import "github.com/vilmis04/eurovision-game-service/internal/admin"

type Score struct {
	Country  string         `json:"country"`
	Year     uint16         `json:"year"`
	User     string         `json:"user"`
	GameType admin.GameType `json:"gameType"`
	InFinal  bool           `json:"inFinal"`
	Position int            `json:"position"`
}

type ScoreResponse struct {
	Country  string `json:"country"`
	InFinal  bool   `json:"inFinal"`
	Position int    `json:"position"`
}

type CountryResult struct {
	Name     string
	Position int
	Score    uint16
}

var Points = map[string]uint16{
	"semi":   5,
	"winner": 25,
	"0":      12,
	"1":      8,
	"2":      5,
	"3":      3,
	"4":      2,
	"5":      1,
}
