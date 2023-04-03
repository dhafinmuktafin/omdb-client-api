package external_api

import (
	"context"
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/core/config"
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/types"
	"github.com/opentracing/opentracing-go"
	"net/http"
	"reflect"
	"testing"
	"time"

	circuit "github.com/eapache/go-resiliency/breaker"
	"github.com/jarcoal/httpmock"
)

func TestExternalAPI_GetSearchMovieByKeyword(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	breakers := map[int]*circuit.Breaker{}
	breakers[types.BreakerSearchMovieByKeyword] = circuit.New(10, 1, 5*time.Second)

	cfg := &config.Config{
		CircuitBreaker: types.CircuitBreakerConfig{
			ErrorThreshold:   10,
			SuccessThreshold: 1,
			Timeout:          5,
		},
	}

	type fields struct {
		Config         *config.Config
		Client         *http.Client
		CircuitBreaker *circuit.Breaker
	}
	type args struct {
		ctx        context.Context
		pageNumber int64
		keyword    string
	}

	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResponse types.SearchMovieResult
		wantErr      bool
		mock         func()
	}{
		{
			name: "TestGetSearchMovieByKeyword_success",
			args: args{
				pageNumber: 1,
				keyword:    "Batman",
				ctx:        context.Background(),
			},
			wantResponse: types.SearchMovieResult{},
			wantErr:      true,
			mock: func() {
				httpmock.Reset()
				httpmock.RegisterResponder("GET", "http://www.omdbapi.com/?apikey=faf7e5bb&s=Batman&page=2", func(req *http.Request) (*http.Response, error) {
					resp := types.SearchMovieResult{}
					return httpmock.NewJsonResponse(http.StatusOK, resp)
				})
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			var span opentracing.Span
			span, tt.args.ctx = opentracing.StartSpanFromContext(tt.args.ctx, "")
			defer span.Finish()
			e := &ExternalAPI{
				Config:         cfg,
				Client:         &http.Client{Timeout: 5 * time.Second},
				CircuitBreaker: breakers,
			}

			got, _, err := e.GetMovieListByKeyword(tt.args.ctx, tt.args.pageNumber, tt.args.keyword)
			if (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.wantResponse) {
				t.Errorf("ExternalAPI.GetSearchMovieByKeyword() got = %v, wantResp %v", got, tt.wantResponse)
				return
			}
		})
	}
}

func TestExternalAPI_GetMovieDetails(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	breakers := map[int]*circuit.Breaker{}
	breakers[types.BreakerGetMovieDetails] = circuit.New(10, 1, 5*time.Second)

	cfg := &config.Config{
		CircuitBreaker: types.CircuitBreakerConfig{
			ErrorThreshold:   10,
			SuccessThreshold: 1,
			Timeout:          5,
		},
	}

	type fields struct {
		Config         *config.Config
		Client         *http.Client
		CircuitBreaker *circuit.Breaker
	}
	type args struct {
		ctx     context.Context
		movieId string
	}

	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResponse types.MovieDetails
		wantErr      bool
		mock         func()
	}{
		{
			name: "TestExternalAPI_GetMovieDetails_success",
			args: args{
				movieId: "Batman",
				ctx:     context.Background(),
			},
			wantResponse: types.MovieDetails{},
			wantErr:      true,
			mock: func() {
				httpmock.Reset()
				httpmock.RegisterResponder("GET", "http://www.omdbapi.com/?apikey=faf7e5bb&s=Batman&page=2", func(req *http.Request) (*http.Response, error) {
					resp := types.MovieDetails{}
					return httpmock.NewJsonResponse(http.StatusOK, resp)
				})
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			var span opentracing.Span
			span, tt.args.ctx = opentracing.StartSpanFromContext(tt.args.ctx, "")
			defer span.Finish()
			e := &ExternalAPI{
				Config:         cfg,
				Client:         &http.Client{Timeout: 5 * time.Second},
				CircuitBreaker: breakers,
			}

			got, _, err := e.GetMovieDetails(tt.args.ctx, tt.args.movieId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.wantResponse) {
				t.Errorf("ExternalAPI.GetSearchMovieByKeyword() got = %v, wantResp %v", got, tt.wantResponse)
				return
			}
		})
	}
}
