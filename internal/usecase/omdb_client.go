package usecase

import (
	"context"
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/types"
)

type OmdbClientUseCase interface {
	GetMovieListByKeyword(context.Context, int64, string) (types.SearchMovieResult, error)
	GetMovieByMovieId(context.Context, string) (types.MovieDetails, error)
}
