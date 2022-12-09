package constants

const (
	ContentRepositoryFile = `package repository

import (
	"@gomod/pkg/infra/db"
)

type Repository interface {
}

type @domain-capRepo struct {
	DBList *db.DatabaseList
}

func New@domain-capRepo(dbList *db.DatabaseList) @domain-capRepo {
	return @domain-capRepo{
		DBList: dbList,
	}
}`
	ContentUsecaseFile = `package usecase

import (
	"@gomod/config"
	repo "@gomod/internal/wrapper/repository"
	"@gomod/pkg/infra/db"
	"github.com/sirupsen/logrus"
)

type Usecase interface {
}

type @domain-capUsecase struct {
	Repo   repo.Repository
	Conf   *config.Config
	DBList *db.DatabaseList
	Log    *logrus.Logger
}

func New@domain-capUsecase(repository repo.Repository, conf *config.Config, dbList *db.DatabaseList, logger *logrus.Logger) @domain-capUsecase {
	return @domain-capUsecase{
		Repo:   repository,
		Conf:   conf,
		DBList: dbList,
		Log:    logger,
	}
}`

	ContentHandlerFile = `package delivery

import (
	"@gomod/config"
	"@gomod/internal/wrapper/usecase"
	"github.com/sirupsen/logrus"
)

type @domain-capHandler struct {
	Usecase usecase.Usecase
	Conf    *config.Config
	Log     *logrus.Logger
}

func New@domain-capHandler(uc usecase.Usecase, conf *config.Config, logger *logrus.Logger) @domain-capHandler {
	return @domain-capHandler{
		Usecase: uc,
		Conf:    conf,
		Log:     logger,
	}
}`

	ContentRoutesFile = `package @domain-lower

import (
	"@gomod/internal/wrapper/handler"
	"github.com/gofiber/fiber/v2"
)

func NewRoutes(api fiber.Router, handler handler.Handler) {
}`

	ContentImport = `import (
@import
)`
)

func ContentPackage(name string) string {
	return "package " + name
}
