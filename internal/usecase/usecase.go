package usecase

import (
	"github.com/haikalvidya/football-league/config"
	"github.com/haikalvidya/football-league/internal/middlewares"
	"github.com/haikalvidya/football-league/internal/repository"
)

type Usecase struct {
	League        ILeagueUsecase
	ContainLetter IContainLetterUsecase
}

type usecaseType struct {
	Repo       *repository.Repository
	Middleware *middlewares.CustomMiddleware
	ServerInfo *config.ServerConfig
}

func NewUsecase(repo *repository.Repository, mid *middlewares.CustomMiddleware, serverInfo *config.ServerConfig) *Usecase {
	usc := &usecaseType{Repo: repo, Middleware: mid, ServerInfo: serverInfo}

	return &Usecase{
		League:        (*leagueUsecase)(usc),
		ContainLetter: (*containLetterUsecase)(usc),
	}
}
