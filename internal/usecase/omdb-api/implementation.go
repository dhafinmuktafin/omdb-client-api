package omdbClient

import (
	"context"
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/types"
	"github.com/prometheus/common/log"
)

func (o OmdbClient) GetMovieListByKeyword(ctx context.Context, page int64, keyword string) (result types.SearchMovieResult, err error) {

	var url string
	result, url, err = o.ExtAPI.GetMovieListByKeyword(ctx, page, keyword)
	if err != nil {
		log.Errorf("[Repository][Init] error initialize mysql %+v", err)
		return
	}

	go func() {
		input := map[string]interface{}{
			"url":             url,
			"requestPayload":  url,
			"responsePayload": result.Response,
		}
		if err := o.MySQLRepo.InsertLogging(ctx, input); err != nil {
			log.Errorf("[GetMovieListByKeyword] input: %v\n err InsertLogging: %+v\n ", input, err)
		}
	}()

	return
}

func (o OmdbClient) GetMovieByMovieId(ctx context.Context, movieId string) (result types.MovieDetails, err error) {

	var url string
	result, url, err = o.ExtAPI.GetMovieDetails(ctx, movieId)
	if err != nil {
		log.Errorf("[Repository][Init] error initialize mysql %+v", err)
		return
	}

	go func() {
		input := map[string]interface{}{
			"url":              url,
			"request_payload":  url,
			"response_payload": result,
		}
		if err := o.MySQLRepo.InsertLogging(ctx, input); err != nil {
			log.Errorf("[GetMovieListByKeyword] input: %v\n err InsertLogging: %+v\n ", input, err)
		}
	}()

	return
}
