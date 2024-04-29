package score

import (
	"cmp"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"slices"

	"github.com/vilmis04/eurovision-game-service/internal/admin"
	"github.com/vilmis04/eurovision-game-service/internal/country"
)

type Service struct {
	storage        *Repo
	adminService   *admin.Service
	countryService *country.Service
}

func NewService() *Service {
	return &Service{
		storage:        NewRepo(),
		adminService:   admin.NewService(),
		countryService: country.NewService(),
	}
}

func (s *Service) InitializeScores(user string, year uint16, gameType admin.GameType) ([]ScoreResponse, error) {
	encodedCountries, err := s.countryService.GetCountryList(fmt.Sprint(year), fmt.Sprint(gameType), "")
	if err != nil {
		return nil, err
	}
	var countries []country.Country
	err = json.Unmarshal(*encodedCountries, &countries)
	if err != nil {
		return nil, err
	}

	scores, err := s.storage.InitializeScores(user, year, gameType, &countries)
	if err != nil {
		return nil, err
	}

	return scores, nil
}

func (s *Service) UpdateScore(user string, request *http.Request) error {
	encodedConfig, err := s.adminService.GetConfig()
	if err != nil {
		return err
	}
	var config admin.Admin
	err = json.Unmarshal(*encodedConfig, &config)
	if err != nil {
		return err
	}

	var body ScoreResponse
	err = json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		return fmt.Errorf("service: %v", err)
	}

	score, err := s.GetScore(user, body.Country, config.Year)
	if err != nil {
		return err
	}
	score.InFinal = body.InFinal
	score.Position = body.Position

	err = s.storage.UpdateScore(user, score)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetScore(user string, country string, year uint16) (*Score, error) {
	score, err := s.storage.GetScore(user, country, year)
	if err != nil {
		return nil, err
	}

	return score, nil
}

func (s *Service) GetAllScores(user string) (*[]byte, error) {
	encodedConfig, err := s.adminService.GetConfig()
	if err != nil {
		return nil, err
	}
	var config admin.Admin
	err = json.Unmarshal(*encodedConfig, &config)
	if err != nil {
		return nil, err
	}

	scores, err := s.storage.GetAllScores(user, config.GameType, config.Year)
	if err != nil {
		return nil, fmt.Errorf("failed to get scores for user %s: %v", user, err)
	}
	if len(scores) == 0 {
		scores, err = s.InitializeScores(user, config.Year, config.GameType)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize scores for user %s: %v", user, err)
		}
	}

	encodedScores, err := json.Marshal(scores)
	if err != nil {
		return nil, err
	}

	return &encodedScores, nil
}

func (s *Service) SortCountryList(countries []country.Country) (semiWinners []string, finalCountryList map[string]CountryResult) {
	var finalistCountryResult = []CountryResult{}
	finalCountryList = make(map[string]CountryResult)
	for _, country := range countries {
		if country.IsInFinal || country.GameType == admin.GameTypeFinal {
			finalistCountryResult = append(finalistCountryResult, CountryResult{
				Name:  country.Name,
				Score: country.Score,
			})
		}
		if country.IsInFinal {
			semiWinners = append(semiWinners, country.Name)
		}
	}
	slices.SortStableFunc(finalistCountryResult, func(a, b CountryResult) int {
		return cmp.Compare(b.Score, a.Score)
	})
	for index, country := range finalistCountryResult {
		finalCountryList[country.Name] = CountryResult{
			Name:     country.Name,
			Score:    country.Score,
			Position: int8(index + 1),
		}
	}
	return semiWinners, finalCountryList
}

func (s *Service) CalculateSemiScore(semiWinners []string, scores []Score) uint16 {
	var scoredPoints uint16 = 0
	for _, score := range scores {
		if !score.InFinal {
			continue
		}
		if slices.Contains(semiWinners, score.Country) {
			scoredPoints = scoredPoints + Points["semi"]
		}
	}
	return scoredPoints
}

func (s *Service) CalculateFinalScore(finalCountryList map[string]CountryResult, scores []Score) uint16 {
	var finalScore uint16 = 0
	for _, score := range scores {
		finalPosition := finalCountryList[score.Country].Position
		diff := finalPosition - score.Position
		absDiff := math.Abs(float64(diff))
		fmt.Printf("pos: %v, guessed: %v, country: %s, diff: %v, absDiff: %v\n", finalPosition, score.Position, score.Country, diff, absDiff)
		if absDiff > 5 {
			continue
		}
		if absDiff == 0 && finalPosition == 1 {
			finalScore = finalScore + Points["winner"]
			continue
		}
		finalScore = finalScore + Points[fmt.Sprint(absDiff)]
	}
	return finalScore
}

func (s *Service) CalculateTotalScore(countries []country.Country, scores []Score) uint16 {
	semiWinners, finalCountryList := s.SortCountryList(countries)
	semiScore := s.CalculateSemiScore(semiWinners, scores)
	finalScore := s.CalculateFinalScore(finalCountryList, scores)

	return finalScore + semiScore
}
