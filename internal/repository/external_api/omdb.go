package external_api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/types"
	"github.com/prometheus/common/log"
	"net/http"
)

func (e *ExternalAPI) GetMovieListByKeyword(ctx context.Context, pageNumber int64, keyword string) (result types.SearchMovieResult, url string, err error) {

	err = e.CircuitBreaker[types.BreakerSearchMovieByKeyword].Run(func() error {
		url = fmt.Sprintf("%s?apikey=%s&page=%d&s=%s", types.OmdbAPIURL, types.APIKey, pageNumber, keyword)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}

		req.WithContext(ctx)
		resp, err := e.Client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			log.Errorf("[GetMovieListByKeyword] Error decode response body: %+v\n", err)
			return err
		}

		return nil
	})

	return result, url, err
}

func (e *ExternalAPI) GetMovieDetails(ctx context.Context, movieId string) (result types.MovieDetails, url string, err error) {

	err = e.CircuitBreaker[types.BreakerGetMovieDetails].Run(func() error {
		url := fmt.Sprintf("%s?apikey=%s&i=%s", types.OmdbAPIURL, types.APIKey, movieId)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}

		req.WithContext(ctx)
		resp, err := e.Client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			log.Errorf("[GetMovieDetails] Error decode response body: %+v\n", err)
			return err
		}

		return nil
	})

	return result, url, err
}
