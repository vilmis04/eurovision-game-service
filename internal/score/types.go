package score

type Score struct {
	Country string `json:"country"`
	Year    uint16 `json:"year"`
	User    string `json:"user"`
	InFinal bool   `json:"inFinal"`
	Points  uint16 `json:"points"`
}

type UpadateScoreRequestBody struct {
	InFinal *bool   `json:"inFinal"`
	Points  *uint16 `json:"points"`
}
