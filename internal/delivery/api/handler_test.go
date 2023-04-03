package api

import (
	pb "github.com/dhafinmuktafin/omdb-client-api/internal/delivery/grpc/proto"
	"github.com/dhafinmuktafin/omdb-client-api/internal/delivery/util"
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/core/config"
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/types"
	"github.com/dhafinmuktafin/omdb-client-api/internal/usecase"
	mockUsecase "github.com/dhafinmuktafin/omdb-client-api/testfile/usecase"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestAPI_HealthCheck(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockCase := mockUsecase.NewMockOmdbClientUseCase(mockCtrl)
	router := mux.NewRouter()
	mWare := util.NewMiddleware(router, 10)
	respWriter := httptest.NewRecorder()
	cfg := &config.Config{}

	req, _ := http.NewRequest("GET", "/health", nil)

	type fields struct {
		OmdbClientCase usecase.OmdbClientUseCase
		Middleware     *util.HttpMiddleware
		Cfg            *config.Config
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *types.JSONResponse
		wantErr  bool
	}{
		{
			name: "test_api_health",
			fields: fields{
				OmdbClientCase: mockCase,
				Middleware:     mWare,
				Cfg:            cfg,
			},
			args: args{
				w: respWriter,
				r: req,
			},
			wantResp: &types.JSONResponse{
				Header: types.JSONRespHeader{
					ServerProcessTime: "",
					Message:           "",
					Reason:            "",
					Code:              200,
				},
				Data:        []string{"OmdbClientAPI service is up and running."},
				ErrorString: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &OmdbClientAPI{
				UseCase:    tt.fields.OmdbClientCase,
				Middleware: tt.fields.Middleware,
				Config:     tt.fields.Cfg,
			}
			gotResp, err := c.HealthCheck(tt.args.w, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("HealthCheck() get error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("HealthCheck() get = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestOmdbAPIClient_GetSearchMovieListByKeyword(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUseCase := mockUsecase.NewMockOmdbClientUseCase(mockCtrl)
	router := mux.NewRouter()
	midWare := util.NewMiddleware(router, 10)
	respWriter := httptest.NewRecorder()
	cfg := &config.Config{}

	req, _ := http.NewRequest("GET", "/omdb-client/search?page=1&keyword=Batman", nil)

	searchMovieResult := types.SearchMovieResult{
		Search: []*pb.GetMovieByKeywordResponse_Search{
			{
				Title:  "Batman",
				Year:   "2019",
				ImdbID: "aaaa",
				Type:   "movie",
				Poster: "https://m.media-amazon.com/images/M/XkEyXkFqcGdeQXVy_SX300.jpg",
			},
		},
		TotalResults: "1",
		Response:     "True",
	}

	type fields struct {
		Middleware *util.HttpMiddleware
		UseCase    usecase.OmdbClientUseCase
		Cfg        *config.Config
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *types.JSONResponse
		wantErr  bool
		mock     func()
	}{
		{
			name: "TestOmdbAPIClient_GetSearchMovieListByKeyword: success",
			fields: fields{
				Middleware: midWare,
				UseCase:    mockUseCase,
				Cfg:        cfg,
			},
			args: args{
				w: respWriter,
				r: req,
			},
			wantResp: &types.JSONResponse{
				Header: types.JSONRespHeader{
					ServerProcessTime: "",
					Message:           "",
					Reason:            "",
					Code:              200,
				},
				Data:        searchMovieResult,
				ErrorString: "",
			},
			wantErr: false,
			mock: func() {
				mockUseCase.EXPECT().GetMovieListByKeyword(gomock.Any(), gomock.Any(), gomock.Any()).Return(searchMovieResult, nil).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			c := &OmdbClientAPI{
				Middleware: tt.fields.Middleware,
				UseCase:    tt.fields.UseCase,
				Config:     tt.fields.Cfg,
			}
			gotResp, err := c.GetSearchMovieByKeyword(tt.args.w, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("OmdbClientAPI.GetSearchMovieByKeyword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("OmdbClientAPI.GetSearchMovieByKeyword() got= %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestOmdbAPIClient_GetMovieDetailByMovieId(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUseCase := mockUsecase.NewMockOmdbClientUseCase(mockCtrl)
	router := mux.NewRouter()
	midWare := util.NewMiddleware(router, 10)
	respWriter := httptest.NewRecorder()
	cfg := &config.Config{}

	req, _ := http.NewRequest("GET", "/omdb-client/movie/tt0372784", nil)

	movieResult := types.MovieDetails{
		Title:    "a",
		Year:     "b",
		Rated:    "c",
		Released: "d",
		Runtime:  "e",
		Genre:    "f",
		Director: "g",
		Writer:   "f",
		Actors:   "h",
		Plot:     "i",
		Language: "j",
		Country:  "k",
		Awards:   "l",
		Poster:   "m",
		Ratings: []*pb.GetMovieByMovieIdResponse_Rating{
			{
				Source: "1",
				Value:  "2",
			},
		},
		Metascore:  "n",
		ImdbRating: "o",
		ImdbVotes:  "p",
		ImdbID:     "q",
		Type:       "r",
		DVD:        "s",
		BoxOffice:  "t",
		Production: "u",
		Website:    "v",
		Response:   "True",
	}

	type fields struct {
		Middleware *util.HttpMiddleware
		UseCase    usecase.OmdbClientUseCase
		Cfg        *config.Config
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *types.JSONResponse
		wantErr  bool
		mock     func()
	}{
		{
			name: "TestOmdbAPIClient_GetMovieDetailByMovieId: success",
			fields: fields{
				Middleware: midWare,
				UseCase:    mockUseCase,
				Cfg:        cfg,
			},
			args: args{
				w: respWriter,
				r: req,
			},
			wantResp: &types.JSONResponse{
				Header: types.JSONRespHeader{
					ServerProcessTime: "",
					Message:           "",
					Reason:            "",
					Code:              200,
				},
				Data:        movieResult,
				ErrorString: "",
			},
			wantErr: false,
			mock: func() {
				mockUseCase.EXPECT().GetMovieByMovieId(gomock.Any(), gomock.Any()).Return(movieResult, nil).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			c := &OmdbClientAPI{
				Middleware: tt.fields.Middleware,
				UseCase:    tt.fields.UseCase,
				Config:     tt.fields.Cfg,
			}
			gotResp, err := c.GetMovieDetailByMovieId(tt.args.w, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("OmdbClientAPI.GetSearchMovieByKeyword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("OmdbClientAPI.GetSearchMovieByKeyword() got= %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
