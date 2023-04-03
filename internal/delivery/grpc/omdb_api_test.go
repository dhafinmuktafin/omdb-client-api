package grpc

import (
	"context"
	pb "github.com/dhafinmuktafin/omdb-client-api/internal/delivery/grpc/proto"
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/core/config"
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/types"
	"github.com/dhafinmuktafin/omdb-client-api/internal/usecase"
	mockUsecase "github.com/dhafinmuktafin/omdb-client-api/testfile/usecase"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
	"time"
)

func TestOmdbService_GetMovieListByKeyword(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockOmdbAPIClientUseCase := mockUsecase.NewMockOmdbClientUseCase(mockCtrl)

	grpcOut := &pb.GetMovieByKeywordResponse{
		SearchList: []*pb.GetMovieByKeywordResponse_Search{
			{
				Title:  "1",
				Year:   "2",
				ImdbID: "3",
				Type:   "4",
				Poster: "5",
			},
		},
		TotalResults: "1",
		Response:     "True",
	}

	type fields struct {
		cfg         *config.Config
		localTime   *time.Location
		omdbUseCase usecase.OmdbClientUseCase
	}
	type args struct {
		ctx context.Context
		in  *pb.GetMovieByKeywordRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOut *pb.GetMovieByKeywordResponse
		wantErr bool
		mock    func()
	}{
		{
			name: "test_case_success",
			fields: fields{
				cfg:         &config.Config{},
				localTime:   &time.Location{},
				omdbUseCase: mockOmdbAPIClientUseCase,
			},
			args: args{
				ctx: context.Background(),
				in: &pb.GetMovieByKeywordRequest{
					Page:    1,
					Keyword: "Batman",
				},
			},
			wantOut: grpcOut,
			wantErr: false,
			mock: func() {
				caseReturn := types.SearchMovieResult{
					Search: []*pb.GetMovieByKeywordResponse_Search{
						{
							Title:  "1",
							Year:   "2",
							ImdbID: "3",
							Type:   "4",
							Poster: "5",
						},
					},
					TotalResults: "1",
					Response:     "True",
				}
				mockOmdbAPIClientUseCase.EXPECT().GetMovieListByKeyword(gomock.Any(), gomock.Any(), gomock.Any()).Return(caseReturn, nil).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			omdbAPIClientService := &OmdbClientService{
				cfg:               tt.fields.cfg,
				localTime:         tt.fields.localTime,
				omdbClientUseCase: tt.fields.omdbUseCase,
			}
			gotOut, err := omdbAPIClientService.GetMovieListByKeyword(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("OmdbService.GetMovieListByKeyword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("OmdbService.GetMovieListByKeyword() got= %v, want= %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestOmdbService_GetMovieByMovieId(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockOmdbAPIClientUseCase := mockUsecase.NewMockOmdbClientUseCase(mockCtrl)

	grpcOut := &pb.GetMovieByMovieIdResponse{
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
		cfg         *config.Config
		localTime   *time.Location
		omdbUseCase usecase.OmdbClientUseCase
	}
	type args struct {
		ctx context.Context
		in  *pb.GetMovieByMovieIdRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOut *pb.GetMovieByMovieIdResponse
		wantErr bool
		mock    func()
	}{
		{
			name: "test_case_success",
			fields: fields{
				cfg:         &config.Config{},
				localTime:   &time.Location{},
				omdbUseCase: mockOmdbAPIClientUseCase,
			},
			args: args{
				ctx: context.Background(),
				in: &pb.GetMovieByMovieIdRequest{
					MovieId: "aaa",
				},
			},
			wantOut: grpcOut,
			wantErr: false,
			mock: func() {
				caseReturn := types.MovieDetails{
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
				mockOmdbAPIClientUseCase.EXPECT().GetMovieByMovieId(gomock.Any(), gomock.Any()).Return(caseReturn, nil).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			omdbAPIClientService := &OmdbClientService{
				cfg:               tt.fields.cfg,
				localTime:         tt.fields.localTime,
				omdbClientUseCase: tt.fields.omdbUseCase,
			}
			gotOut, err := omdbAPIClientService.GetMovieByMovieId(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("OmdbService.GetMovieByMovieId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("OmdbService.GetMovieByMovieId() got= %v, want= %v", gotOut, tt.wantOut)
			}
		})
	}
}
