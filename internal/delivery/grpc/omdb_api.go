package grpc

import (
	"context"
	pb "github.com/dhafinmuktafin/omdb-client-api/internal/delivery/grpc/proto"
	"github.com/prometheus/common/log"
)

func (o OmdbClientService) GetMovieListByKeyword(ctx context.Context, request *pb.GetMovieByKeywordRequest) (resp *pb.GetMovieByKeywordResponse, err error) {

	result, err := o.omdbClientUseCase.GetMovieListByKeyword(ctx, request.Page, request.Keyword)
	if err != nil {
		log.Errorf("[GetMovieListByKeyword] Error GetMovieListByKeyword %v", err)
		return nil, err
	}
	resp = &pb.GetMovieByKeywordResponse{
		SearchList:   result.Search,
		TotalResults: result.TotalResults,
		Response:     result.Response,
	}

	return
}

func (o OmdbClientService) GetMovieByMovieId(ctx context.Context, request *pb.GetMovieByMovieIdRequest) (resp *pb.GetMovieByMovieIdResponse, err error) {

	result, err := o.omdbClientUseCase.GetMovieByMovieId(ctx, request.MovieId)
	if err != nil {
		log.Errorf("[GetMovieByMovieId] Error GetMovieByMovieId %v", err)
		return nil, err
	}

	resp = &pb.GetMovieByMovieIdResponse{
		Title:      result.Title,
		Year:       result.Year,
		Rated:      result.Rated,
		Released:   result.Released,
		Runtime:    result.Runtime,
		Genre:      result.Genre,
		Director:   result.Director,
		Writer:     result.Writer,
		Actors:     result.Actors,
		Plot:       result.Plot,
		Language:   result.Language,
		Country:    result.Country,
		Awards:     result.Awards,
		Poster:     result.Poster,
		Ratings:    result.Ratings,
		Metascore:  result.Metascore,
		ImdbRating: result.ImdbRating,
		ImdbVotes:  result.ImdbVotes,
		ImdbID:     result.ImdbID,
		Type:       result.Type,
		DVD:        result.DVD,
		BoxOffice:  result.BoxOffice,
		Production: result.Production,
		Website:    result.Website,
		Response:   result.Response,
	}

	return
}
