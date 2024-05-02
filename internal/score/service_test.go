package score

import (
	"testing"

	"github.com/vilmis04/eurovision-game-service/internal/country"
)

var countries = []country.Country{
	{
		Name:       "Serbia",
		Code:       "rs",
		Year:       2024,
		GameType:   "final",
		Score:      5,
		IsInFinal:  true,
		Artist:     "TEYA DORA",
		Song:       "RAMONDA",
		OrderSemi:  2,
		OrderFinal: 0,
	},
	{
		Name:       "Ireland",
		Code:       "ie",
		Year:       2024,
		GameType:   "final",
		Score:      50,
		IsInFinal:  true,
		Artist:     "Bambie Thug",
		Song:       "Doomsday Blue",
		OrderSemi:  4,
		OrderFinal: 0,
	},
	{
		Name:       "Poland",
		Code:       "pl",
		Year:       2024,
		GameType:   "final",
		Score:      10,
		IsInFinal:  true,
		Artist:     "LUNA",
		Song:       "The Tower",
		OrderSemi:  6,
		OrderFinal: 0,
	},
	{
		Name:       "Croatia",
		Code:       "hr",
		Year:       2024,
		GameType:   "final",
		Score:      45,
		IsInFinal:  true,
		Artist:     "Baby Lasagna",
		Song:       "Rim Tim Tagi Dim",
		OrderSemi:  7,
		OrderFinal: 0,
	},
	{
		Name:       "Iceland",
		Code:       "is",
		Year:       2024,
		GameType:   "final",
		Score:      15,
		IsInFinal:  true,
		Artist:     "Hera Björk",
		Song:       "Scared of Heights",
		OrderSemi:  8,
		OrderFinal: 0,
	},
	{
		Name:       "Slovenia",
		Code:       "si",
		Year:       2024,
		GameType:   "final",
		Score:      40,
		IsInFinal:  false,
		Artist:     "Raiven",
		Song:       "Veronika",
		OrderSemi:  9,
		OrderFinal: 0,
	},
	{
		Name:       "Finland",
		Code:       "fi",
		Year:       2024,
		GameType:   "final",
		Score:      20,
		IsInFinal:  false,
		Artist:     "Windows95man",
		Song:       "No Rules!",
		OrderSemi:  10,
		OrderFinal: 0,
	},
	{
		Name:       "Moldova",
		Code:       "md",
		Year:       2024,
		GameType:   "final",
		Score:      35,
		IsInFinal:  false,
		Artist:     "Natalia Barbu",
		Song:       "In The Middle",
		OrderSemi:  11,
		OrderFinal: 0,
	},
	{
		Name:       "Azerbaijan",
		Code:       "az",
		Year:       2024,
		GameType:   "final",
		Score:      25,
		IsInFinal:  false,
		Artist:     "FAHREE feat. Ilkin Dovlatov",
		Song:       "Özünlə Apar",
		OrderSemi:  12,
		OrderFinal: 0,
	},
	{
		Name:       "Australia",
		Code:       "au",
		Year:       2024,
		GameType:   "final",
		Score:      30,
		IsInFinal:  false,
		Artist:     "Electric Fields",
		Song:       "One Milkali (One Blood)",
		OrderSemi:  13,
		OrderFinal: 0,
	},
}

var scores = []Score{
	{
		Country:  "Serbia",
		Year:     2024,
		User:     "user1",
		GameType: "final",
		InFinal:  true,
		Position: 2,
	},
	{
		Country:  "Ireland",
		Year:     2024,
		User:     "user1",
		GameType: "final",
		InFinal:  false,
		Position: 1,
	},
	{
		Country:  "Poland",
		Year:     2024,
		User:     "user1",
		GameType: "final",
		InFinal:  true,
		Position: 6,
	},
	{
		Country:  "Croatia",
		Year:     2024,
		User:     "user1",
		GameType: "final",
		InFinal:  false,
		Position: 7,
	},
	{
		Country:  "Iceland",
		Year:     2024,
		User:     "user1",
		GameType: "final",
		InFinal:  true,
		Position: 8,
	},
	{
		Country:  "Slovenia",
		Year:     2024,
		User:     "user1",
		GameType: "final",
		InFinal:  false,
		Position: 9,
	},
	{
		Country:  "Finland",
		Year:     2024,
		User:     "user1",
		GameType: "final",
		InFinal:  true,
		Position: 10,
	},
	{
		Country:  "Moldova",
		Year:     2024,
		User:     "user1",
		GameType: "final",
		InFinal:  false,
		Position: 4,
	},
	{
		Country:  "Azerbaijan",
		Year:     2024,
		User:     "user1",
		GameType: "final",
		InFinal:  true,
		Position: 3,
	},
	{
		Country:  "Australia",
		Year:     2024,
		User:     "user1",
		GameType: "final",
		InFinal:  false,
		Position: 5,
	},
}

var service = NewService()

func TestSortCountryList(t *testing.T) {
	semiWinners, finalCountryList := service.sortCountryList(countries)
	if len(semiWinners) != 5 {
		t.Errorf("finalists should be 5, but got %v", len(semiWinners))
	}
	if len(finalCountryList) != 10 {
		t.Errorf("finalCountryList should be 10, but got %v", len(finalCountryList))
	}

}

func TestCalculateSemiScore(t *testing.T) {
	semiWinners, _ := service.sortCountryList(countries)
	semiScore := service.calculateSemiScore(semiWinners, scores)
	if semiScore != 15 {
		t.Errorf("semiScore should be 15, but got %v", semiScore)
	}
}

func TestCalculateFinalScore(t *testing.T) {
	_, finalCountryList := service.sortCountryList(countries)
	finalScore := service.calculateFinalScore(finalCountryList, scores)
	if finalScore != 71 {
		t.Errorf("finalScore should be 71, but got %v", finalScore)
	}
}

func TestCalculateTotalScore(t *testing.T) {
	totalScore := service.CalculateTotalScore(countries, scores)
	if totalScore != 86 {
		t.Errorf("totalScore should be 86, but got %v", totalScore)
	}
}
