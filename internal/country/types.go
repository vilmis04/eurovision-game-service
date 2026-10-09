package country

import (
	"strconv"
	"strings"

	"github.com/vilmis04/eurovision-game-service/internal/admin"
	"github.com/vilmis04/eurovision-game-service/internal/utils"
)

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

const maxFieldLength = 255

func parseYear(year string) (int, error) {
	parsed, err := strconv.Atoi(year)
	if err != nil || parsed < 1900 || parsed > 3000 {
		return 0, utils.BadRequest("invalid year")
	}

	return parsed, nil
}

func parseName(name string) (string, error) {
	if name == "" || len(name) > maxFieldLength {
		return "", utils.BadRequest("invalid country name")
	}

	return name, nil
}

func validText(value *string) bool {
	return value != nil && strings.TrimSpace(*value) != "" && len(*value) <= maxFieldLength
}

func (r *CreateCountryRequest) validate() error {
	if !validText(r.Name) || !validText(r.Code) || !validText(r.Artist) || !validText(r.Song) || r.GameType == nil || r.OrderSemi == nil {
		return utils.BadRequest("missing or invalid country fields")
	}
	switch admin.GameType(*r.GameType) {
	case admin.GameTypeSemi1, admin.GameTypeSemi2, admin.GameTypeFinal:
	default:
		return utils.BadRequest("invalid game type")
	}

	return nil
}
