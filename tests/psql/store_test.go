package psql_test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/yudgxe/jetstyle-rest-api/internal/app"
)

var databaseURL string

func TestMain(m *testing.M) {
	configPath := os.Getenv("DATABASE_TEST_CONFIG")

	if configPath == "" {
		configPath = "../../tests/psql/configs/config.toml"
	} else {
		filepath.Join("../../", configPath)
	}

	config := app.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatalf(err.Error())
	}

	databaseURL = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DB.Host,
		config.DB.Port,
		config.DB.User,
		config.DB.Password,
		config.DB.Name,
		config.DB.SSLMode,
	)

	os.Exit(m.Run())
}
