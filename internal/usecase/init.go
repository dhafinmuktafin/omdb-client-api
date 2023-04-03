package usecase

import (
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/core/config"
	"github.com/dhafinmuktafin/omdb-client-api/internal/repository"
	omdbClientUseCase "github.com/dhafinmuktafin/omdb-client-api/internal/usecase/omdb-api"
)

type UseCase struct {
	OmdbClientUseCase OmdbClientUseCase
}

func Init(cfg *config.Config, repository *repository.RepositoryTypes) (usecase UseCase) {
	omdbClientUsecase := omdbClientUseCase.New(cfg, repository)
	usecase = UseCase{
		OmdbClientUseCase: omdbClientUsecase,
	}
	return
}
