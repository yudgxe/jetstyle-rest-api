package app

import (
	"log"
	"net/http"

	"github.com/yudgxe/jetstyle-rest-api/internal/store/psql"

	"github.com/yudgxe/jetstyle-rest-api/pkg/database"
)

func setupPsqlStore(c *Config) (*psql.Store, error) {
	db, err := database.NewPostgres(database.PostgresConnInfo{
		Host:     c.DB.Host,
		Port:     c.DB.Port,
		User:     c.DB.User,
		Password: c.DB.Password,
		Name:     c.DB.Name,
		SSLMode:  c.DB.SSLMode,
	})
	if err != nil {
		return nil, err
	}

	return psql.New(db), nil
}

func CreateUser(c *Config, login, password string) error {
	store, err := setupPsqlStore(c)
	if err != nil {
		return err
	}

	if err := NewCreator(store).CreateUser(login, password); err != nil {
		return err
	}

	return nil
}

func Start(c *Config) error {
	store, err := setupPsqlStore(c)
	if err != nil {
		return err
	}
	s := NewServer(store)

	log.Printf("Server starting on %s port", c.BindAddr)
	return http.ListenAndServe(c.BindAddr, s)
}
