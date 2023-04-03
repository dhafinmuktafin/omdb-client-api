package types

import (
	pb "github.com/dhafinmuktafin/omdb-client-api/internal/delivery/grpc/proto"
	"time"
)

//ServerConfig is a struct
type ServerConfig struct {
	Port     string
	GRPCPort string
}

//DatabaseConfig is a struct
type DatabaseConfig struct {
	Slave                  string
	Master                 string
	MaxOpenConn            int
	MaxIdleConn            int
	ConnMaxLifetimeSeconds time.Duration
}

//HTTPConfig is a struct
type HTTPConfig struct {
	Timeout             time.Duration
	MaxIdleConns        int
	MaxIdleConnsPerHost int
	IdleConnTimeout     time.Duration
}

//CircuitBreakerConfig is a struct
type CircuitBreakerConfig struct {
	Timeout          time.Duration
	ErrorThreshold   int
	SuccessThreshold int
}

type SearchMovieResult struct {
	Search       []*pb.GetMovieByKeywordResponse_Search `json:"Search"`
	TotalResults string                                 `json:"totalResults"`
	Response     string                                 `json:"Response"`
}

type SearchDetail struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type MovieDetails struct {
	Title      string                                 `json:"Title"`
	Year       string                                 `json:"Year"`
	Rated      string                                 `json:"Rated"`
	Released   string                                 `json:"Released"`
	Runtime    string                                 `json:"Runtime"`
	Genre      string                                 `json:"Genre"`
	Director   string                                 `json:"Director"`
	Writer     string                                 `json:"Writer"`
	Actors     string                                 `json:"Actors"`
	Plot       string                                 `json:"Plot"`
	Language   string                                 `json:"Language"`
	Country    string                                 `json:"Country"`
	Awards     string                                 `json:"Awards"`
	Poster     string                                 `json:"Poster"`
	Ratings    []*pb.GetMovieByMovieIdResponse_Rating `json:"Ratings"`
	Metascore  string                                 `json:"Metascore"`
	ImdbRating string                                 `json:"imdbRating"`
	ImdbVotes  string                                 `json:"imdbVotes"`
	ImdbID     string                                 `json:"imdbID"`
	Type       string                                 `json:"Type"`
	DVD        string                                 `json:"DVD"`
	BoxOffice  string                                 `json:"BoxOffice"`
	Production string                                 `json:"Production"`
	Website    string                                 `json:"Website"`
	Response   string                                 `json:"Response"`
}

type Rating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}
