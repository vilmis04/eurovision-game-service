package country

import "github.com/vilmis04/eurovision-game-service/internal/admin"

type Country struct {
	Name       string         `json:"name"`
	Code       string         `json:"code"`
	Year       uint16         `json:"year"`
	GameType   admin.GameType `json:"gameType"`
	Score      uint16         `json:"score"`
	IsInFinal  bool           `json:"isInFinal"`
	Artist     string         `json:"artist"`
	Song       string         `json:"song"`
	OrderSemi  uint8          `json:"orderSemi"`
	OrderFinal uint8          `json:"orderFinal"`
}

type CreateCountryRequest struct {
	Name      *string         `json:"name"`
	Code      *string         `json:"code"`
	GameType  *admin.GameType `json:"gameType"`
	Artist    *string         `json:"artist"`
	Song      *string         `json:"song"`
	OrderSemi *uint8          `json:"orderSemi"`
}

type UpdateCountryRequest struct {
	Score      *uint16 `json:"score"`
	IsInFinal  *bool   `json:"isInFinal"`
	OrderSemi  *uint8  `json:"orderSemi"`
	OrderFinal *uint8  `json:"orderFinal"`
}

type CountrySummary struct {
	Name       string `json:"name"`
	Code       string `json:"code"`
	Artist     string `json:"artist"`
	Song       string `json:"song"`
	OrderSemi  uint8  `json:"orderSemi"`
	OrderFinal uint8  `json:"orderFinal"`
}
