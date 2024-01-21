package admin

import (
	"database/sql"
	"encoding/json"
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
	db, err := sql.Open("postgres", s.connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (s *adminService) GetConfig() (*[]byte, error) {
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

	encodedConfig, err := json.Marshal(&config)
	if err != nil {
		return nil, err
	}

	return &encodedConfig, nil
}

func formatConfigUpdateQuery(property string) string {
	return fmt.Sprintf("UPDATE admin_config SET %v=$1 WHERE id=1", property)
}

func (s *adminService) UpdateConfig(req *http.Request) error {
	var body adminConfigRequestBody

	err := utils.DecodeRequestJson(req, &body)
	if err != nil {
		return err
	}

	db, err := s.connectToDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if body.GameType != nil {
		_, err = db.Exec(formatConfigUpdateQuery("gameType"), *body.GameType)
		if err != nil {
			return err
		}
	}
	if body.Year != nil {
		_, err = db.Exec(formatConfigUpdateQuery("year"), *body.Year)
		if err != nil {
			return err
		}
	}
	if body.IsVotingAcitve != nil {
		_, err = db.Exec(formatConfigUpdateQuery("isVotingActive"), *body.IsVotingAcitve)
		if err != nil {
			return err
		}
	}

	return nil
}
