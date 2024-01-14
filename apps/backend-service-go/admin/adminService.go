package admin

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/vilmis04/eurovision-game-monorepo/tree/main/apps/backend-service-go/utils"
)

type adminService struct {
	connStr string
}

func NewService() *adminService {
	return &adminService{
		connStr: fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB")),
	}
}

func (s *adminService) connectToDB() (*sql.DB, error) {
	fmt.Println(s.connStr)
	db, err := sql.Open("postgres", s.connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (s *adminService) GetConfig() (*AdminConfig, error) {
	db, err := s.connectToDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var config AdminConfig
	var id uint8

	row := db.QueryRow("SELECT * FROM admin_config WHERE id=1")
	err = row.Scan(&id, &config.Year, &config.GameType, &config.IsVotingAcitve)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

type requestBody struct {
	Year           *uint16   `json:"year"`
	GameType       *GameType `json:"gameType"`
	IsVotingAcitve *bool     `json:"isVotingActive"`
}

func (s *adminService) UpdateConfig(req *http.Request) error {
	config, err := s.GetConfig()
	if err != nil {
		return err
	}
	var body requestBody

	err = utils.DecodeJson(req, &body)
	if err != nil {
		return err
	}

	if body.GameType != nil {
		config.GameType = *body.GameType
	}
	if body.Year != nil {
		config.Year = *body.Year
	}
	if body.IsVotingAcitve != nil {
		config.IsVotingAcitve = *body.IsVotingAcitve
	}

	// TODO: update database
	fmt.Printf("Updated config: %v\n", config)
	return nil
}

func (s *adminService) GetCountryDetails() (*CountryDetails, error) {
	// TODO: add get country details
	fmt.Println("Get country details")

	return &CountryDetails{}, nil
}

func (s *adminService) UpdateCountryDetails() error {
	// TODO: Take an object, update accordingly and return error
	fmt.Println("Update country details")

	return nil
}

func (s *adminService) CreateCountryDetails() error {
	// TODO: Take an object, create and return error
	fmt.Println("Create details for new country")

	return nil
}

func (s *adminService) DeleteCountry() error {
	// TODO: Take an object, create and return error
	fmt.Println("Delete by country details by id")

	return nil
}
