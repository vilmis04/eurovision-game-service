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

	if !config.IsVotingAcitve {
		return fmt.Errorf("voting is not active")
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

func (s *Service) GetAllScores(user string, allGameTypes bool) (*[]byte, error) {
	encodedConfig, err := s.adminService.GetConfig()
	if err != nil {
		return nil, err
	}
	var config admin.Admin
	err = json.Unmarshal(*encodedConfig, &config)
	if err != nil {
		return nil, err
	}

	var gameType admin.GameType = ""
	if !allGameTypes {
		gameType = config.GameType
	}
	scores, err := s.storage.GetAllScores(user, gameType, config.Year)
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

func (s *Service) GetUserScore(user string) (uint16, error) {
	adminConfigResponse, err := s.adminService.GetConfig()
	if err != nil {
		return 0, err
	}
	var adminConfig admin.Admin
	err = json.Unmarshal(*adminConfigResponse, &adminConfig)
	if err != nil {
		return 0, err
	}

	countryResponse, err := s.countryService.GetCountryList(fmt.Sprint(adminConfig.Year), "", "")
	if err != nil {
		return 0, err
	}
	var countries []country.Country
	err = json.Unmarshal(*countryResponse, &countries)
	if err != nil {
		return 0, err
	}

	scoresResponse, err := s.GetAllScores(user, true)
	if err != nil {
		return 0, err
	}

	var scores []Score
	err = json.Unmarshal(*scoresResponse, &scores)
	if err != nil {
		return 0, err
	}

	skipFinalCalculation := false
	if adminConfig.GameType != admin.GameTypeFinal || adminConfig.IsVotingAcitve {
		skipFinalCalculation = true
	}

	return s.CalculateTotalScore(countries, scores, skipFinalCalculation), nil
}

func (s *Service) sortCountryList(countries []country.Country) (semiWinners []string, finalCountryList map[string]CountryResult) {
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
			Position: index + 1,
		}
	}
	return semiWinners, finalCountryList
}

func (s *Service) calculateSemiScore(semiWinners []string, scores []Score) uint16 {
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

func (s *Service) calculateFinalScore(finalCountryList map[string]CountryResult, scores []Score, skipFinalCalculation bool) uint16 {
	var finalScore uint16 = 0
	if skipFinalCalculation {
		return finalScore
	}
	for _, score := range scores {
		finalPosition := finalCountryList[score.Country].Position
		if score.Position == 0 || finalPosition == 0 {
			continue
		}
		diff := finalPosition - score.Position
		absDiff := math.Abs(float64(diff))
		if absDiff > 5 {
			continue
		}
		if absDiff == 0 && finalPosition == 1 {
			finalScore = finalScore + Points["winner"]
			continue
		}
		fmt.Printf("diff: %v, absDiff: %v, finalPosition: %v, score.Position: %v, Points: %v, country: %v\n", diff, absDiff, finalPosition, score.Position, Points[fmt.Sprint(absDiff)], score.Country)
		finalScore = finalScore + Points[fmt.Sprint(absDiff)]
	}
	return finalScore
}

func (s *Service) CalculateTotalScore(countries []country.Country, scores []Score, skipFinalCalculation bool) uint16 {
	semiWinners, finalCountryList := s.sortCountryList(countries)
	semiScore := s.calculateSemiScore(semiWinners, scores)
	finalScore := s.calculateFinalScore(finalCountryList, scores, skipFinalCalculation)

	return finalScore + semiScore
}
