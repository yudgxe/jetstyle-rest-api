package app

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/yudgxe/jetstyle-rest-api/internal/model"
	"github.com/yudgxe/jetstyle-rest-api/internal/store"
	"github.com/yudgxe/jetstyle-rest-api/internal/transport/rest/handler"
	"github.com/yudgxe/jetstyle-rest-api/pkg/middleware"

	_ "github.com/yudgxe/jetstyle-rest-api/docs"

	httpSwagger "github.com/swaggo/http-swagger/v2"

	"golang.org/x/crypto/bcrypt"
)

type Server struct {
	router *mux.Router
	store  store.Store
}

func NewServer(store store.Store) *Server {
	s := &Server{
		router: mux.NewRouter(),
		store:  store,
	}

	s.configurate()

	return s
}

func (s *Server) configurate() {
	s.router.StrictSlash(true)
	s.router.Use(middleware.Logging)

	s.router.NotFoundHandler = handler.NotFoundHandler()
	s.router.MethodNotAllowedHandler = handler.NotFoundHandler()
	
	sub := s.router.PathPrefix("/tasks").Subrouter()
	sub.Use(func(h http.Handler) http.Handler {
		return middleware.BasicAuth(h, func(login, password string) (bool, error) {
			user, err := s.store.User().FindByLogin(context.TODO(), login)
			if err != nil {
				return false, err
			}

			if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password)); err != nil {
				return false, err
			}

			return true, nil
		})
	})
	handler.NewTaskHandler(s.store).Bind(sub)
	

	s.router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	
	
	
}

func (s *Server) CreateUser(c *Config, login, password string) error {
	user := &model.User{
		Login:    login,
		Password: password,
	}

	if err := user.Validate(); err != nil {
		return err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.HashedPassword = string(passwordHash)

	if err := s.store.User().Create(context.Background(), user); err != nil {
		return err
	}

	log.Printf("User created his id = %d", user.ID)

	return nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
