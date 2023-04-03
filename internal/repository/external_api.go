package repository

import (
	"context"
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/types"
)

type ExternalAPIRepo interface {
	GetMovieListByKeyword(context.Context, int64, string) (types.SearchMovieResult, string, error)
	GetMovieDetails(context.Context, string) (types.MovieDetails, string, error)
}
