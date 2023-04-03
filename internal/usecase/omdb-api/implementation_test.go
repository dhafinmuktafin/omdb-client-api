package omdbClient

import (
	"context"
	"errors"
	pb "github.com/dhafinmuktafin/omdb-client-api/internal/delivery/grpc/proto"
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/core/config"
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/types"
	"github.com/dhafinmuktafin/omdb-client-api/internal/repository"
	mock_repository "github.com/dhafinmuktafin/omdb-client-api/testfile/repository"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func Test_GetMovieListByKeyword(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockMySQLRepo := mock_repository.NewMockMySQLRepository(mockCtrl)
	mockExtAPI := mock_repository.NewMockExternalAPIRepo(mockCtrl)

	result := types.SearchMovieResult{
		Search: []*pb.GetMovieByKeywordResponse_Search{
			{
				Title:  "1",
				Year:   "1",
				ImdbID: "1",
				Type:   "1",
				Poster: "1",
			},
		},
		TotalResults: "1",
		Response:     "True",
	}

	type fields struct {
		Config     *config.Config
		MySQLRepo  repository.MySQLRepository
		ExtAPIRepo repository.ExternalAPIRepo
	}

	type args struct {
		ctx        context.Context
		pageNumber int64
		keyword    string
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantErr  bool
		wantData types.SearchMovieResult
		mock     func()
	}{
		{
			name: "TestGetMovieListByKeyword_success",
			fields: fields{
				Config:     &config.Config{},
				MySQLRepo:  mockMySQLRepo,
				ExtAPIRepo: mockExtAPI,
			},
			args: args{
				ctx:        context.Background(),
				pageNumber: 1,
				keyword:    "Batman",
			},
			wantErr:  false,
			wantData: result,
			mock: func() {
				mockExtAPI.EXPECT().GetMovieListByKeyword(gomock.Any(), gomock.Any(), gomock.Any()).Return(result, "aa", nil)
				mockMySQLRepo.EXPECT().InsertLogging(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "TestGetMovieListByKeyword_error",
			fields: fields{
				Config:     &config.Config{},
				MySQLRepo:  mockMySQLRepo,
				ExtAPIRepo: mockExtAPI,
			},
			args: args{
				ctx:        context.Background(),
				pageNumber: 1,
				keyword:    "Batman",
			},
			wantErr:  true,
			wantData: result,
			mock: func() {
				mockExtAPI.EXPECT().GetMovieListByKeyword(gomock.Any(), gomock.Any(), gomock.Any()).Return(result, "aa", errors.New("expected error"))
			},
		},
	}

	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			c := &OmdbClient{
				Config:    tt.fields.Config,
				MySQLRepo: tt.fields.MySQLRepo,
				ExtAPI:    tt.fields.ExtAPIRepo,
			}
			gotResult, err := c.GetMovieListByKeyword(tt.args.ctx, tt.args.pageNumber, tt.args.keyword)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMovieListByKeyword() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResult, tt.wantData) {
				t.Errorf("GetMovieListByKeyword() = %v, want %v", gotResult, tt.wantData)
			}
		})
	}
}

func Test_GetMovieByMovieId(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockMySQLRepo := mock_repository.NewMockMySQLRepository(mockCtrl)
	mockExtAPI := mock_repository.NewMockExternalAPIRepo(mockCtrl)

	result := types.MovieDetails{
		Title:    "1",
		Year:     "1",
		Rated:    "1",
		Released: "1",
		Runtime:  "1",
		Genre:    "1",
		Director: "1",
		Writer:   "1",
		Actors:   "1",
		Plot:     "1",
		Language: "1",
		Country:  "1",
		Awards:   "1",
		Poster:   "1",
		Ratings: []*pb.GetMovieByMovieIdResponse_Rating{
			{
				Source: "2",
				Value:  "2",
			},
		},
		Metascore:  "1",
		ImdbRating: "1",
		ImdbVotes:  "1",
		ImdbID:     "1",
		Type:       "1",
		DVD:        "1",
		BoxOffice:  "1",
		Production: "1",
		Website:    "1",
		Response:   "True",
	}

	type fields struct {
		Config     *config.Config
		MySQLRepo  repository.MySQLRepository
		ExtAPIRepo repository.ExternalAPIRepo
	}

	type args struct {
		ctx     context.Context
		movieId string
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantErr  bool
		wantData types.MovieDetails
		mock     func()
	}{
		{
			name: "TestGetMovieListByKeyword_success",
			fields: fields{
				Config:     &config.Config{},
				MySQLRepo:  mockMySQLRepo,
				ExtAPIRepo: mockExtAPI,
			},
			args: args{
				ctx:     context.Background(),
				movieId: "adsd",
			},
			wantErr:  false,
			wantData: result,
			mock: func() {
				mockExtAPI.EXPECT().GetMovieDetails(gomock.Any(), gomock.Any()).Return(result, "aa", nil)
				mockMySQLRepo.EXPECT().InsertLogging(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "TestGetMovieListByKeyword_error",
			fields: fields{
				Config:     &config.Config{},
				MySQLRepo:  mockMySQLRepo,
				ExtAPIRepo: mockExtAPI,
			},
			args: args{
				ctx:     context.Background(),
				movieId: "adsd",
			},
			wantErr:  true,
			wantData: result,
			mock: func() {
				mockExtAPI.EXPECT().GetMovieDetails(gomock.Any(), gomock.Any()).Return(result, "aa", errors.New("expected error"))
			},
		},
	}

	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			c := &OmdbClient{
				Config:    tt.fields.Config,
				MySQLRepo: tt.fields.MySQLRepo,
				ExtAPI:    tt.fields.ExtAPIRepo,
			}
			gotResult, err := c.GetMovieByMovieId(tt.args.ctx, tt.args.movieId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMovieByMovieId() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResult, tt.wantData) {
				t.Errorf("GetMovieByMovieId() = %v, want %v", gotResult, tt.wantData)
			}
		})
	}
}
