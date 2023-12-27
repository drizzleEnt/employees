package app

import (
	"context"
	"log"

	"github.com/drizzleent/emplyees/internal/api"
	"github.com/drizzleent/emplyees/internal/client/db"
	"github.com/drizzleent/emplyees/internal/client/db/pg"
	"github.com/drizzleent/emplyees/internal/config"
	"github.com/drizzleent/emplyees/internal/config/env"
	"github.com/drizzleent/emplyees/internal/repository"
	"github.com/drizzleent/emplyees/internal/repository/apipg"
	"github.com/drizzleent/emplyees/internal/service"
	emplService "github.com/drizzleent/emplyees/internal/service/employe"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	httpConfig config.HTTPConfig

	dbClient db.Client

	service service.ApiService

	repository repository.ApiRepository

	handler *api.Handler
}

func NewServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if nil == s.pgConfig {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}
		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if nil == s.httpConfig {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}
		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if nil == s.dbClient {
		cl, err := pg.New(ctx, s.PGConfig().Address())
		if err != nil {
			log.Fatalf("Failed to create db client %s", err.Error())
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("Failed to ping db %s", err.Error())
		}
		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) ApiRepository(ctx context.Context) repository.ApiRepository {
	if nil == s.repository {
		s.repository = apipg.NewRepository(s.DBClient(ctx))
	}

	return s.repository
}

func (s *serviceProvider) ApiService(ctx context.Context) service.ApiService {
	if nil == s.service {
		s.service = emplService.NewService(s.ApiRepository(ctx))
	}

	return s.service
}

func (s *serviceProvider) Handler(ctx context.Context) *api.Handler {
	if nil == s.handler {
		s.handler = api.NewHandler(s.ApiService(ctx))
	}

	return s.handler
}
