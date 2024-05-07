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
	scores, err := s.storage.InitializeScores(user, year, gameType)
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
	if config.GameType == admin.GameTypeFinal {
		score.Position = body.Position
	} else {
		score.InFinal = body.InFinal
	}

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

	scores := []ScoreResponse{}
	if allGameTypes || config.GameType != admin.GameTypeFinal {
		semiScores, err := s.storage.GetAllOrSemiScores(user, config.GameType, config.Year)
		if err != nil {
			return nil, fmt.Errorf("failed to get scores for user %s: %v", user, err)
		}
		scores = append(scores, semiScores...)
		if len(scores) == 0 && !allGameTypes {
			fmt.Println("-------------")
			fmt.Println(len(scores))
			fmt.Println("-------------")
			scores, err = s.InitializeScores(user, config.Year, config.GameType)
			if err != nil {
				return nil, fmt.Errorf("failed to initialize scores for user %s: %v", user, err)
			}
		}
	}

	if allGameTypes || config.GameType == admin.GameTypeFinal {
		finalScores, err := s.storage.GetFinalScores(user, config.Year)
		if err != nil {
			return nil, fmt.Errorf("failed to get scores for user %s: %v", user, err)
		}
		scores = append(scores, finalScores...)
		if len(finalScores) < 25 && !allGameTypes {
			initFinalScores, err := s.InitializeScores(user, config.Year, admin.GameTypeFinal)
			if err != nil {
				return nil, fmt.Errorf("failed to initialize scores for user %s: %v", user, err)
			}
			scores = append(scores, initFinalScores...)
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
	finalists := []country.Country{}
	for _, country := range countries {
		if country.IsInFinal {
			finalists = append(finalists, country)
		}
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

	return s.CalculateTotalScore(finalists, scores, skipFinalCalculation), nil
}

func (s *Service) sortCountryList(countries []country.Country) map[string]CountryResult {
	slices.SortStableFunc(countries, func(a, b country.Country) int {
		return cmp.Compare(b.Score, a.Score)
	})

	finalCountryList := make(map[string]CountryResult)
	for index, country := range countries {
		finalCountryList[country.Name] = CountryResult{
			Name:     country.Name,
			Score:    country.Score,
			Position: index + 1,
		}
	}
	return finalCountryList
}

func (s *Service) calculateSemiScore(finalCountryList map[string]CountryResult, scores []Score) uint16 {
	var scoredPoints uint16 = 0
	for country := range finalCountryList {
		scoreIndex := slices.IndexFunc(scores, func(score Score) bool {
			return score.Country == country
		})
		if scoreIndex == -1 {
			continue
		}
		if scores[scoreIndex].InFinal {
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
	for country, results := range finalCountryList {
		scoreIndex := slices.IndexFunc(scores, func(score Score) bool {
			return score.Country == country
		})
		if scoreIndex == -1 {
			continue
		}
		score := scores[scoreIndex]

		scoredPosition := score.Position
		finalPosition := results.Position
		if scoredPosition == 0 || finalPosition == 0 {
			continue
		}
		diff := finalPosition - scoredPosition
		absDiff := math.Abs(float64(diff))
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

func (s *Service) CalculateTotalScore(countries []country.Country, scores []Score, skipFinalCalculation bool) uint16 {
	finalCountryList := s.sortCountryList(countries)
	semiScore := s.calculateSemiScore(finalCountryList, scores)
	finalScore := s.calculateFinalScore(finalCountryList, scores, skipFinalCalculation)

	return finalScore + semiScore
}
